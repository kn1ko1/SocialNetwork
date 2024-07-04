export function GroupDetailsUserList({
  userList,
  groupId,
  groupMembers,
  AddGroupUser
}) {
  return /*#__PURE__*/React.createElement("div", {
    className: "userListGroup"
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline'
    }
  }, "Invite User"), userList !== null && userList.length !== groupMembers.length ? userList.filter(user => !groupMembers.some(member => member.userId === user.userId)).map((user, index) => /*#__PURE__*/React.createElement("div", {
    key: index,
    style: {
      padding: "10px"
    }
  }, /*#__PURE__*/React.createElement("span", null, user.username), /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-primary",
    style: {
      marginLeft: "10px"
    },
    onClick: () => AddGroupUser(user.userId, groupId, "groupInvite")
  }, "Add to Group"))) : /*#__PURE__*/React.createElement("p", null, "Look's like everyone's part of this group!"));
}