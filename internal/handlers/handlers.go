package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AutomationMK/bookings/internal/config"
	"github.com/AutomationMK/bookings/internal/driver"
	"github.com/AutomationMK/bookings/internal/forms"
	"github.com/AutomationMK/bookings/internal/helpers"
	"github.com/AutomationMK/bookings/internal/models"
	"github.com/AutomationMK/bookings/internal/render"
	"github.com/AutomationMK/bookings/internal/repository"
	"github.com/AutomationMK/bookings/internal/repository/dbrepo"
	"github.com/go-chi/chi"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Reserve handles the make-reservation page
func (m *Repository) Reserve(w http.ResponseWriter, r *http.Request) {
	// grab reservation from the session
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "You need to check for available rooms before making a reservation")
		http.Redirect(w, r, "/search-availability", http.StatusTemporaryRedirect)
		return
	}

	room, err := m.DB.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res.Room.RoomName = room.RoomName

	m.App.Session.Put(r.Context(), "reservation", res)

	// add the reservation to template data any map
	data := make(map[string]any)
	data["reservation"] = res

	// reformat the date.Time to string
	ad := res.ArrivalDate.Format("1/2/2006")
	dd := res.DepartureDate.Format("1/2/2006")
	// add the date strings to template data string map
	stringMap := make(map[string]string)
	stringMap["arrival_date"] = ad
	stringMap["departure_date"] = dd

	render.Template(w, r, "reservation.page.tmpl", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// PostReserve handles the posting of a reservation form
func (m *Repository) PostReserve(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("Can't get reservation data from session"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Email = r.Form.Get("email")
	reservation.Phone = r.Form.Get("phone")

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]any)
		data["reservation"] = reservation

		render.Template(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	restriction := models.RoomRestriction{
		ArrivalDate:   reservation.ArrivalDate,
		DepartureDate: reservation.DepartureDate,
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get item from sessio")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]any)
	data["reservation"] = reservation

	ad := reservation.ArrivalDate.Format("1/2/2006")
	dd := reservation.DepartureDate.Format("1/2/2006")
	stringMap := make(map[string]string)
	stringMap["arrival_date"] = ad
	stringMap["departure_date"] = dd

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}

// Availability handles the search-availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles post data from search-availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// parse date data
	ad := r.Form.Get("arrival_date")
	dd := r.Form.Get("departure_date")
	layout := "1/2/2006"
	arrivalDate, err := time.Parse(layout, ad)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	departureDate, err := time.Parse(layout, dd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err := m.DB.SearchAvailabilityForAllRooms(arrivalDate, departureDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(rooms) == 0 {
		// no availability
		m.App.Session.Put(r.Context(), "error", "No availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		ArrivalDate:   arrivalDate,
		DepartureDate: departureDate,
	}

	m.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "available-rooms.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

type jsonResponse struct {
	OK            bool   `json:"ok"`
	Message       string `json:"message"`
	RoomID        string `json:"room_id"`
	ArrivalDate   string `json:"arrival_date"`
	DepartureDate string `json:"departure_date"`
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	ad := r.Form.Get("arrival_date")
	dd := r.Form.Get("departure_date")

	layout := "1/2/2006"
	arrivalDate, err := time.Parse(layout, ad)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	departureDate, err := time.Parse(layout, dd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	available, err := m.DB.SearchAvailabilityByDatesByRoomID(arrivalDate, departureDate, roomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	resp := jsonResponse{
		OK:            available,
		Message:       "",
		ArrivalDate:   ad,
		DepartureDate: dd,
		RoomID:        strconv.Itoa(roomID),
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// BookRoom takes url parameters and builds reservation session
// user is redirected to the /make-reservation route
func (m *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ad := r.URL.Query().Get("ad")
	dd := r.URL.Query().Get("dd")

	layout := "1/2/2006"
	arrivalDate, err := time.Parse(layout, ad)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	departureDate, err := time.Parse(layout, dd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var res models.Reservation

	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res.Room.RoomName = room.RoomName
	res.RoomID = roomID
	res.ArrivalDate = arrivalDate
	res.DepartureDate = arrivalDate

	m.App.Session.Put(r.Context(), "reservation", res)

	log.Println(roomID, arrivalDate, departureDate)
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}

// Contact handles the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Rooms handles the Rooms page
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "rooms.page.tmpl", &models.TemplateData{})
}

// Deluxe handles the Deluxe room page
func (m *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "deluxe-room.page.tmpl", &models.TemplateData{})
}

// Premium handles the Premium suite page
func (m *Repository) Premium(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "premium-suite.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID

	m.App.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
