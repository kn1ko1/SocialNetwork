import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { fetchUsername } from "./components/shared/FetchUsername.js";
import { fetchGroupName } from "./components/shared/FetchGroupName.js";
const {
  useState,
  useEffect
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
  const [usersIFollow, setUsersIFollow] = useState([]);
  const [usersFollowMe, setUsersFollowMe] = useState([]);
  const [groupsPartOf, setGroupsPartOf] = useState([]);
  const [uniqueUsernames, setUniqueUsernames] = useState([]);
  let messages = document.getElementById("messages");
  useEffect(() => {
    const fetchUserAndGroupData = async () => {
      try {
        const promises = [];
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followerUserUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groupUsers`));
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

        // Extract usernames from userUsersIFollowData and usersFollowMeData
        const usersIFollowUsernames = await Promise.all(userUsersIFollowData.map(userUser => fetchUsername(userUser.subjectId)));
        const usersFollowMeUsernames = await Promise.all(userUsersFollowMeData.map(userFollower => fetchUsername(userFollower.subjectId)));
        const groupsPartOfGroupNames = await Promise.all(groupsPartOfData.map(group => fetchGroupName(group.groupId)));

        // Update the state with the extracted usernames
        setUsersIFollow(usersIFollowUsernames);
        setUsersFollowMe(usersFollowMeUsernames);
        setGroupsPartOf(groupsPartOfGroupNames);
        const uniqueUsernames = Array.from(new Set([...usersIFollowUsernames, ...usersFollowMeUsernames]));
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
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Chat"), /*#__PURE__*/React.createElement("h3", null, "Users"), /*#__PURE__*/React.createElement("ul", null, uniqueUsernames.map((username, index) => /*#__PURE__*/React.createElement("li", {
    key: index
  }, /*#__PURE__*/React.createElement("a", {
    href: "#"
  }, username)))), /*#__PURE__*/React.createElement("h3", null, "Groups"), /*#__PURE__*/React.createElement("ul", null, groupsPartOf.map((groupName, index) => /*#__PURE__*/React.createElement("li", {
    key: index
  }, /*#__PURE__*/React.createElement("a", {
    href: "#"
  }, groupName)))), /*#__PURE__*/React.createElement("ul", {
    id: "messages",
    style: messageStyle
  }), /*#__PURE__*/React.createElement("form", {
    id: "chatbox",
    onSubmit: handleSubmit
  }, /*#__PURE__*/React.createElement("textarea", {
    onChange: handleMessages
  }), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "send")));
}