package sqlite

import (
	"database/sql"
	events "socialnetwork/sqlite/EVENTS"
	group_users "socialnetwork/sqlite/GROUP_USERS"
	messages "socialnetwork/sqlite/MESSAGES"
	posts "socialnetwork/sqlite/POSTS"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Retrieves data for the user's homepage including posts and comments
func GetGroupDataForUser(identityDB, businessDb *sql.DB, groupId int) (transport.GroupData, error) {
	var GroupData transport.GroupData
	var err error
	// Get all users, for inviting purposes
	GroupData.AllUsers, err = GetAllUsersTransport(identityDB)
	if err != nil {
		utils.HandleError("Error in GetAllUsersTransport in GetGroupDataForUser", err)

	}

	// Gets a list of all users in the group
	GroupData.GroupUsers, err = group_users.GetGroupUsersByGroupId(identityDB, groupId)
	if err != nil {
		utils.HandleError("Error in GetGroupUsersByGroupId in GetGroupDataForUser", err)

	}

	// for _, groupUser := range groupUsers {
	// 	for _, user := range GroupData.AllUsers {
	// 		if user.UserId == groupUser.UserId {
	// 			GroupData.GroupUsersWithUsernames = append(GroupData.GroupUsersWithUsernames, user)
	// 			break
	// 		}
	// 	}
	// }

	// for i, groupUser := range groupUsers {
	// 	for _, user := range GroupData.AllUsers {
	// 		if user.UserId == groupUser.UserId {
	// 			GroupData.GroupUsersWithUsernames[i].IsMember = true
	// 			break
	// 		}
	// 	}
	// }

	// Get posts with for this group
	GroupData.GroupPosts, err = posts.GetPostsByGroupId(businessDb, groupId)
	if err != nil {
		utils.HandleError("Error in GetPostsByGroupId in GetGroupDataForUser", err)
	}

	// Get groups chat log
	GroupData.GroupMessages, err = messages.GetMessagesByMessageTypeandTargetId(businessDb, "group", groupId)
	if err != nil {
		utils.HandleError("Error in GetMessagesByMessageTypeandTargetId in GetGroupDataForUser", err)
	}

	// GetPostsPrivate retrieves private posts for the given followerId
	GroupData.GroupEvents, err = events.GetEventsByGroupId(businessDb, groupId)
	if err != nil {
		utils.HandleError("Error in GetEventsByGroupId in GetGroupDataForUser", err)
	}

	return GroupData, nil
}
