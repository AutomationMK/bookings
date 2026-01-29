package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/AutomationMK/bookings/internal/models"
	"golang.org/x/crypto/bcrypt"
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
	err := m.DB.QueryRow(ctx, stmt,
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

	_, err := m.DB.Exec(ctx, stmt,
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

	err := m.DB.QueryRow(ctx, stmt, roomID, start, end).Scan(&numRows)
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
		SELECT r.id, r.room_name, r.created_at, r.updated_at, r.bed_type, r.room_area, r.room_view, r.room_description, r.room_features, r.photo_links, r.room_route
		FROM rooms AS r
		WHERE r.id NOT IN (
			SELECT rr.room_id
			FROM room_restrictions AS rr
			WHERE
				$1 < rr.departure_date AND
				$2 > rr.arrival_date
		);`

	rows, err := m.DB.Query(ctx, stmt, start, end)
	if err != nil {
		return rooms, err
	}
	defer rows.Close()

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.BedType,
			&room.RoomArea,
			&room.RoomView,
			&room.RoomDescription,
			&room.RoomFeatures,
			&room.PhotoLinks,
			&room.RoomRoute,
		)
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
		SELECT id, room_name, created_at, updated_at, bed_type, room_area, room_view, room_description, room_features, photo_links, room_route
		FROM rooms
		WHERE id = $1;`

	row := m.DB.QueryRow(ctx, stmt, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
		&room.BedType,
		&room.RoomArea,
		&room.RoomView,
		&room.RoomDescription,
		&room.RoomFeatures,
		&room.PhotoLinks,
		&room.RoomRoute,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}

// GetRoomByRoute gets room data from database by room_route
func (m *postgresDBRepo) GetRoomByRoute(route string) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	stmt := `
		SELECT id, room_name, created_at, updated_at, bed_type, room_area, room_view, room_description, room_features, photo_links, room_route
		FROM rooms
		WHERE room_route = $1;`

	row := m.DB.QueryRow(ctx, stmt, route)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
		&room.BedType,
		&room.RoomArea,
		&room.RoomView,
		&room.RoomDescription,
		&room.RoomFeatures,
		&room.PhotoLinks,
		&room.RoomRoute,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}

// GetAllRooms returns all rooms in the database or an error if encountered
func (m *postgresDBRepo) GetAllRooms() ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	stmt := `
		SELECT id, room_name, created_at, updated_at, bed_type, room_area, room_view, room_description, room_features, photo_links, room_route
		FROM rooms;`

	rows, err := m.DB.Query(ctx, stmt)
	if err != nil {
		return rooms, err
	}
	defer rows.Close()

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.BedType,
			&room.RoomArea,
			&room.RoomView,
			&room.RoomDescription,
			&room.RoomFeatures,
			&room.PhotoLinks,
			&room.RoomRoute,
		)
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

// GetRoomCount gets the amont of rooms in the database
func (m *postgresDBRepo) GetRoomCount() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `SELECT COUNT(*) FROM rooms`

	var count int
	row := m.DB.QueryRow(ctx, stmt)
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetUserByID returns a user by id
func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		SELECT id, first_name, last_name, email, password, access_level, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	row := m.DB.QueryRow(ctx, stmt, id)

	var u models.User
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUser updates a user in the database
func (m *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		UPDATE users SET
			first_name = $1,
			last_name = $2,
			email = $3,
			access_level = $4,
			updated_at = $5
	`

	_, err := m.DB.Exec(ctx, stmt,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// Authenticate authenticates a user
func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		SELECT id, password
		FROM users
		WHERE email = $1
	`

	var id int
	var hashedPassword string

	row := m.DB.QueryRow(ctx, stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", nil
	}

	return id, hashedPassword, nil
}

// GetAllReservations returns all rooms in the database or an error if encountered
func (m *postgresDBRepo) GetAllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation

	stmt := `
		SELECT r.id, r.first_name, r.last_name, r.email, r.phone,
			r.arrival_date, r.departure_date, r.room_id, r.created_at,
			r.updated_at, rm.id, rm.room_name, rm.created_at,
			rm.updated_at, rm.bed_type, rm.room_area, rm.room_view,
			rm.room_description, rm.room_features, rm.photo_links, rm.room_route
		FROM reservations AS r
		LEFT JOIN rooms as rm on (r.room_id = rm.id)
		ORDER by r.arrival_date ASC;`

	rows, err := m.DB.Query(ctx, stmt)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()

	for rows.Next() {
		var reservation models.Reservation
		err := rows.Scan(
			&reservation.ID,
			&reservation.FirstName,
			&reservation.LastName,
			&reservation.Email,
			&reservation.Phone,
			&reservation.ArrivalDate,
			&reservation.DepartureDate,
			&reservation.RoomID,
			&reservation.CreatedAt,
			&reservation.UpdatedAt,
			&reservation.Room.ID,
			&reservation.Room.RoomName,
			&reservation.Room.CreatedAt,
			&reservation.Room.UpdatedAt,
			&reservation.Room.BedType,
			&reservation.Room.RoomArea,
			&reservation.Room.RoomView,
			&reservation.Room.RoomDescription,
			&reservation.Room.RoomFeatures,
			&reservation.Room.PhotoLinks,
			&reservation.Room.RoomRoute,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, reservation)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}
