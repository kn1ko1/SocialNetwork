export function GroupDetailsUserList({ userList, groupId, groupMembers, AddGroupUser }) {
    return (
        <div className="userListGroup">
            <h2>UserList</h2>
            {userList !== null && userList.length !== groupMembers.length ? (
                userList
                    .filter(user => !groupMembers.some(member => member.userId === user.userId))
                    .map((user, index) => (
                        <div key={index}>
                            <span>{user.username}</span>
                            <button onClick={() => AddGroupUser(user.userId, groupId, "groupInvite")}>Add to Group</button>
                        </div>
                    ))
            ) : (
                <p>Look's like everyone's part of this group!</p>
            )}
        </div>
    );
}
