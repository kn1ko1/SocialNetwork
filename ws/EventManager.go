package ws

type EventManager struct {
	Events map[int]*Event
}

func NewEventManager() *EventManager {
	ret := new(EventManager)
	ret.Events = make(map[int]*Event)
	return ret
}

func (m *EventManager) Start() {
	defaultEvent := NewEvent(0)
	m.Events[defaultEvent.EventID] = defaultEvent
	go defaultEvent.Run()
}
