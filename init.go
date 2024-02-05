package main

import "socialnetwork/sqlite"

func init() {

	sqlite.InitIdentityDatabase()
	sqlite.InitBusinessDatabase()
}
