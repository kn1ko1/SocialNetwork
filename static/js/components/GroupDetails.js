const {
  useState,
  useEffect
} = React;
export function GroupDetails({
  group
}) {
  const [userList, setUserList] = useState([]);
  const [groupMembers, setGroupMembers] = useState([]);
  const [groupPosts, setGroupPosts] = useState([]);
  const [groupMessages, setGroupMessages] = useState([]);
  const [groupEvents, setGroupEvents] = useState([]);
  useEffect(() => {
    const fetchGroupData = async () => {
      try {
        const promises = [];
        promises.push(fetch(`http://localhost:8080/api/users/transport`));
        promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/groupUsers`));
        promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/posts`));
        promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/messages`));
        promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/events`));
        const results = await Promise.all(promises);
        const userListResponse = results[0];
        const groupMembersResponse = results[1];
        const postsResponse = results[2];
        const messagesResponse = results[3];
        const eventsResponse = results[4];
        if (!userListResponse.ok) {
          throw new Error('Failed to fetch user list');
        }
        if (!groupMembersResponse.ok) {
          throw new Error('Failed to fetch group members');
        }
        if (!postsResponse.ok) {
          throw new Error('Failed to fetch group posts');
        }
        if (!messagesResponse.ok) {
          throw new Error('Failed to fetch group messages');
        }
        if (!eventsResponse.ok) {
          throw new Error('Failed to fetch group eventsResponse');
        }
        const userListData = await userListResponse.json();
        const groupMembersData = await groupMembersResponse.json();
        const postsData = await postsResponse.json();
        const messagesData = await messagesResponse.json();
        const eventsData = await eventsResponse.json();
        setUserList(userListData);
        setGroupMembers(groupMembersData);
        setGroupPosts(postsData);
        setGroupMessages(messagesData);
        setGroupEvents(eventsData);
        console.log("This is GroupMembersData:", groupMembersData);
      } catch (error) {
        console.error('Error fetching group posts:', error);
      }
    };
    fetchGroupData();
  }, [group.groupId]);

  // const UserList = ({ userList }) => {
  const handleAddToGroup = userId => {
    console.log('Adding user to group with groupId:', group.groupId);
    console.log('User ID:', userId);
    AddGroupUser({
      groupId: group.groupId,
      userId: userId
    }); // Call AddGroupUser function with groupId and userId
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "group-details"
  }, /*#__PURE__*/React.createElement("h2", null, group.title), /*#__PURE__*/React.createElement("p", null, group.description), /*#__PURE__*/React.createElement(PostFormGroup, {
    group: group
  }), /*#__PURE__*/React.createElement("div", {
    className: "userList"
  }, /*#__PURE__*/React.createElement("h2", null, "UserList"), userList !== null && userList.length > 0 ? userList.map((user, index) => /*#__PURE__*/React.createElement("div", {
    key: index
  }, /*#__PURE__*/React.createElement("span", null, user.username), /*#__PURE__*/React.createElement("button", {
    onClick: () => handleAddToGroup(user.userId)
  }, "Add to Group"))) : /*#__PURE__*/React.createElement("p", null, "No Users?!")), /*#__PURE__*/React.createElement("div", {
    className: "groupMembers"
  }, /*#__PURE__*/React.createElement("h2", null, "Group Members"), groupMembers !== null && groupMembers.length > 0 ? groupMembers.map((member, index) => {
    // Find the user object corresponding to the member's userId
    const user = userList.find(user => user.userId === member.userId);
    return /*#__PURE__*/React.createElement("div", {
      key: index
    }, user ? user.username : 'Unknown User');
  }) : /*#__PURE__*/React.createElement("p", null, "It's just you... Maybe you should invite someone?")), /*#__PURE__*/React.createElement("div", {
    id: "groupPosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Posts"), groupPosts !== null ? groupPosts.map(post => /*#__PURE__*/React.createElement("li", {
    key: post.createdAt
  }, post.body)) : /*#__PURE__*/React.createElement("div", {
    id: "groupPosts"
  }, "There are no posts in this groups yet")), /*#__PURE__*/React.createElement("div", {
    className: "groupMessages"
  }, /*#__PURE__*/React.createElement("h2", null, "Messages"), groupMessages !== null && groupMessages.length > 0 ? groupMessages.map((message, index) => /*#__PURE__*/React.createElement("div", {
    key: index
  }, message.body)) : /*#__PURE__*/React.createElement("p", null, "No Messages")), /*#__PURE__*/React.createElement("div", {
    className: "groupEvents"
  }, /*#__PURE__*/React.createElement("h2", null, "Events"), groupEvents !== null && groupEvents.length > 0 ? groupEvents.map((event, index) => /*#__PURE__*/React.createElement("div", {
    key: index
  }, event.title)) : /*#__PURE__*/React.createElement("p", null, "No Events")));
}