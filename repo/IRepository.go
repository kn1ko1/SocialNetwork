// Message from Matt
// Please do not touch this file

package repo

import (
	"socialnetwork/transport"
)

type IRepository interface {
	// Model Tables
	IComments
	IEvents
	IGroups
	IMessages
	INotifications
	IPosts
	IUsers
	// Link Tables
	IEventUsers
	IGroupUsers
	IPostUsers
	IUserUsers

	// Need to be accurately re-defined as composition of DB functions OUTSIDE
	// of the Repo interface - Repo interface represents only retrieval of DB data
	// Not Transformation - (see ORMs)

	// Users
	GetAllUsersTransport() ([]transport.UserTransport, error)
	// Home (name tbc)
	GetHomeDataForUser(userId int) (transport.HomeModel, error)

	//Profile
	GetProfileDataForUser(userId int) (transport.ProfileModel, error)
	UpdateIsPublic(userId int, isPublic bool) error

	//Group
	//GetGroupDataForUser(groupId int) (transport.GroupData, error)
}
