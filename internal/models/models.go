package models

import "time"

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users holds user data
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms holds room data
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions holds restrictions data
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations holds reservation data
type Reservations struct {
	ID            int
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	ArrivalDate   time.Time
	DepartureDate time.Time
	RoomID        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Rooms
}

// RoomRestrictions holds room restriction data
type RoomRestrictions struct {
	ID            int
	ArrivalDate   time.Time
	DepartureDate time.Time
	room_id       int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Rooms
	Reservation   Reservations
	Restriction   Restrictions
}
