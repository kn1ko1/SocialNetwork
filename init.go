package main

import (
	"socialnetwork/utils"

	"socialnetwork/ws"
)

var (
	socketGroupManager *ws.SocketGroupManager
)

func init() {

	utils.InitIdentityDatabase()
	utils.InitBusinessDatabase()

	socketGroupManager = ws.NewSocketGroupManager()
	socketGroupManager.Start()
}
