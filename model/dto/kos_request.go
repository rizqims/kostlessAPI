package dto

type KosRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	RoomCount   int    `json:"roomCount"`
	Coordinate  string `json:"coordinate"`
	Description string `json:"description"`
	Rules       string `json:"rules"`
}
