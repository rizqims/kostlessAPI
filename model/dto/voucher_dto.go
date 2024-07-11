package dto

import "time"

type CreateVoucherReq struct {
	Name          string    `json:"name"`
	ExpiredDate   time.Time `json:"expiredDate"`
	SeekerID      string    `json:"seekerId"`
	PercentAmount int       `json:"percentAmount"`
}

// {
// 	"name":"a",
// 	"expiredDate":"b",
// 	"seekerID":"c",
// 	"percentAmount":"open"
// }