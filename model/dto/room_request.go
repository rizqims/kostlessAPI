package dto

type RoomRequest struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Avail       string `json:"avail"`
	Price       int    `json:"price"`
}

// {
// 	"name":"",
// 	"type":"",
// 	"description":"",
// 	"avail":"",
// 	"price":""
// }