package models

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