package sqlite

import (
	"database/sql"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Retrieves limited user info, for user list on homepage,  from the USERS table
func GetAllUsersTransport(database *sql.DB) ([]transport.UserTransport, error) {
	rows, err := database.Query("SELECT UserId, Username FROM USERS ORDER BY Username ASC")
	if err != nil {
		utils.HandleError("Error executing SELECT UserId, Username FROM USERS statement.", err)
		return nil, err
	}
	defer rows.Close()

	var users []transport.UserTransport

	for rows.Next() {
		var user transport.UserTransport
		err := rows.Scan(
			&user.UserId,
			&user.Username,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetAllUsersTransport.", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetAllUsersTransport.", err)
		return nil, err
	}

	return users, nil
}
