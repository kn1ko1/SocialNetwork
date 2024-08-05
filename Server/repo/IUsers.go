package repo

import "socialnetwork/Server/models"

type IUsers interface {
	// User
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUsersByPublic() ([]models.User, error)
	GetUserById(userId int) (models.User, error)
	// GetUserByEmail(email string) (models.User, error)
	// GetUserByUsername(username string) (models.User, error)
	GetUserByUsernameOrEmail(usernameOrEmail string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUserById(userId int) error
	DeleteAllUsers() error
}
