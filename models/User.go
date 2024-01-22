package models

import (
	"errors"
	"regexp"
	"time"
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

const (
	minUsernameLength = 4
	maxUsernameLength = 50
	emailRegex        = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

var (
	minDOB = int64((time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC)).Unix())
	maxDOB = int64((time.Now()).Unix())
	re     = regexp.MustCompile(emailRegex)
)

func (u *User) Validate() error {
	if u.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if u.DOB < minDOB || u.DOB > maxDOB {
		return errors.New("invalid DOB field")
	}
	if !re.MatchString(u.Email) {
		return errors.New("invalid Email field")
	}
	if u.EncryptedPassword == "" {
		return errors.New("cEncryptedPassword must not be empty")
	}
	if u.FirstName == "" {
		return errors.New("first name must not be empty")
	}
	if u.LastName == "" {
		return errors.New("last name must not be empty")
	}
	if u.UpdatedAt < u.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field. cannot be before 'CreatedAt' field")
	}
	if len(u.Username) < minUsernameLength || len(u.Username) > maxUsernameLength {
		return errors.New("invalid length of username")
	}
	return nil
}
