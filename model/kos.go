package model

import (
	"time"
)

type Kos struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	RoomCount   int       `json:"room_count"`
	Coordinate  string    `json:"coordinate"`
	Description string    `json:"description"`
	Rules       string    `json:"rules"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rooms       []Room    `json:"rooms"`
}
