package repository

import (
	"time"

	"github.com/AutomationMK/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	GetRoomByRoute(route string) (models.Room, error)
	GetAllRooms() ([]models.Room, error)
	GetRowCount(table string) (int, error)
}
