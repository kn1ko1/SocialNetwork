package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

func GetUsernameByUserId(identityDb *sql.DB, userId int) (transport.UserTransport, error) {

	var userTransport transport.UserTransport
	userTransport.UserId = userId

	err := identityDb.QueryRow("SELECT Username FROM USERS WHERE UserId = ?", userId).
		Scan(
			&userTransport.Username,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("user not found", err)
		return userTransport, errors.New("user not found")
	case err != nil:
		utils.HandleError("Error retrieving user by userId", err)
		return userTransport, err
	}

	return userTransport, nil
}
