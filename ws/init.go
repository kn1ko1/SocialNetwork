package ws

var (
	groupManager *GroupManager
	eventManager *EventManager
)

func init() {
	groupManager = NewGroupManager()
	groupManager.Start()

	eventManager = NewEventManager()
	eventManager.Start()
}
