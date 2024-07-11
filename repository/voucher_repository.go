package repository

import (
	"database/sql"
	"kostless/model"
	"kostless/model/dto"
	"time"
)

type VoucherRepo interface {
	CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error)
	DeleteExpiredVoucher() error
}

type voucherRepo struct {
	db *sql.DB
}

func (v *voucherRepo) CreateVoucher(payload dto.CreateVoucherReq) (model.Voucher, error) {
	trans, err := v.db.Begin()
	if err != nil {
		return model.Voucher{}, err
	}

	var newVoucher model.Voucher
	err = v.db.QueryRow(`INSERT INTO vouchers (name, expired_date, seeker_id, percent_amount, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at`,
		payload.Name,
		payload.ExpiredDate,
		payload.SeekerID,
		payload.PercentAmount,
		time.Now(),
		time.Now()).Scan(
		&newVoucher.ID,
		&newVoucher.CreatedAt,
		&newVoucher.UpdatedAt,
	)
	if err != nil {
		trans.Rollback()
		return model.Voucher{}, err
	}
	newVoucher.Name = payload.Name
	newVoucher.ExpiredDate = payload.ExpiredDate
	newVoucher.SeekerID = payload.SeekerID
	newVoucher.PercentAmount = payload.PercentAmount
	err = trans.Commit()
	if err != nil {
		return model.Voucher{}, err
	}

	return newVoucher, nil
}

func (v *voucherRepo) DeleteExpiredVoucher() error{
	now := time.Now()
	_, err := v.db.Exec(`DELETE FROM vouchers WHERE expired_date < $1`, now)
	if err != nil {
		return err
	}
	return nil
}

func NewVoucherRepo(db *sql.DB) VoucherRepo {
	return &voucherRepo{
		db: db,
	}
}
