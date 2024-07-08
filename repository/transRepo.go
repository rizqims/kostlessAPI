package repository

import (
	"database/sql"
	"kostless/model"
	"time"
)

type TransRepo interface {
  CreateTrans(payload model.Trans) (model.Trans, error)
}

type transRepo struct {
  db *sql.DB
}

func (t *transRepo) CreateTrans(payload model.Trans) (model.Trans, error){
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

func NewTransRepo(db *sql.DB) TransRepo{
	return &transRepo{
		db: db,
	}
}