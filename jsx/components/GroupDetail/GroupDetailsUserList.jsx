export function GroupDetailsUserList({ userList, groupId, groupMembers, AddGroupUser }) {
    return (
        <div className="userListGroup">
            <h2 style={{textDecoration: 'underline'}}>User List</h2>
            {userList !== null && userList.length !== groupMembers.length ? (
                userList
                    .filter(user => !groupMembers.some(member => member.userId === user.userId))
                    .map((user, index) => (
                        <div key={index}>
                            <span>{user.username}</span>
                            <button type="button" className="btn btn-primary" style={{ marginLeft: "10px" }} onClick={() => AddGroupUser(user.userId, groupId, "groupInvite")}>Add to Group</button>
                        </div>
                    ))
            ) : (
                <p>Look's like everyone's part of this group!</p>
            )}
        </div>
    );
}
