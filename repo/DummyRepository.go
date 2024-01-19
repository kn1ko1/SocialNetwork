package repo

import (
	"database/sql"
	"log"
	"socialnetwork/models"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriver       = "sqlite3"
	identityDbPath = "./sqlite/data/Identity.db"
	businessDbPath = "./sqlite/data/Business.db"
)

// Field commented for implementation example reasons:
// Should take as a field a private DB instance
//
// The repo instance is then responsible for maintaining
// the concrete DB connection
type DummyRepository struct {
	identityDb *sql.DB
	businessDb *sql.DB
}

// Constructor function
func NewDummyRepository() *DummyRepository {
	// The DB field would be constructed properly here!
	identityDb, err := sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Fatal(err)
	}
	businessDb, err := sql.Open(dbDriver, businessDbPath)
	if err != nil {
		log.Fatal(err)
	}
	return &DummyRepository{identityDb: identityDb, businessDb: businessDb}
}

func (r *DummyRepository) CreatePost(p models.Post) (models.Post, error) {
	// Call the SQLite specific instance, for example:
	//
	// return sqlite.CreatePost(r.db, p)

	// Since it is ommitted for simplicity, let's simulate it:
	p.UserId = 1
	return p, nil
}
