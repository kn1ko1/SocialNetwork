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
export const renderGroupDetails = group => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(GroupDetails, {
    group: group
  }), pageContainer);
};
export function GroupDetails({
  group
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
      promises.push(fetch(`http://localhost:8080/api/users/transport`));
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
      const postsWithComments = await fetchCommentsForPosts(postsData);
      setUserList(userListData);
      setGroupMembers(groupMembersData);
      setGroupPosts(postsWithComments);
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
    try {
      const response = await fetch('http://localhost:8080/api/notifications', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(notificationtData)
      });
      console.log('Response:', response);
      if (response.ok) {
        // Handle success response
        console.log('Group user added successfully!');
      } else {
        // Handle error response
        console.error('Failed to add group user:', response.statusText);
      }
    } catch (error) {
      console.error('Error adding group user:', error);
    }
  }
  return /*#__PURE__*/React.createElement("div", {
    className: "group-details"
  }, /*#__PURE__*/React.createElement("h2", null, group.title), /*#__PURE__*/React.createElement("p", null, group.description), group.isMember ? /*#__PURE__*/React.createElement("div", {
    id: "groupData"
  }, /*#__PURE__*/React.createElement(PostFormGroup, {
    group: group,
    fetchFunc: () => fetchGroupData(group.groupId)
  }), /*#__PURE__*/React.createElement(EventForm, {
    group: group
  }), /*#__PURE__*/React.createElement(GroupDetailsUserList, {
    userList: userList,
    groupId: group.groupId,
    groupMembers: groupMembers,
    AddGroupUser: AddGroupUser
  }), /*#__PURE__*/React.createElement("div", {
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
  }, /*#__PURE__*/React.createElement(PostCard, {
    post: post,
    comments: post.comments,
    showCommentForm: true,
    fetchFunc: () => fetchGroupData(group.groupId)
  }))) : /*#__PURE__*/React.createElement("div", {
    id: "groupPosts"
  }, "There are no posts in this groups yet")), /*#__PURE__*/React.createElement("div", {
    className: "groupMessages"
  }, /*#__PURE__*/React.createElement("h2", null, "Messages"), groupMessages !== null && groupMessages.length > 0 ? groupMessages.map((message, index) => /*#__PURE__*/React.createElement("div", {
    key: index
  }, message.body)) : /*#__PURE__*/React.createElement("p", null, "No Messages")), /*#__PURE__*/React.createElement(GroupDetailsEvents, {
    groupEvents: groupEvents
  })) : /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("div", null, "You are not a member yet"), /*#__PURE__*/React.createElement("button", {
    onClick: () => AddGroupUser(group.creatorId, group.groupId, "groupRequest")
  }, "Request to join group")));
}