package model

import (
	"time"

	"github.com/google/uuid"
)

type Kos struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	RoomCount   int       `json:"room_count"`
	Coordinate  string    `json:"coordinate"`
	Description string    `json:"description"`
	Rules       string    `json:"rules"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
