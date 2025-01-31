package model

import "time"

type User struct {
	Id           string    `json:"id"`
	Fullname     string    `json:"fullname"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phoneNumber"`
	PhotoProfile string    `json:"photoProfile"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// register
// {
// 	"fullname":"",
// 	"username":"",
// 	"password":"",
// 	"email":"",
// 	"phoneNumber":"",
// 	"photoProfile":""
// }

// login
// {
// 	"username":"",
// 	"password":""
// }