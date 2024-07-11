package dto

type CreateVoucherReq struct {
	Name          string    `json:"name"`
	ExpiredDate   string `json:"expiredDate"`
	SeekerID      string    `json:"seekerId"`
	PercentAmount int       `json:"percentAmount"`
}

// {
// 	"name":"a",
// 	"expiredDate":"b",
// 	"seekerID":"c",
// 	"percentAmount":"open"
// }