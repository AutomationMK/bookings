package models

import "time"

// User holds user data
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room holds room data
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction holds restriction data
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation holds reservation data
type Reservation struct {
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
	Room          Room
}

// RoomRestriction holds room restriction data
type RoomRestriction struct {
	ID            int
	ArrivalDate   time.Time
	DepartureDate time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}
