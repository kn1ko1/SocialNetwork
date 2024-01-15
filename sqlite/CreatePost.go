package sqlite

import (
	"database/sql"
	"matthewhope/example-architecture/models"
)

func CreatePost(db *sql.DB, p models.Post) (models.Post, error) {
	// not implemented
	//
	// do all the SQL stuff to add the Post to the db & get a result
	return p, nil
}
