import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { fetchUsername } from "./components/shared/FetchUsername.js";
import { fetchGroupById } from "./components/shared/FetchGroupById.js";
const {
  useState,
  useEffect
} = React;
const GROUP_CHAT_MESSAGE = 1;
const PRIVATE_MESSAGE = 2;
const CREATE_EVENT = 3;
export const renderChat = ({
  socket
}) => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Chat, {
    socket: socket
  }), pageContainer);
};
export function Chat({
  socket
}) {
  const {
    currentUserId
  } = getCurrentUserId();
  const [sendMessage, setSendMessage] = useState("");
  const [receiveMessage, setReceiveMessage] = useState("");
  const [groupsPartOf, setGroupsPartOf] = useState([]);
  const [uniqueUsernames, setUniqueUsernames] = useState([]);
  let messages = document.getElementById("messages");
  useEffect(() => {
    console.log("currentUserId", currentUserId);
    const fetchUserAndGroupData = async () => {
      try {
        const promises = [];
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followerUserUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groups`));
        const results = await Promise.all(promises);
        const userUsersIFollowResponse = results[0];
        const userUsersFollowMeResponse = results[1];
        const groupsPartOfResponse = results[2];
        if (!userUsersIFollowResponse.ok) {
          throw new Error('Failed to fetch usersIFollow list');
        }
        if (!userUsersFollowMeResponse.ok) {
          throw new Error('Failed to fetch usersFollowMe list');
        }
        if (!groupsPartOfResponse.ok) {
          throw new Error('Failed to fetch groupsPartOf list');
        }
        const userUsersIFollowData = await userUsersIFollowResponse.json();
        const userUsersFollowMeData = await userUsersFollowMeResponse.json();
        const groupsPartOfData = await groupsPartOfResponse.json();
        setGroupsPartOf(groupsPartOfData);
        let usersIFollowUsernames = null;
        // Extract usernames from userUsersIFollowData and usersFollowMeData
        if (userUsersIFollowData != null) {
          usersIFollowUsernames = await Promise.all(userUsersIFollowData.map(userUser => fetchUsername(userUser.subjectId)));
        }
        let usersFollowMeUsernames = null;
        if (userUsersFollowMeData != null) {
          usersFollowMeUsernames = await Promise.all(userUsersFollowMeData.map(userFollower => fetchUsername(userFollower.subjectId)));
        }
        let uniqueUsernames = null;
        if (usersIFollowUsernames != null & usersFollowMeUsernames != null) {
          uniqueUsernames = Array.from(new Set([...usersIFollowUsernames, ...usersFollowMeUsernames]));
        } else if (usersIFollowUsernames == null) {
          uniqueUsernames = usersFollowMeUsernames;
        } else if (usersFollowMeUsernames == null) {
          uniqueUsernames = usersIFollowUsernames;
        }
        setUniqueUsernames(uniqueUsernames);
        console.log("Unique Usernames:", uniqueUsernames);
        console.log("groupsPartOfGroupNames:", groupsPartOfGroupNames);
      } catch (error) {
        console.error('Error fetching possible chat options list:', error);
      }
    };
    if (currentUserId !== null) {
      fetchUserAndGroupData();
    }
  }, [currentUserId]);
  const handleMessages = e => {
    setSendMessage(e.target.value);
  };
  const [isChatboxVisible, setChatboxVisible] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const [selectedGroup, setSelectedGroup] = useState(null);
  const handleUserClick = username => {
    setSelectedUser(username);
    setSelectedGroup(null); // Clear the selected group when selecting a user
    setChatboxVisible(true);
  };
  const handleGroupClick = group => {
    setSelectedGroup(group);
    setSelectedUser(null); // Clear the selected user when selecting a group
    setChatboxVisible(true);
  };
  const handleSubmit = e => {
    e.preventDefault();
    let bodymessage = {
      body: sendMessage,
      messageType: "group",
      senderId: currentUserId,
      targetId: 100
    };
    let obj = {
      code: GROUP_CHAT_MESSAGE,
      body: JSON.stringify(bodymessage)
    };
    socket.send(JSON.stringify(obj));
    setSendMessage("");

    // Toggle the value of isChatboxVisible when a chat is selected
    setChatboxVisible(!isChatboxVisible);
  };
  socket.onmessage = function (e) {
    let data = JSON.parse(e.data);
    let msg = JSON.parse(data.body).body;
    // setReceiveMessage(msg)
    // console.log("receiveMessage:", receiveMessage)
    let entry = document.createElement("li");
    entry.appendChild(document.createTextNode(msg));
    messages.appendChild(entry);
  };
  const messageStyle = {
    color: "orange"
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Chat"), /*#__PURE__*/React.createElement("h3", null, "Users"), uniqueUsernames && uniqueUsernames.length > 0 ? /*#__PURE__*/React.createElement("ul", null, uniqueUsernames.map((username, index) => /*#__PURE__*/React.createElement("li", {
    key: index
  }, /*#__PURE__*/React.createElement("a", {
    href: "#",
    onClick: () => handleUserClick(username)
  }, username)))) : /*#__PURE__*/React.createElement("p", null, "You're not following/followed by any users"), /*#__PURE__*/React.createElement("h3", null, "Groups"), groupsPartOf && groupsPartOf.length > 0 ? /*#__PURE__*/React.createElement("ul", null, groupsPartOf.map((group, index) => /*#__PURE__*/React.createElement("li", {
    key: index
  }, /*#__PURE__*/React.createElement("a", {
    href: "#",
    onClick: () => handleGroupClick(group)
  }, group.title)))) : /*#__PURE__*/React.createElement("p", null, "You're not part of any groups"), /*#__PURE__*/React.createElement("ul", {
    id: "messages",
    style: {
      ...messageStyle,
      display: isChatboxVisible ? "block" : "none"
    }
  }, selectedUser && /*#__PURE__*/React.createElement("li", null, "Chat with ", selectedUser), selectedGroup && /*#__PURE__*/React.createElement("li", null, "Chat in ", selectedGroup.title)), /*#__PURE__*/React.createElement("form", {
    id: "chatbox",
    onSubmit: handleSubmit,
    style: {
      display: isChatboxVisible ? "block" : "none"
    }
  }, /*#__PURE__*/React.createElement("textarea", {
    onChange: handleMessages
  }), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "send")));
}
;