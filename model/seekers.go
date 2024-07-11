package model

import "time"

type Seekers struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Fullname     string    `json:"fullname"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phoneNumber"`
	AtitudePoits int       `json:"atitudePoints"`
	Status       string    `json:"status"`
	PhotoProfile string    `json:"photoProfile"`
	RoomId       string    `json:"roomId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// register
// {
// 	"username":"",
// 	"password":"",
// 	"fullname":"",
// 	"email":"",
// 	"phoneNumber":"",
// 	"status":"",
// 	"photoProfile":""
// }

// login
// {
// 	"username":"khalid",
// 	"password":"123"
//   }