package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"socialnetwork/utils"

	_ "github.com/mattn/go-sqlite3"
)

//var Database *sql.DB

func InitIdentityDatabase() {
	// Initialization for identity.db
	var err error
	identityDatabasePath := "./sqlite/data/Identity.db"

	identityDB, err := sql.Open("sqlite3", identityDatabasePath)
	if err != nil {
		utils.HandleError("Unable to open identity database", err)
	}

	log.Println("Connected to Identity SQLite database")

	// Apply "up" migrations from SQL files for identity.db
	RunMigrations(identityDB, "./sqlite/migrations/identity", "up")
	if err != nil {
		utils.HandleError("Error applying 'up' migrations for identity.db: ", err)
	}

	WipeDatabaseOnCommandNew(identityDB, identityDatabasePath)

	defer identityDB.Close()
}

func InitBusinessDatabase() {
	// Initialization for business.db
	var err error
	businessDatabasePath := "./sqlite/data/Business.db"

	businessDB, err := sql.Open("sqlite3", businessDatabasePath)
	if err != nil {
		utils.HandleError("Unable to open business database", err)
	}

	log.Println("Connected to Business SQLite database")

	// Apply "up" migrations from SQL files for business.db
	RunMigrations(businessDB, "./sqlite/migrations/business", "up")
	if err != nil {
		utils.HandleError("Error applying 'up' migrations for business.db: ", err)
	}

	WipeDatabaseOnCommandNew(businessDB, businessDatabasePath)

	defer businessDB.Close()
}

func RunMigrations(Database *sql.DB, migrationDir, direction string) {
	// Read files in the migration directory
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		utils.HandleError("Error reading migration directory", err)
		return
	}

	// Iterate through migration files
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		// Check if it's an "up" migration file
		if direction == "up" && !isUpMigration(fileName) {
			continue
		}
		// Check if it's a "down" migration file
		if direction == "down" && !isDownMigration(fileName) {
			continue
		}

		// Build the full path to the migration file
		migrationPath := migrationDir + "/" + fileName

		// Read SQL content from the migration file
		sqlBytes, err := os.ReadFile(migrationPath)
		if err != nil {
			message := fmt.Sprintf("error reading migration file %s", migrationPath)
			utils.HandleError(message, err)
		}

		// Execute the SQL content on the database
		_, err = Database.Exec(string(sqlBytes))
		if err != nil {
			message := fmt.Sprintf("error executing migration %s:", migrationPath)
			utils.HandleError(message, err)
		}
	}
}

func isUpMigration(fileName string) bool {
	return len(fileName) > 3 && fileName[len(fileName)-7:] == "_up.sql"
}

func isDownMigration(fileName string) bool {
	return len(fileName) > 5 && fileName[len(fileName)-9:] == "_down.sql"
}

// This function will delete the database if "go run . new" is typed in command line.
func WipeDatabaseOnCommandNew(database *sql.DB, path string) {
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			// Rollback the last migration (uncomment if needed)
			RunMigrations(database, path, "down")
			fmt.Println("Dropped all tables")
		}
	}
}
