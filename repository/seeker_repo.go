package repository

import (
	"database/sql"
	"fmt"
	"kostless/model"
	"time"
)

// interface
type SeekerRepo interface {
	CreatedNewSeeker(payload model.Seekers) (model.Seekers, error)
	GetBySeeker(username string) (model.Seekers, error)
	GetByID(id string)(model.Seekers, error)
}

// struct
type seekerRepo struct {
	db *sql.DB
}

// GetBySeeker implements SeekerRepo.
func (s *seekerRepo) GetBySeeker(username string) (model.Seekers, error) {
	var seeker model.Seekers
	err := s.db.QueryRow(`SELECT id, username, password, fullname, phone_number, status, created_at, updated_at FROM seekers WHERE username=$1`, username).Scan(&seeker.Id, &seeker.Username, &seeker.Password, &seeker.Fullname, &seeker.Email, &seeker.Status, &seeker.CreatedAt, &seeker.UpdatedAt)
	if err != nil {
		return model.Seekers{}, err
	}
	return seeker, nil
}

func (s *seekerRepo) GetByID(id string) (model.Seekers, error) {
	var seeker model.Seekers
	err := s.db.QueryRow(`SELECT id, username, password, fullname, email, phone_number, status, created_at, updated_at FROM seekers WHERE id=$1`, id).Scan(&seeker.Id, &seeker.Username, &seeker.Password, &seeker.Fullname, &seeker.Email, &seeker.Status, &seeker.CreatedAt, &seeker.UpdatedAt)
	if err != nil {
		return model.Seekers{}, err
	}
	return seeker, nil
}

// CreatedNewUser implements UserRepo.
func (s *seekerRepo) CreatedNewSeeker(payload model.Seekers) (model.Seekers, error) {
	var seeker model.Seekers
	err := s.db.QueryRow(`INSERT INTO seekers (username, password, fullname, email, phone_number, status, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id, username, fullname, email, phone_number,  status,  created_at`, payload.Username, payload.Password, payload.Fullname, payload.Email, payload.PhoneNumber, payload.Status, time.Now()).Scan(&seeker.Id, &seeker.Username, &seeker.Fullname, &seeker.Email, &seeker.PhoneNumber, &seeker.Status, &seeker.CreatedAt)

	if err != nil {
		fmt.Print("err===", err)
		return model.Seekers{}, err
	}
	return seeker, nil
}

func NewUserSeeker(database *sql.DB) SeekerRepo {
	return &seekerRepo{db: database}
}
