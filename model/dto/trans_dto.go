package dto

import "time"

type TransCreateReq struct {
	RoomID    string `json:"roomId"`
	SeekerID  string `json:"seekerId"`
	StartDate string `json:"startDate"`
	VoucherID string `json:"voucherId"`
	Months    int    `json:"months"`
	PayLater  bool   `json:"payLater"`
	DueDate   string `json:"dueDate"`
}

type UpdatePaylaterReq struct {
	SeekerID string `json:"seekerId"`
	TransID  string `json:"transId"`
	Total    int    `json:"total"`
}

type UpdatePaylaterRes struct {
	TransID          string    `json:"transId"`
	IsOverdue        bool      `json:"isOverdue"`
	DueDate          time.Time `json:"dueDate"`
	CurrentTime      time.Time `json:"currentTime"`
	TotalWithPenalty int       `json:"totalWithPenalty"`
}

type AccPayment struct {
	TransID       string `json:"transId"`
	PaymentStatus string `json:"paymentStatus"`
}

// req json
// {
// 	"roomId": "18b3b4e9-0dd3-4803-9677-a2e83bcbc935",
// 	"seekerId": "7701b589-b342-4e4f-8f4f-6d572c9b8dcc",
// 	"startDate":"2024-05-05",
// 	"voucherId": "0",
// 	"months": 3,
// 	"payLater": false,
// 	"dueDate": "2023-03-03"
// }

// {
// 	"transID":"l",
// 	"payment_status":"",
// 	""
// }
