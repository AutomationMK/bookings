package dbrepo

import (
	"time"

	"github.com/AutomationMK/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts all info from a Reservation model
// into the reservations table
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction adds a RoomRestriction model to the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
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

	return rooms, nil
}

// GetRoomByID gets room data from database by room ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	return room, nil
}

// GetRoomByRoute gets room data from database by room_route
func (m *testDBRepo) GetRoomByRoute(route string) (models.Room, error) {
	var room models.Room

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
