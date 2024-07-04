package database

import (
	"socialnetwork/Server/ws"
)

var (
	socketGroupManager *ws.SocketGroupManager
)

func init() {

	InitIdentityDatabase()
	InitBusinessDatabase()

	socketGroupManager = ws.NewSocketGroupManager()
	socketGroupManager.Start()
}
