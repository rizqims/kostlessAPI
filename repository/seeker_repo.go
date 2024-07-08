package repository

import (
	"database/sql"
	"kostless-api/model"
	"time"
)

// interface
type SeekerRepo interface {
	CreatedNewSeeker(payload model.Seekers) (model.Seekers, error)
}

// struct
type seekerRepo struct {
	db *sql.DB
}

// CreatedNewUser implements UserRepo.
func (s *seekerRepo) CreatedNewSeeker(payload model.Seekers) (model.Seekers, error) {
	var seeker model.Seekers
	err := s.db.QueryRow(`INSERT INTO seekers (username, password, fullname, email, phone_number, attitude_points, status, room_id, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id, username, fullname, email, phone_number, attitude_points, status, room_id, created_at`, payload.Username, payload.Password, payload.Fullname, payload.Email, payload.PhoneNumber, payload.AtitudePoits, payload.Status, payload.RoomId, time.Now()).Scan(&seeker.Id, &seeker.Username, &seeker.Fullname, &seeker.Email, &seeker.PhoneNumber, &seeker.AtitudePoits, &seeker.Status, &seeker.CreatedAt)

	if err != nil {
		return model.Seekers{}, err
	}
	return seeker, nil
}

func NewUserSeeker(database *sql.DB) SeekerRepo {
	return &seekerRepo{db: database}
}

