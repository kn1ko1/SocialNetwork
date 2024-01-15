package repo

import (
	"database/sql"
	"log"
	"socialnetwork/models"

	_ "github.com/mattn/go-sqlite3"
)

// Field commented for implementation example reasons:
// Should take as a field a private DB instance
//
// The repo instance is then responsible for maintaining
// the concrete DB connection
type DummyRepository struct {
	db *sql.DB
}

// Constructor function
func NewDummyRepository() *DummyRepository {
	// The DB field would be constructed properly here!
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	return &DummyRepository{db: db}
}

func (r *DummyRepository) CreatePost(p models.Post) (models.Post, error) {
	// Call the SQLite specific instance, for example:
	//
	// return sqlite.CreatePost(r.db, p)

	// Since it is ommitted for simplicity, let's simulate it:
	p.UserId = 1
	return p, nil
}
