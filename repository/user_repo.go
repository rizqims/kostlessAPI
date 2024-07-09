package repository

import (
	"database/sql"
	"kostless-api/model"
	"time"
)

// interface
type UserRepo interface {
	CreatedNewUser(payload model.User) (model.User, error)
	GetByUsername(username string) (model.User, error)
	GetUserById(id string) (model.User, error)
	PutUpdateUserProf(user model.User) error
}

// struct
type userRepo struct {
	db *sql.DB
}

// GetUserById implements UserRepo.
func (u *userRepo) GetUserById(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, fullname, username, password, email, phone_number, created_at, updated_at FROM users WHERE id=$1`, id).Scan(&user.Id, &user.Fullname, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// PutUpdateUserProf implements UserRepo.
func (u *userRepo) PutUpdateUserProf(user model.User) error {
	query := `UPDATE users SET fullname=$1, username=$2, password=$3, email=$4, phone_number=$5, updated_at=$6 WHERE id=$7`
    _, err := u.db.Exec(query, user.Fullname, user.Username, user.Password, user.Email, user.PhoneNumber, user.UpdatedAt, user.Id)
    return err
}

// GetByUsername implements UserRepo.
func (u *userRepo) GetByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, fullname, username, password, email, phone_number, created_at, updated_at FROM users WHERE username=$1`, username).Scan(&user.Id, &user.Fullname, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// CreatedNewUser implements UserRepo
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
