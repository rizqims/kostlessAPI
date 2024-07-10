package repository

import (
	"database/sql"
	//"errors"
	"kostless/model"
	//"kostless-api/util"
	"time"
)

// interface
type UserRepo interface {
	CreatedNewUser(payload model.User) (model.User, error)
	GetByUsername(username string) (model.User, error)
	GetUserById(id string) (model.User, error)
	PutUpdateUserProf(id string, user model.User) error
}

// struct
type userRepo struct {
	db *sql.DB
}

// GetUserById implements UserRepo.
func (u *userRepo) GetUserById(id string) (model.User, error) {

	var user model.User
	err := u.db.QueryRow(`SELECT id, fullname, username, password, email, phone_number, photo_profile, created_at, updated_at FROM users WHERE id=$1`, id).Scan(&user.Id, &user.Fullname, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// PutUpdateUserProf implements UserRepo.
func (u *userRepo) PutUpdateUserProf(id string, user model.User) error {
	query := `UPDATE users SET fullname=$1, username=$2, password=$3, email=$4, phone_number=$5, photo_profile=$6, updated_at=$7 WHERE id=$8`
	_, err := u.db.Exec(query, user.Fullname, user.Username, user.Password, user.Email, user.PhoneNumber, user.PhotoProfile, time.Now(), user.Id)
	return err
}

// GetByUsername implements UserRepo.
func (u *userRepo) GetByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, fullname, username, password, email, phone_number, photo_profile, created_at, updated_at FROM users WHERE username=$1`, username).Scan(&user.Id, &user.Fullname, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// CreatedNewUser implements UserRepo
func (u *userRepo) CreatedNewUser(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`INSERT INTO users (fullname, username, password, email, phone_number, photo_profile, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id, fullname, username, email, phone_number, photo_profile,created_at`, payload.Fullname, payload.Username, payload.Password, payload.Email, payload.PhoneNumber, payload.PhotoProfile, time.Now()).Scan(&user.Id, &user.Fullname, &user.Username, &user.Email, &user.PhoneNumber, &user.PhotoProfile, &user.CreatedAt)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func NewUserRepo(database *sql.DB) UserRepo {
	return &userRepo{db: database}
}
