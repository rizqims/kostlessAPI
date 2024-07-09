package model

import "time"

type Trans struct {
	ID            string    `json:"id"`
	RoomID        string    `json:"roomId"`
	SeekerID      string    `json:"seekerId"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Discount      int       `json:"discount"`
	Total         int       `json:"total"`
	PayLater      bool      `json:"payLater"`
	DueDate       time.Time `json:"dueDate"`
	PaymentStatus string    `json:"paymentStatus"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}