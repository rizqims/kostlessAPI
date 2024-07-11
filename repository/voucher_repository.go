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
	GetAllVoucher() ([]model.Voucher, error)
	GetVoucherBySeekerID(id string) ([]model.Voucher, error)
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
	expired, err := time.Parse(`2006-01-02`, payload.ExpiredDate)
	err = v.db.QueryRow(`INSERT INTO vouchers (name, expired_date, seeker_id, percent_amount, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at`,
		payload.Name,
		&expired,
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
	newVoucher.ExpiredDate = expired
	newVoucher.SeekerID = payload.SeekerID
	newVoucher.PercentAmount = payload.PercentAmount
	err = trans.Commit()
	if err != nil {
		return model.Voucher{}, err
	}

	return newVoucher, nil
}

func (v *voucherRepo) DeleteExpiredVoucher() error {
	now := time.Now()
	_, err := v.db.Exec(`DELETE FROM vouchers WHERE expired_date < $1`, now)
	if err != nil {
		return err
	}
	return nil
}

func (v *voucherRepo) GetAllVoucher() ([]model.Voucher, error) {
	var voucherList []model.Voucher
	rows, err := v.db.Query(`SELECT * FROM vouchers`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var voucher model.Voucher
		err := rows.Scan(
			&voucher.ID,
			&voucher.Name,
			&voucher.ExpiredDate,
			&voucher.SeekerID,
			&voucher.PercentAmount,
			&voucher.CreatedAt,
			&voucher.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		voucherList = append(voucherList, voucher)
	}
	return voucherList, nil
}

func (v *voucherRepo) GetVoucherBySeekerID(id string) ([]model.Voucher, error){
	var voucherList []model.Voucher
	rows, err := v.db.Query(`SELECT * FROM vouchers WHERE seeker_id=$1`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var voucher model.Voucher
		err := rows.Scan(
			&voucher.ID,
			&voucher.Name,
			&voucher.ExpiredDate,
			&voucher.SeekerID,
			&voucher.PercentAmount,
			&voucher.CreatedAt,
			&voucher.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		voucherList = append(voucherList, voucher)
	}
	return voucherList, nil
}

func NewVoucherRepo(db *sql.DB) VoucherRepo {
	return &voucherRepo{
		db: db,
	}
}
