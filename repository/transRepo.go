package repository

import (
	"database/sql"
	"fmt"
	"kostless/model"
	"kostless/model/dto"
	"time"
)

type TransRepo interface {
	CreateTrans(payload model.Trans) (model.Trans, error)
	GetTransByID(id string) (model.Trans, error)
	GetTransHistory() ([]model.Trans, error)
	GetPaylaterList() ([]model.Trans, error)
	GetTransByMonth(startDate, endDate string) ([]model.Trans, error)
	UpdatePaylater(payload dto.UpdatePaylaterReq) (model.Trans, error)
}

type transRepo struct {
	db *sql.DB
}

func (t *transRepo) CreateTrans(payload model.Trans) (model.Trans, error) {
	trans, err := t.db.Begin()
	if err != nil {
		return model.Trans{}, err
	}

	err = t.db.QueryRow(`INSERT INTO bookings (room_id, seeker_id, start_date, end_date, total, pay_later, due_date, payment_status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, created_at, updated_at`,
		payload.RoomID,
		payload.SeekerID,
		payload.StartDate,
		payload.EndDate,
		payload.Total,
		payload.PayLater,
		payload.DueDate,
		payload.PaymentStatus,
		time.Now(),
		time.Now()).Scan(
		&payload.ID,
		&payload.CreatedAt,
		&payload.UpdatedAt,
	)
	if err != nil {
		trans.Rollback()
		return model.Trans{}, err
	}

	err = trans.Commit()
	if err != nil {
		return model.Trans{}, err
	}

	return payload, nil
}

func (t *transRepo) GetTransByID(id string) (model.Trans, error) {
	var trans model.Trans
	var disc, total sql.NullInt64
	err := t.db.QueryRow(`SELECT * FROM bookings WHERE id=$1`, id).Scan(
		&trans.ID,
		&trans.RoomID,
		&trans.SeekerID,
		&trans.StartDate,
		&trans.EndDate,
		&disc,
		&total,
		&trans.PayLater,
		&trans.DueDate,
		&trans.PaymentStatus,
		&trans.CreatedAt,
		&trans.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return model.Trans{}, err
	}

	trans.Discount = int(disc.Int64)
	trans.Total = int(total.Int64)

	return trans, err
}

func (u *transRepo) GetTransHistory() ([]model.Trans, error) {
	rows, err := u.db.Query(`SELECT * FROM bookings`)
	if err != nil {
		return nil, fmt.Errorf("GetTransHistoryRepo: get trans error: ", err)
	}

	var transList = []model.Trans{}
	var disc, total sql.NullInt64
	for rows.Next() {
		var trans model.Trans
		err := rows.Scan(
			&trans.ID,
			&trans.RoomID,
			&trans.SeekerID,
			&trans.StartDate,
			&trans.EndDate,
			&disc,
			&total,
			&trans.PayLater,
			&trans.DueDate,
			&trans.PaymentStatus,
			&trans.CreatedAt,
			&trans.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		trans.Discount = int(disc.Int64)
		trans.Total = int(total.Int64)
		transList = append(transList, trans)
	}
	return transList, nil
}

func (t *transRepo) GetPaylaterList() ([]model.Trans, error) {
	rows, err := t.db.Query(`SELECT * FROM bookings WHERE pay_later=$1`, true)
	if err != nil {
		return nil, fmt.Errorf("GetPaylaterListRepo: get trans error: ", err)
	}

	var transList = []model.Trans{}
	var disc, total sql.NullInt64
	for rows.Next() {
		var trans model.Trans
		err := rows.Scan(
			&trans.ID,
			&trans.RoomID,
			&trans.SeekerID,
			&trans.StartDate,
			&trans.EndDate,
			&disc,
			&total,
			&trans.PayLater,
			&trans.DueDate,
			&trans.PaymentStatus,
			&trans.CreatedAt,
			&trans.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		trans.Discount = int(disc.Int64)
		trans.Total = int(total.Int64)
		transList = append(transList, trans)
	}
	return transList, nil
}

func (t *transRepo) GetTransByMonth(startDate, endDate string) ([]model.Trans, error) {
	rows, err := t.db.Query(`SELECT * FROM bookings WHERE start_date >= $1 AND start_date <= $2`, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var transList = []model.Trans{}
	var disc, total sql.NullInt64
	for rows.Next() {
		var trans model.Trans
		err := rows.Scan(
			&trans.ID,
			&trans.RoomID,
			&trans.SeekerID,
			&trans.StartDate,
			&trans.EndDate,
			&disc,
			&total,
			&trans.PayLater,
			&trans.DueDate,
			&trans.PaymentStatus,
			&trans.CreatedAt,
			&trans.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transList = append(transList, trans)
	}
	fmt.Println(transList)
	return transList, nil
}

func (t *transRepo) UpdatePaylater(payload dto.UpdatePaylaterReq) (model.Trans, error){
	var updatedTrans model.Trans
	err := t.db.QueryRow(`UPDATE bookings SET pay_later=false, updated_at=$1 WHERE id=$2 RETURNING due_date, total, seeker_id`, time.Now(), payload.TransID).Scan(&updatedTrans.DueDate, &updatedTrans.Total, &updatedTrans.SeekerID)
	if err != nil {
		return model.Trans{}, err
	}
	return updatedTrans, nil
}

func NewTransRepo(db *sql.DB) TransRepo {
	return &transRepo{
		db: db,
	}
}
