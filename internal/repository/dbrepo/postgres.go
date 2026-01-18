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

// SearchAvailabilityByDatesByRoomID returns true if the room is avalailable
// returns false if not available
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		SELECT COUNT(id)
		FROM room_restrictions
		WHERE
			room_id = $1 AND
			$2 < departure_date AND
			$3 > arrival_date;`

	var numRows int

	err := m.DB.QueryRowContext(ctx, stmt, roomID, start, end).Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

// SearchAvailabityForALLRooms returns a slice of available rooms if any
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	stmt := `
		SELECT r.id, r.room_name
		FROM rooms AS r
		WHERE r.id NOT IN (
			SELECT rr.room_id
			FROM room_restrictions AS rr
			WHERE
				$1 < rr.departure_date AND
				$2 > rr.arrival_date
		);`

	rows, err := m.DB.QueryContext(ctx, stmt, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomByID gets room data from database by room ID
func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	stmt := `
		SELECT id, room_name, created_at, updated_at
		FROM rooms
		WHERE id = $1;`

	row := m.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}
