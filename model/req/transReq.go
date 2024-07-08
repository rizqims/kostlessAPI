package req

import "time"

type TransCreateReq struct {
	RoomID    string    `json:"roomId"`
	SeekerID  string    `json:"seekerId"`
	StartDate time.Time `json:"startDate"`
	Months    int `json:"months"`
	PayLater  bool      `json:"payLater"`
	DueDate   time.Time `json:"dueDate"`
}
