package handlers

import (
	"net/http"

	"github.com/AutomationMK/bookings/pkg/config"
	"github.com/AutomationMK/bookings/pkg/models"
	"github.com/AutomationMK/bookings/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform buisiness logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Welcome to GO"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact handles the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Rooms handles the Rooms page
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "rooms.page.tmpl", &models.TemplateData{})
}

// Deluxe handles the Deluxe room page
func (m *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "deluxe-room.page.tmpl", &models.TemplateData{})
}

// Premium handles the Premium suite page
func (m *Repository) Premium(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "premium-suite.page.tmpl", &models.TemplateData{})
}
