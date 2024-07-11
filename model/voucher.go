package model

import "time"

type Voucher struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	ExpiredDate   time.Time `json:"expiredDate"`
	SeekerID      string    `json:"seekerId"`
	PercentAmount int       `json:"percentAmount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}