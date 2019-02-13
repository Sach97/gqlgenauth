package model

import (
	"golang.org/x/crypto/bcrypt"
)

//User struct represents our User model
type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt string `db:"created_at"`
	Confirmed bool
	// EmailConfirmedAt string `db:"email_confirmed_at"`
	// EmailSentAt string `db:"email_confirmed_at"`
	Roles []*Role
}

//HashedPassword hash user password
func (user *User) HashedPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

//ComparePassword compare the given password with the password in db
func (user *User) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
