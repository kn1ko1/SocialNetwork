package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"socialnetwork/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func InitIdentityDatabase() {

	dirPath := "./sqlite/migrations/identity"

	// Open the directory
	dir, direrr := os.Open(dirPath)
	if direrr != nil {
		fmt.Println("Error opening directory:", direrr)
		return
	}
	defer dir.Close()

	// Read the files in the directory
	fileInfos, fileInfoserr := dir.Readdir(-1)
	if fileInfoserr != nil {
		fmt.Println("Error reading directory:", fileInfoserr)
		return
	}

	// Print the names of the files
	fmt.Println("Files in", dirPath+":")
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
	// Initialization for identity.db
	var err error

	identityDB, err := sql.Open("sqlite3", "./sqlite/data/Identity.db")
	if err != nil {
		utils.HandleError("Unable to open identity database", err)
	}

	log.Println("Connected to Identity SQLite database")

	// Apply "up" migrations from SQL files for identity.db
	if err := runMigrations(identityDB, dirPath, "up"); err != nil {
		utils.HandleError("Error applying 'up' migrations for identity.db: ", err)
	}

	defer identityDB.Close()
}

func InitBusinessDatabase() {
	// Initialization for business.db
	var err error

	businessDB, err := sql.Open("sqlite3", "./sqlite/data/Business.db")
	if err != nil {
		utils.HandleError("Unable to open business database", err)
	}

	log.Println("Connected to Business SQLite database")

	// Apply "up" migrations from SQL files for business.db
	if err := runMigrations(businessDB, "./sqlite/migrations/business", "up"); err != nil {
		utils.HandleError("Error applying 'up' migrations for business.db: ", err)
	}

	defer businessDB.Close()
}

func runMigrations(database *sql.DB, migrationDir, direction string) error {
	// Create SQLite driver instance
	// driver, err := sqlite.WithInstance(database, &sqlite.Config{})
	// if err != nil {
	// 	return err
	// }

	// Create migrate instance
	m, err := migrate.New(
		migrationDir,
		"sqlite/data/Identity.db",
	)
	if err != nil {
		log.Println("here")
		return err
	}
	log.Println("fsdf")
	// Perform migration based on direction
	var migrationErr error
	if direction == "up" {
		migrationErr = m.Up()
	} else if direction == "down" {
		migrationErr = m.Steps(-1)
	} else {
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	// Handle migration error
	if migrationErr != nil && migrationErr != migrate.ErrNoChange {
		return migrationErr
	}

	// Log completion
	log.Print("Migrations completed successfully")

	return nil
}
