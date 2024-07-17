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
	// Determine the correct base directory depending on the environment (local or Docker)
	var baseDir string
	if runningInDocker() {
		baseDir = "/Database"
	} else {
		wd, err := os.Getwd()

		log.Println("working directory wd:", wd)

		if err != nil {
			log.Fatal("Unable to get current working directory:", err)
		}
		baseDir = filepath.Join(wd, "Database")

		log.Println("baseDir:", baseDir)
	}

	// Define the path to the database file
	dbPath := filepath.Join(baseDir, "Identity.db")

	log.Println("dbPath:", dbPath)

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		log.Fatal("Unable to create database directory:", err)
	}

	// Open the SQLite database
	identityDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Unable to open identity database:", err)
	}
	defer identityDB.Close()

	// Ensure foreign keys are enabled
	if _, err := identityDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	log.Println("Connected to Identity SQLite database at:", dbPath)

	// Define the path to migrations directory
	var migrationsDir string
	if runningInDocker() {
		migrationsDir = filepath.Join(baseDir, "migrations", "identity")
	} else {
		migrationsDir = filepath.Join(baseDir, "..", "migrations", "identity")
	}

	// Adjust the migration paths if necessary
	dbURL := fmt.Sprintf("sqlite://%s", dbPath)
	migrationsURL := fmt.Sprintf("file://%s", migrationsDir)
	runMigrations(dbURL, migrationsURL)
}

func InitBusinessDatabase() {
	// Determine the correct base directory depending on the environment (local or Docker)
	var baseDir string
	if runningInDocker() {
		baseDir = "/Database"
	} else {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal("Unable to get current working directory:", err)
		}
		baseDir = filepath.Join(wd, "Database")
	}

	// Define the path to the database file
	dbPath := filepath.Join(baseDir, "Business.db")

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		log.Fatal("Unable to create database directory:", err)
	}

	// Open the SQLite database
	businessDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Unable to open business database:", err)
	}
	defer businessDB.Close()

	// Ensure foreign keys are enabled
	if _, err := businessDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	log.Println("Connected to Business SQLite database at:", dbPath)

	// Define the path to migrations directory
	var migrationsDir string
	if runningInDocker() {
		migrationsDir = filepath.Join(baseDir, "migrations", "business")
	} else {
		migrationsDir = filepath.Join(baseDir, "..", "migrations", "business")
	}

	// Adjust the migration paths if necessary
	dbURL := fmt.Sprintf("sqlite://%s", dbPath)
	migrationsURL := fmt.Sprintf("file://%s", migrationsDir)
	runMigrations(dbURL, migrationsURL)
}

// // Function to determine if running inside Docker
// func runningInDocker() bool {
// 	// Check if /proc/1/cgroup or /proc/self/cgroup has "docker" in it
// 	// This is a simple and commonly used method to detect Docker
// 	_, err := os.Stat("/.dockerenv")
// 	log.Println("Running in Docker:", err)
// 	return !os.IsNotExist(err)
// }

func runMigrations(databaseURL, migrationsDir string) {

	log.Println("migrationsDir string:", migrationsDir)
	log.Println("databaseURL string:", databaseURL)

	m, err := migrate.New(migrationsDir, databaseURL)
	if err != nil {
		log.Fatal("Error creating migrations instance:", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations:", err)
	}
}

func runningInDocker() bool {
	// Docker sets the HOSTNAME environment variable, which can be used to detect if running in Docker
	_, inDocker := os.LookupEnv("HOSTNAME")
	log.Println("Running in Docker:", inDocker)
	return inDocker
}

// func InitIdentityDatabase() {
// 	// Get the current working directory
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal("Unable to get current working directory:", err)
// 	}

// 	// Define the relative path for the database
// 	dbDir := filepath.Join(wd, "..", "Database")
// 	dbPath := filepath.Join(dbDir, "Identity.db")

// 	// Ensure the directory exists
// 	if err := os.MkdirAll(dbDir, 0755); err != nil {
// 		log.Fatal("Unable to create database directory:", err)
// 	}

// 	identityDB, err := sql.Open("sqlite3", dbPath)
// 	if err != nil {
// 		log.Fatal("Unable to open identity database:", err)
// 	}
// 	defer identityDB.Close()

// 	// Ensure the database file is created
// 	if _, err := identityDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
// 		log.Fatal("Error initializing database:", err)
// 	}

// 	log.Println("Connected to Identity SQLite database at:", dbPath)

// 	// Adjust the migration paths if necessary
// 	dbURL := fmt.Sprintf("sqlite://%s", dbPath)
// 	migrationsDir := fmt.Sprintf("file://%s", filepath.Join(dbDir, "migrations", "identity"))
// 	runMigrations(dbURL, migrationsDir)
// }

// func InitBusinessDatabase() {
// 	// Get the current working directory
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal("Unable to get current working directory:", err)
// 	}

// 	// Define the relative path for the database
// 	dbDir := filepath.Join(wd, "..", "Database")
// 	dbPath := filepath.Join(dbDir, "Business.db")

// 	// Ensure the directory exists
// 	if err := os.MkdirAll(dbDir, 0755); err != nil {
// 		log.Fatal("Unable to create database directory:", err)
// 	}

// 	businessDB, err := sql.Open("sqlite3", dbPath)
// 	if err != nil {
// 		log.Fatal("Unable to open business database:", err)
// 	}
// 	defer businessDB.Close()

// 	// Ensure the database file is created
// 	if _, err := businessDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
// 		log.Fatal("Error initializing database:", err)
// 	}

// 	log.Println("Connected to Identity SQLite database at:", dbPath)

// 	// Adjust the migration paths if necessary
// 	dbURL := fmt.Sprintf("sqlite://%s", dbPath)
// 	migrationsDir := fmt.Sprintf("file://%s", filepath.Join(dbDir, "migrations", "business"))
// 	runMigrations(dbURL, migrationsDir)
// }

// // func InitBusinessDatabase() {
// // 	businessDB, err := sql.Open("sqlite3", "Business.db")
// // 	if err != nil {
// // 		log.Fatal("Unable to open business database:", err)
// // 	}
// // 	defer businessDB.Close()
// // 	log.Println("Connected to Business SQLite database")

// // 	runMigrations("sqlite://../sqlite/data/Business.db", "file://../sqlite/migrations/business")
// // }

// func runMigrations(databaseURL, migrationsDir string) {

// 	log.Println("migrationsDir string", migrationsDir)
// 	log.Println("databaseURL string", databaseURL)

// 	m, err := migrate.New(migrationsDir, databaseURL)
// 	if err != nil {
// 		log.Fatal("Error creating migrations instance:", err)
// 	}
// 	defer m.Close()

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal("Error applying migrations:", err)
// 	}
// }
