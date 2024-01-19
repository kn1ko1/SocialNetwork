package models

import (
	"errors"
)

const (
	minUsernameLength = 4
	maxUsernameLength = 50
)

type User struct {
	UserId            int
	Bio               string
	CreatedAt         int64
	DOB               int64
	Email             string
	EncryptedPassword string
	FirstName         string
	ImageUrl          string
	IsPublic          bool
	LastName          string
	UpdatedAt         int64
	Username          string
}

func (u *User) Validate() error {
	// All of the stuff
	if len(u.Username) < minUsernameLength || len(u.Username) > maxUsernameLength {
		return errors.New("invalid length of username")
	}
	return nil
}
