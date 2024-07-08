package model

import "time"

type Seekers struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Fullname     string    `json:"fullname"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	AtitudePoits int       `json:"atitude_points"`
	Status       string    `json:"status"`
	RoomId       string    `json:"room_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
