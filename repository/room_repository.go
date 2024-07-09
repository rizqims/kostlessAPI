package repository

import (
	"database/sql"
	"kostless/model"
	"time"
)

type RoomRepository interface {
	CreateRoom(room model.Room) (model.Room, error)
	GetRoomByID(id string) (model.Room, error)
	GetRoomByAvailability(availability string) ([]model.Room, error)
	GetRoomByPriceLowerThanOrEqual(price int) ([]model.Room, error)
}

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *roomRepository {
	return &roomRepository{db}
}

func (r *roomRepository) CreateRoom(room model.Room) (model.Room, error) {
	query := `INSERT INTO rooms (name, type, description, avail, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var id string
	timeNow := time.Now()
	err := r.db.QueryRow(query, room.Name, room.Type, room.Description, room.Avail, room.Price, timeNow, timeNow).Scan(&id)
	if err != nil {
		return model.Room{}, err
	}

	room.ID = id
	return room, nil
}

func (r *roomRepository) GetRoomByID(id string) (model.Room, error) {
	query := `SELECT id, name, type, description, avail, price, created_at, updated_at FROM rooms WHERE id = $1`
	var room model.Room
	err := r.db.QueryRow(query, id).Scan(&room.ID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) GetRoomByAvailability(availability string) ([]model.Room, error) {
	query := `SELECT id, name, type, description, avail, price, created_at, updated_at FROM rooms WHERE avail = $1`
	rows, err := r.db.Query(query, availability)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepository) GetRoomByPriceLowerThanOrEqual(price int) ([]model.Room, error) {
	query := `SELECT id, name, type, description, avail, price, created_at, updated_at FROM rooms WHERE price <= $1`
	rows, err := r.db.Query(query, price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
