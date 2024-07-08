package repository

import (
	"database/sql"
	"kostless-api/model"
	"time"
)

// interface
type UserRepo interface {
	CreatedNewUser(payload model.User) (model.User, error)
}

// struct
type userRepo struct {
	db *sql.DB
}

// CreatedNewUser implements UserRepo.
func (u *userRepo) CreatedNewUser(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`INSERT INTO users (fullname, username, password, email, phone_number, updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, fullname, username, email, phone_number, created_at`, payload.Fullname, payload.Username, payload.Password, payload.Email, payload.PhoneNumber, time.Now()).Scan(&user.Id, &user.Fullname, &user.Username, &user.Email, &user.PhoneNumber, &user.CreatedAt)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func NewUserRepo(database *sql.DB) UserRepo {
	return &userRepo{db: database}
}
