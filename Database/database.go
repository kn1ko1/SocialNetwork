package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

//run `migrate --help` in terminal to explore migrate package.

func InitIdentityDatabase() {
	identityDB, err := sql.Open("sqlite3", "./sqlite/data/Identity.db")
	if err != nil {
		log.Fatal("Unable to open identity database:", err)
	}
	defer identityDB.Close()

	log.Println("Connected to Identity SQLite database")

	runMigrations("sqlite://./sqlite/data/Identity.db", "file://./sqlite/migrations/identity")

}

func InitBusinessDatabase() {
	businessDB, err := sql.Open("sqlite3", "./sqlite/data/Business.db")
	if err != nil {
		log.Fatal("Unable to open business database:", err)
	}
	defer businessDB.Close()

	log.Println("Connected to Business SQLite database")

	runMigrations("sqlite://./sqlite/data/Business.db", "file://./sqlite/migrations/business")

}

func runMigrations(databaseURL, migrationsDir string) {

	m, err := migrate.New(migrationsDir, databaseURL)

	if err != nil {
		log.Fatal("Error creating migrations instance:", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations:", err)
	}
}
