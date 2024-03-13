package sqlite

import (
	"database/sql"
	"log"
	"socialnetwork/models"
	"socialnetwork/utils"
	"time"
)

// Adds user into the gdatabase *sql.DBiven database
func CreateUser(database *sql.DB, User models.User) (models.User, error) {

	query := `INSERT INTO USERS (
		Bio,
		CreatedAt, 
		DOB,
		Email,
		EncryptedPassword,
		FirstName,
		ImageURL,
		IsPublic,
		LastName,
		UpdatedAt,
		Username
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
		return User, err
	}
	res, err := statement.Exec(
		User.Bio,
		User.CreatedAt,
		User.DOB,
		User.Email,
		User.EncryptedPassword,
		User.FirstName,
		User.ImageURL,
		User.IsPublic,
		User.LastName,
		User.UpdatedAt,
		User.Username)
	if err != nil {
		utils.HandleError("Error executing statement.", err)

		return User, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return User, err
	}
	t := time.Unix(User.DOB/1000, 0)
	log.Println("[sqlite/CreateUser] DOB converted back ", t.Format("02-01-2006"))
	User.UserId = int(id)
	return User, nil
}
