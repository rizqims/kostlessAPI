package repository

import (
	"database/sql"
	"kostless/model"
	"time"
)

type KosRepository interface {
	CreateKos(kos model.Kos) (model.Kos, error)
}

type kosRepository struct {
	db *sql.DB
}

func NewKosRepository(db *sql.DB) *kosRepository {
	return &kosRepository{db}
}

func (r *kosRepository) CreateKos(kos model.Kos) (model.Kos, error) {
	query := `INSERT INTO kos (name, address, room_count, coordinate, description, rules, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	timeNow := time.Now()
	err := r.db.QueryRow(query, kos.Name, kos.Address, kos.RoomCount, kos.Coordinate, kos.Description, kos.Rules, timeNow, timeNow).Scan(&kos.ID)
	if err != nil {
		return model.Kos{}, err
	}

	kos.CreatedAt = timeNow
	kos.UpdatedAt = timeNow

	return kos, nil
}
