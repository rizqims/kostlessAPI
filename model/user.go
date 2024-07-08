package model

import "time"

type User struct {
	Id          string    `json:"id"`
	Fullname    string    `json:"fullname"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
