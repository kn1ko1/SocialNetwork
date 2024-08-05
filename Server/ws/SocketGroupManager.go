package ws

type SocketGroupManager struct {
	SocketGroups map[int]*SocketGroup
}

func NewSocketGroupManager() *SocketGroupManager {
	ret := new(SocketGroupManager)
	ret.SocketGroups = make(map[int]*SocketGroup)
	return ret
}

func (m *SocketGroupManager) Start() {
	defaultSocketGroup := NewSocketGroup(0)
	m.SocketGroups[defaultSocketGroup.SocketGroupID] = defaultSocketGroup
	go defaultSocketGroup.Run()
}
