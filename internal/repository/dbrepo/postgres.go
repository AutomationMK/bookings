package dbrepo

import (
	"context"
	"time"

	"github.com/AutomationMK/bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts all info from a Reservation model
// into the reservations table
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// create a context that limits the sql connection to 3 sec
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	stmt := `
		INSERT INTO reservations (
			first_name,
			last_name,
			email,
			phone,
			arrival_date,
			departure_date,
			room_id,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`

	// add Reservation model items and execute the query
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.ArrivalDate,
		res.DepartureDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsertRoomRestriction adds a RoomRestriction model to the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO room_restrictions (
			arrival_date,
			departure_date,
			room_id,
			reservation_id,
			created_at,
			updated_at,
			restriction_id
		) VALUES($1, $2, $3, $4, $5, $6, $7);`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.ArrivalDate,
		r.DepartureDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	)
	if err != nil {
		return err
	}

	return nil
}
