package req

type TransCreateReq struct {
	RoomID    string    `json:"roomId"`
	SeekerID  string    `json:"seekerId"`
	StartDate string `json:"startDate"`
	Months    int       `json:"months"`
	PayLater  bool      `json:"payLater"`
	DueDate   string `json:"dueDate"`
}

// req json
// {
// 	"roomId": "18b3b4e9-0dd3-4803-9677-a2e83bcbc935",
// 	"seekerId": "7701b589-b342-4e4f-8f4f-6d572c9b8dcc",
// 	"startDate":"2024-05-05",
// 	"months": 3,
// 	"payLater": false,
// 	"dueDate": "2023-03-03"
// }