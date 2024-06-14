import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { PostFormGroup } from "./components/GroupDetail/PostFormGroup.js";
import { EventForm } from "./components/GroupDetail/EventForm.js";
import { GroupDetailsUserList } from "./components/GroupDetail/GroupDetailsUserList.js";
import { PostCard } from "./components/shared/PostCard.js";
import { GroupDetailsEvents } from "./components/GroupDetail/GroupDetailsEvent.js";
import { fetchCommentsForPosts } from "./components/shared/FetchCommentsForPosts.js";
const {
  useState,
  useEffect
} = React;
export const renderGroupDetails = (group, socket) => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(GroupDetails, {
    group: group,
    socket: socket
  }), pageContainer);
};
export function GroupDetails({
  group,
  socket
}) {
  const {
    currentUserId
  } = getCurrentUserId();
  const [userList, setUserList] = useState([]);
  const [groupMembers, setGroupMembers] = useState([]);
  const [groupPosts, setGroupPosts] = useState([]);
  const [groupMessages, setGroupMessages] = useState([]);
  const [groupEvents, setGroupEvents] = useState([]);
  if (group.isMember) {
    useEffect(() => {
      fetchGroupData(group.groupId);
    }, [group.groupId]);
  }
  const fetchGroupData = async groupId => {
    try {
      const promises = [];
      promises.push(fetch(`http://localhost:8080/api/users`));
      promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/groupUsers`));
      promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/posts`));
      promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/messages`));
      promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/events`));
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
      if (postsData != null) {
        const postsWithComments = await fetchCommentsForPosts(postsData);
        setGroupPosts(postsWithComments);
      } else {
        setGroupPosts(null);
      }
      setGroupMessages(messagesData);
      setGroupEvents(eventsData);
    } catch (error) {
      console.error('Error fetching group data:', error);
    }
  };
  async function AddGroupUser(userId, groupId, notificationType) {
    const notificationtData = {
      notificationType: notificationType,
      objectId: groupId,
      senderId: currentUserId,
      status: "pending",
      targetId: userId
    };
    console.log('notificationtData:', notificationtData);
    let codeNum = 4;
    if (notificationType == "groupInvite") {
      codeNum = 5;
    }
    try {
      let obj = {
        code: codeNum,
        body: JSON.stringify(notificationtData)
      };
      socket.send(JSON.stringify(obj));
    } catch (error) {
      console.error('Error adding group user:', error);
    }
  }
  const groupDetailsStyle = {
    maxWidth: '1000px',
    background: 'linear-gradient(to bottom, #c7ddef, #ffffff)',
    // Light blue/grey to white gradient
    borderRadius: '10px',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
    // Optional: Add shadow for depth
    padding: '40px',
    margin: 'auto',
    marginBottom: '20px',
    // Adjust spacing between post cards
    border: '1px solid #ccc' // Add a thin border
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "group-details container"
  }, /*#__PURE__*/React.createElement("div", {
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline',
      fontSize: '24px',
      textAlign: 'center'
    }
  }, group.title), /*#__PURE__*/React.createElement("p", {
    style: {
      fontSize: '20px',
      textAlign: 'center'
    }
  }, group.description)), group.isMember ? /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-lg-3 text-center"
  }, /*#__PURE__*/React.createElement("div", {
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement("h3", {
    style: {
      textDecoration: 'underline'
    }
  }, "Group Members"), /*#__PURE__*/React.createElement("div", {
    className: "groupMembers"
  }, groupMembers !== null && groupMembers.length > 0 ? groupMembers.map((member, index) => {
    const user = userList.find(user => user.userId === member.userId);
    return /*#__PURE__*/React.createElement("div", {
      key: index
    }, user ? user.username : 'Unknown User');
  }) : /*#__PURE__*/React.createElement("p", null, "It's just you... Maybe you should invite someone?"))), /*#__PURE__*/React.createElement("div", {
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement(GroupDetailsUserList, {
    userList: userList,
    groupId: group.groupId,
    groupMembers: groupMembers,
    AddGroupUser: AddGroupUser
  }))), /*#__PURE__*/React.createElement("div", {
    className: "col-lg-6 text-center"
  }, /*#__PURE__*/React.createElement("div", {
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement(PostFormGroup, {
    group: group,
    fetchFunc: () => fetchGroupData(group.groupId)
  })), /*#__PURE__*/React.createElement("h3", {
    style: {
      textDecoration: 'underline'
    }
  }, "Posts"), /*#__PURE__*/React.createElement("div", {
    id: "groupPosts"
  }, groupPosts !== null ? groupPosts.map(post => /*#__PURE__*/React.createElement("li", {
    key: post.createdAt
  }, /*#__PURE__*/React.createElement(PostCard, {
    post: post,
    comments: post.comments,
    showCommentForm: true,
    fetchFunc: () => fetchGroupData(group.groupId)
  }))) : /*#__PURE__*/React.createElement("div", null, "There are no posts in this group yet")), /*#__PURE__*/React.createElement("div", {
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement("h3", {
    style: {
      textDecoration: 'underline'
    }
  }, "Messages"), /*#__PURE__*/React.createElement("div", {
    className: "groupMessages"
  }, groupMessages !== null && groupMessages.length > 0 ? groupMessages.map((message, index) => /*#__PURE__*/React.createElement("div", {
    key: index
  }, message.body)) : /*#__PURE__*/React.createElement("p", null, "No Messages")))), /*#__PURE__*/React.createElement("div", {
    className: "col-lg-3 text-center"
  }, /*#__PURE__*/React.createElement("div", {
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement(EventForm, {
    group: group,
    socket: socket
  })), /*#__PURE__*/React.createElement("div", {
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement(GroupDetailsEvents, {
    groupEvents: groupEvents
  })))) : /*#__PURE__*/React.createElement("div", {
    className: "text-center",
    style: groupDetailsStyle
  }, /*#__PURE__*/React.createElement("div", null, "You are not a member yet"), /*#__PURE__*/React.createElement("button", {
    onClick: () => AddGroupUser(group.creatorId, group.groupId, "groupRequest")
  }, "Request to join group")));
}