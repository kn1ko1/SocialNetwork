package ws

var (
	manager *GroupManager
)

func init() {
	manager = NewGroupManager()
	manager.Start()
}
