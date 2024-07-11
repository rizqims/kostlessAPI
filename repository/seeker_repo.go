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
	GetSeekerByID(id string) (model.Seekers, error)
	GetAllSeekers() ([]model.Seekers, error)
	UpdateSeeker(id string, seeker model.Seekers) error
	DeleteSeeker(id string) error
	UpdateAttitudePoints(id string, attitudePoints int) error
}

// struct
type seekerRepo struct {
	db *sql.DB
}

// DeleteSeeker implements SeekerRepo.
func (s *seekerRepo) DeleteSeeker(id string) error {
	query := `DELETE FROM seekers WHERE id=$1`
	_, err := s.db.Exec(query, id)
	return err
}

// GetAllSeekers implements SeekerRepo.
func (s *seekerRepo) GetAllSeekers() ([]model.Seekers, error) {
	var seekers []model.Seekers
	query := `SELECT id, fullname, username, password, email, phone_number,  attitude_points, status, photo_profile, room_id, created_at, updated_at FROM seekers`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var seeker model.Seekers
		err := rows.Scan(&seeker.Id, &seeker.Fullname, &seeker.Username, &seeker.Password, &seeker.Email, &seeker.PhoneNumber, &seeker.AtitudePoits, &seeker.Status, &seeker.PhotoProfile, &seeker.RoomId, &seeker.CreatedAt, &seeker.UpdatedAt)
		if err != nil {
			return nil, err
		}
		seekers = append(seekers, seeker)
	}
	return seekers, nil
}

// GetSeekerByID implements SeekerRepo.
func (s *seekerRepo) GetSeekerByID(id string) (model.Seekers, error) {
	var seeker model.Seekers
	query := `SELECT id, fullname, username, password, email, phone_number, attitude_points, status, photo_profile, room_id, created_at, updated_at FROM seekers WHERE id=$1`
	err := s.db.QueryRow(query, id).Scan(&seeker.Id, &seeker.Fullname, &seeker.Username, &seeker.Password, &seeker.Email, &seeker.PhoneNumber, &seeker.AtitudePoits, &seeker.Status, &seeker.PhotoProfile, &seeker.RoomId, &seeker.CreatedAt, &seeker.UpdatedAt)
	return seeker, err
}

// UpdateAttitudePoints implements SeekerRepo.
func (s *seekerRepo) UpdateAttitudePoints(id string, attitudePoints int) error {
	query := `UPDATE seekers SET attitude_points=$1, updated_at=$2 WHERE id=$3`
	_, err := s.db.Exec(query, attitudePoints, time.Now(), id)
	return err
}

// UpdateSeeker implements SeekerRepo.
func (s *seekerRepo) UpdateSeeker(id string, seeker model.Seekers) error {
	query := `UPDATE seekers SET fullname=$1, username=$2, password=$3, email=$4, phone_number=$5, attitude_points=$6, status=$7, photo_profile=$8, room_id=$9, updated_at=$10 WHERE id=$11`
	_, err := s.db.Exec(query, seeker.Fullname, seeker.Username, seeker.Password, seeker.Email, seeker.PhoneNumber, seeker.AtitudePoits, seeker.Status, seeker.PhotoProfile, seeker.RoomId, time.Now(), seeker.Id)
	return err
}

// GetBySeeker implements SeekerRepo.
func (s *seekerRepo) GetBySeeker(username string) (model.Seekers, error) {
	var seeker model.Seekers
	err := s.db.QueryRow(`SELECT id, username, password, fullname, phone_number, attitude_points, status, photo_profile, created_at, updated_at FROM seekers WHERE username=$1`, username).Scan(&seeker.Id, &seeker.Username, &seeker.Password, &seeker.Fullname, &seeker.Email, &seeker.AtitudePoits, &seeker.Status, &seeker.PhotoProfile, &seeker.CreatedAt, &seeker.UpdatedAt)
	if err != nil {
		return model.Seekers{}, err
	}
	return seeker, nil
}

// CreatedNewUser implements UserRepo.
func (s *seekerRepo) CreatedNewSeeker(payload model.Seekers) (model.Seekers, error) {
	var seeker model.Seekers
	err := s.db.QueryRow(`INSERT INTO seekers (username, password, fullname, email, phone_number, attitude_points, status, photo_profile, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id, username, fullname, email, phone_number, attitude_points, status, photo_profile, created_at`, payload.Username, payload.Password, payload.Fullname, payload.Email, payload.PhoneNumber, payload.AtitudePoits, payload.Status, payload.PhotoProfile, time.Now()).Scan(&seeker.Id, &seeker.Username, &seeker.Fullname, &seeker.Email, &seeker.PhoneNumber, &seeker.AtitudePoits, &seeker.Status, &seeker.PhotoProfile, &seeker.CreatedAt)

	if err != nil {
		return model.Seekers{}, err
	}
	return seeker, nil
}

func NewUserSeeker(database *sql.DB) SeekerRepo {
	return &seekerRepo{db: database}
}
