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
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current working directory:", err)
	}

	// Define the relative path for the database
	dbDir := filepath.Join(wd, "..", "Database")
	dbPath := filepath.Join(dbDir, "Identity.db")

	// Ensure the directory exists
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatal("Unable to create database directory:", err)
	}

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
	migrationsDir := fmt.Sprintf("file://%s", filepath.Join(dbDir, "migrations", "identity"))
	runMigrations(dbURL, migrationsDir)
}

func InitBusinessDatabase() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current working directory:", err)
	}

	// Define the relative path for the database
	dbDir := filepath.Join(wd, "..", "Database")
	dbPath := filepath.Join(dbDir, "Business.db")

	// Ensure the directory exists
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatal("Unable to create database directory:", err)
	}

	businessDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Unable to open business database:", err)
	}
	defer businessDB.Close()

	// Ensure the database file is created
	if _, err := businessDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	log.Println("Connected to Business SQLite database at:", dbPath)

	// Adjust the migration paths if necessary
	dbURL := fmt.Sprintf("sqlite://%s", dbPath)
	migrationsDir := fmt.Sprintf("file://%s", filepath.Join(dbDir, "migrations", "business"))
	runMigrations(dbURL, migrationsDir)
}

func runMigrations(databaseURL, migrationsDir string) {

	log.Println("testing path for Docker migrationsDir", migrationsDir)
	log.Println("testing path for Docker databaseURL", databaseURL)

	m, err := migrate.New(migrationsDir, databaseURL)
	if err != nil {
		log.Fatal("Error creating migrations instance:", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations:", err)
	}
	log.Println("Applying Up migrations from", migrationsDir)
}
