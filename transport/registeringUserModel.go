package transport

import "socialnetwork/models"

type RegisteringUser struct {
	models.User
	DOB string `json:"dob"`
}
