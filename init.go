package main

import (
	dbUtils "socialnetwork/Database/databaseUtils"
	"socialnetwork/Server/ws"
)

var (
	socketGroupManager *ws.SocketGroupManager
)

func init() {

	dbUtils.InitIdentityDatabase()
	dbUtils.InitBusinessDatabase()

	socketGroupManager = ws.NewSocketGroupManager()
	socketGroupManager.Start()
}
