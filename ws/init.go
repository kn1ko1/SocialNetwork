package ws

var (
	socketGroupManager *SocketGroupManager
)

func init() {
	socketGroupManager = NewSocketGroupManager()
	socketGroupManager.Start()
}
