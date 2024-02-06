package main

import (
	"os"
	"socialnetwork/sqlite"
	"socialnetwork/utils"
)

func init() {

	identityDB, identityDBErr := sqlite.InitIdentityDatabase()
	if identityDBErr != nil {
		utils.HandleError("Unable to open identity database", identityDBErr)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			sqlite.WipeDatabaseOnCommandNew(identityDB, "sqlite/migrations/identity")

		}
	} else {
		// Apply "up" migrations from SQL files for identity.db
		sqlite.RunMigrations(identityDB, "./sqlite/migrations/identity", "up")
	}

	businessDB, businessDBErr := sqlite.InitBusinessDatabase()
	if businessDBErr != nil {
		utils.HandleError("Unable to open business database", businessDBErr)
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "new" {
			sqlite.WipeDatabaseOnCommandNew(businessDB, "sqlite/migrations/business")

		}
	} else {
		// Apply "up" migrations from SQL files for business.db
		sqlite.RunMigrations(businessDB, "./sqlite/migrations/business", "up")
	}

}
