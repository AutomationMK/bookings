package dbrepo

import (
	"errors"
	"time"

	"github.com/AutomationMK/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts all info from a Reservation model
// into the reservations table
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// if room id is 2, then fail; otherwise pass
	if res.RoomID == 2 {
		return 0, errors.New("Test error in InsertReservation")
	}
	return 1, nil
}

// InsertRoomRestriction adds a RoomRestriction model to the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	// if room id is 1000, then fail; otherwise pass
	if r.RoomID == 1000 {
		return errors.New("Test error in InsertRestriction")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if the room is avalailable
// returns false if not available
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabityForALLRooms returns a slice of available rooms if any
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	// if the two dates are the same then return an error
	if start.Equal(end) {
		return rooms, errors.New("Test error in SearchAvailabilityForAllRooms")
		// if start date is after end date then return empty slice
	} else if start.After(end) {
		return rooms, nil
	}
	// if end date is after start date then return non-zero size slice
	room := models.Room{
		ID:              1,
		RoomName:        "Premium Suite",
		BedType:         "King",
		RoomArea:        123,
		RoomView:        "Garden",
		RoomDescription: "description",
		RoomFeatures:    []string{"feature1"},
		PhotoLinks:      []string{"photolink1.jpg"},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		RoomRoute:       "/static/rooms/test",
	}
	rooms = append(rooms, room)
	return rooms, nil
}

// GetRoomByID gets room data from database by room ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id == 1001 {
		return room, errors.New("Test error from GetRoomByID")
	}
	return room, nil
}

// GetRoomByRoute gets room data from database by room_route
func (m *testDBRepo) GetRoomByRoute(route string) (models.Room, error) {
	var room models.Room

	if route == "/rooms/invalid" {
		return room, errors.New("Test error from GetRoomByRoute")
	}

	return room, nil
}

// GetAllRooms returns all rooms in the database or an error if encountered
func (m *testDBRepo) GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

// GetRoomCount gets the amont of rooms in the database
func (m *testDBRepo) GetRoomCount() (int, error) {
	return 1, nil
}
