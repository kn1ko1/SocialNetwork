package ws

type GroupManager struct {
	Groups map[int]*Group
}

func NewGroupManager() *GroupManager {
	ret := new(GroupManager)
	ret.Groups = make(map[int]*Group)
	return ret
}

func (m *GroupManager) Start() {
	defaultGroup := NewGroup(0)
	m.Groups[defaultGroup.GroupID] = defaultGroup
	go defaultGroup.Run()
}
