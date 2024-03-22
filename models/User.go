package models

import (
	"errors"
	"math/rand"
	"regexp"
	"time"
)

type User struct {
	UserId            int    `json:"userId"`
	Bio               string `json:"bio"`
	CreatedAt         int64  `json:"createdAt"`
	DOB               int64  `json:"dob"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encryptedPassword"`
	FirstName         string `json:"firstName"`
	ImageURL          string `json:"imageURL"`
	IsPublic          bool   `json:"isPublic"`
	LastName          string `json:"lastName"`
	UpdatedAt         int64  `json:"updatedAt"`
	Username          string `json:"username"`
}

const (
	minUsernameLength = 4
	maxUsernameLength = 50
	emailRegex        = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

var (
	minDOB = int64((time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC)).UnixMilli())
	maxDOB = int64((time.Now()).UnixMilli())
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

func GenerateValidUser() *User {
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
		ImageURL:          sutImageURL[idxImageURL],
		IsPublic:          sutIsPublic[idxIsPublic],
		LastName:          sutLastName[idxLastName],
		UpdatedAt:         ctime,
		Username:          sutUsername[idxUsername],
	}
	return u
}

func GenerateMissingFieldUser() *User {
	u := GenerateValidUser()
	missingField := rand.Intn(7)
	switch missingField {
	case 0:
		u.CreatedAt = 0
	case 1:
		u.Email = ""
	case 2:
		u.EncryptedPassword = ""
	case 3:
		u.FirstName = ""
	case 4:
		u.LastName = ""
	case 5:
		u.UpdatedAt = 0
	case 6:
		u.Username = ""
	}
	return u
}

func GenerateInvalidUser() *User {
	u := GenerateValidUser()
	invalidField := rand.Intn(8)
	switch invalidField {
	case 0:
		u.CreatedAt = -u.CreatedAt
	case 1:
		u.DOB = -631238400
	case 2:
		u.Email = ""
	case 3:
		u.EncryptedPassword = ""
	case 4:
		u.FirstName = ""
	case 5:
		u.LastName = ""
	case 6:
		u.UpdatedAt = -u.UpdatedAt
	case 7:
		u.Username = ""
	}
	return u
}
