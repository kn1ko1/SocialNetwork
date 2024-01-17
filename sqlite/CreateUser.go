package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds user into the given database
func CreateUser(database *sql.DB, User *models.User) (*models.User, error) {

	query := "INSERT INTO USERS (" +
		"Bio, " +
		"CreatedAt, " +
		"DOB, " +
		"Email, " +
		"EncryptedPassword, " +
		"FirstName, " +
		"ImageUrl, " +
		"IsPublic, " +
		"LastName, " +
		"UpdatedAt, " +
		"Username" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	statement, err := database.Prepare(query)
	if err != nil {
		return User, err
	}
	res, err := statement.Exec(query,
		User.Bio,
		User.CreatedAt,
		User.DOB,
		User.Email,
		User.EncryptedPassword,
		User.FirstName,
		User.ImageUrl,
		User.IsPublic,
		User.LastName,
		User.UpdatedAt,
		User.Username)
	if err != nil {
		return User, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return User, err
	}
	User.UserId = int(id)
	return User, nil
}
