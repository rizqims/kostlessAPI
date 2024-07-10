package dto

type RoomRequest struct {
	KosID       string `json:"kos_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Avail       string `json:"avail"`
	Price       int    `json:"price"`
}
