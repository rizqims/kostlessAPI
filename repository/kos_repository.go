package repository

import (
	"database/sql"
	"kostless/model"
	"time"
)

type KosRepository interface {
	CreateKos(kos model.Kos) (model.Kos, error)
	UpdateKos(kos model.Kos) (model.Kos, error)
	DeleteKos(id string) error
	GetKosByID(id string) (model.Kos, error)
}

type kosRepository struct {
	db *sql.DB
}

func NewKosRepository(db *sql.DB) *kosRepository {
	return &kosRepository{db}
}

func (r *kosRepository) CreateKos(kos model.Kos) (model.Kos, error) {
	query := `INSERT INTO kos (name, address, room_count, coordinate, description, rules, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`
	timeNow := time.Now()
	err := r.db.QueryRow(query, kos.Name, kos.Address, kos.RoomCount, kos.Coordinate, kos.Description, kos.Rules, timeNow, timeNow).Scan(&kos.ID, &kos.CreatedAt, &kos.UpdatedAt)
	if err != nil {
		return model.Kos{}, err
	}

	return kos, nil
}

func (r *kosRepository) UpdateKos(kos model.Kos) (model.Kos, error) {
	query := `UPDATE kos SET name = $1, address = $2, room_count = $3, coordinate = $4, description = $5, rules = $6, updated_at = $7 WHERE id = $8 RETURNING updated_at`
	err := r.db.QueryRow(query, kos.Name, kos.Address, kos.RoomCount, kos.Coordinate, kos.Description, kos.Rules, time.Now(), kos.ID).Scan(&kos.UpdatedAt)
	if err != nil {
		return model.Kos{}, err
	}

	return kos, nil
}

func (r *kosRepository) DeleteKos(id string) error {
	query := `DELETE FROM kos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *kosRepository) GetKosByID(id string) (model.Kos, error) {
	query := `SELECT id, name, address, room_count, coordinate, description, rules, created_at, updated_at FROM kos WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var kos model.Kos
	err := row.Scan(&kos.ID, &kos.Name, &kos.Address, &kos.RoomCount, &kos.Coordinate, &kos.Description, &kos.Rules, &kos.CreatedAt, &kos.UpdatedAt)
	if err != nil {
		return model.Kos{}, err
	}

	query = `SELECT id, kos_id, name, type, description, avail, price, created_at, updated_at FROM rooms WHERE kos_id = $1`
	rows, err := r.db.Query(query, kos.ID)
	if err != nil {
		return model.Kos{}, err
	}

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.KosID, &room.Name, &room.Type, &room.Description, &room.Avail, &room.Price, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return model.Kos{}, err
		}
		rooms = append(rooms, room)
	}

	kos.Rooms = rooms

	return kos, nil
}
