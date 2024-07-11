package repository

import (
	"database/sql"
	"kostless/model"
	"time"
)

type RoomRepository interface {
	CreateRoom(room model.Room) (model.Room, error)
	GetAllRooms() ([]model.Room, error)
	GetRoomByID(id string) (model.Room, error)
	GetRoomByAvailability(availability string) ([]model.Room, error)
	GetRoomByPriceLowerThanOrEqual(price int) ([]model.Room, error)
	UpdateRoom(room model.Room) (model.Room, error)
	DeleteRoom(id string) error
}

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *roomRepository {
	return &roomRepository{db}
}

func (r *roomRepository) CreateRoom(room model.Room) (model.Room, error) {
	query := `INSERT INTO rooms (kos_id, name, type, description, avail, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`
	timeNow := time.Now()
	err := r.db.QueryRow(query, room.KosID, room.Name, room.Type, room.Description, room.Avail, room.Price, timeNow, timeNow).Scan(&room.ID, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) GetAllRooms() ([]model.Room, error) {
	query := `SELECT id, kos_id, name, type, description, avail, price, created_at, updated_at FROM rooms`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.KosID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepository) GetRoomByID(id string) (model.Room, error) {
	query := `SELECT id, kos_id, name, type, description, avail, price, created_at, updated_at FROM rooms WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var room model.Room
	err := row.Scan(&room.ID, &room.KosID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) GetRoomByAvailability(availability string) ([]model.Room, error) {
	query := `SELECT id, kos_id, name, type, description, avail, price, created_at, updated_at FROM rooms WHERE avail = $1`
	rows, err := r.db.Query(query, availability)
	if err != nil {
		return nil, err
	}

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.KosID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepository) GetRoomByPriceLowerThanOrEqual(price int) ([]model.Room, error) {
	query := `SELECT id, kos_id, name, type, description, avail, price, created_at, updated_at FROM rooms WHERE price <= $1`
	rows, err := r.db.Query(query, price)
	if err != nil {
		return nil, err
	}

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.KosID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepository) UpdateRoom(room model.Room) (model.Room, error) {
	query := `UPDATE rooms SET kos_id = $1, name = $2, type = $3, description = $4, avail = $5, price = $6, updated_at = $7 WHERE id = $8 RETURNING updated_at`
	err := r.db.QueryRow(query, room.KosID, room.Name, room.Type, room.Description, room.Avail, room.Price, time.Now(), room.ID).Scan(&room.UpdatedAt)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) DeleteRoom(id string) error {
	query := `DELETE FROM rooms WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
