package dbUtils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

//run `migrate --help` in terminal to explore migrate package.

func InitIdentityDatabase() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current working directory:", err)
	}
	log.Println("Current working directory:", wd)

	dbPath := filepath.Join(wd, "Identity.db")
	identityDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Unable to open identity database:", err)
	}
	defer identityDB.Close()

	// Ensure the database file is created
	if _, err := identityDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	log.Println("Connected to Identity SQLite database at:", dbPath)

	// Adjust the migration paths if necessary
	dbURL := fmt.Sprintf("sqlite://%s", dbPath)
	migrationsDir := "file://sqlite/migrations/identity"
	runMigrations(dbURL, migrationsDir)
}

func InitBusinessDatabase() {
	businessDB, err := sql.Open("sqlite3", "Business.db")
	if err != nil {
		log.Fatal("Unable to open business database:", err)
	}
	defer businessDB.Close()
	log.Println("Connected to Business SQLite database")

	runMigrations("sqlite://../sqlite/data/Business.db", "file://../sqlite/migrations/business")
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
