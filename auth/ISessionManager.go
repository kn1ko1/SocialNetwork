package auth

import "socialnetwork/models"

type ISessionManager interface {
	Add(string, models.User) error
	Delete(string)
	DeleteUponExpiry(string)
	Get(string) (models.User, error)
	Lifetime() int
}
