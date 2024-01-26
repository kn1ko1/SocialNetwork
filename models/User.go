package models

import (
	"errors"
	"math/rand"
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

func GenerateValidUser(withId bool) *User {
	idxBio := rand.Intn(len(sutBio))
	idxDOB := rand.Intn(len(sutDOB))
	ctime := rand.Int63n(1000) + 1
	idxEmail := rand.Intn(len(sutEmail))
	idxPassword := rand.Intn(len(sutPassword))
	idxFirstName := rand.Intn(len(sutFirstName))
	idxImageURL := rand.Intn(len(sutImageURL))
	idxIsPublic := rand.Intn(len(sutIsPublic))
	idxLastName := rand.Intn(len(sutLastName))
	idxUsername := rand.Intn(len(sutUsername))

	u := &User{
		Bio:               sutBio[idxBio],
		CreatedAt:         ctime,
		DOB:               sutDOB[idxDOB],
		Email:             sutEmail[idxEmail],
		EncryptedPassword: sutPassword[idxPassword],
		FirstName:         sutFirstName[idxFirstName],
		ImageUrl:          sutImageURL[idxImageURL],
		IsPublic:          sutIsPublic[idxIsPublic],
		LastName:          sutLastName[idxLastName],
		UpdatedAt:         ctime,
		Username:          sutUsername[idxUsername],
	}
	if withId {
		u.UserId = rand.Intn(1000)
	}
	return u
}
