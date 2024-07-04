package transport

import "socialnetwork/Server/models"

type RegisteringUser struct {
	models.User
	DOB string `json:"dob"`
}
