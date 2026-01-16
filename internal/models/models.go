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
