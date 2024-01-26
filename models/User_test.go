package models

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Test email addresses
var (
	sutBio = []string{"string 1", "string 2", ""}
	sutDOB = []int64{int64((time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC)).Unix()),
		int64((time.Date(1975, time.May, 15, 0, 0, 0, 0, time.UTC)).Unix()),
		int64((time.Date(1990, time.October, 30, 0, 0, 0, 0, time.UTC)).Unix())}
	sutEmail     = []string{"john.doe@example.com", "another@email.co", "user@123.de"}
	sutFirstName = []string{"Matthew", "Micheal", "Nikoi"}
	sutIsPublic  = []bool{true, false}
	sutLastName  = []string{"Cheetham", "Cornea", "Fenton"}
	sutPassword  = []string{"Password123", "Password456", "Password789"}
	sutUsername  = []string{"User123", "User456", "User789"}
)

func TestUserValidateInvalidFieldExpectError(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, generateInvalidUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err == nil {
				t.Error("expect error for invalid event field")
			}
		})
	}
}

func TestUserValidateMissingFieldExpectError(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, generateMissingFieldUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err == nil {
				t.Error("expect error for missing event field")
			}
		})
	}
}

func TestUserValidateValidExpectNil(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, generateValidUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err != nil {
				t.Error("expect nil for valid event")
			}
		})
	}
}

func generateValidUser() *User {
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
	return u
}

func generateMissingFieldUser() *User {
	u := generateValidUser()
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

func generateInvalidUser() *User {
	u := generateValidUser()
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
