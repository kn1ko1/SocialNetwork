import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { formattedDate } from "./components/shared/FormattedDate.js";
const {
  useState,
  useEffect
} = React;

// const GROUP_CHAT_MESSAGE = 1;
// const PRIVATE_MESSAGE = 2;

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
  const [messageCode, setMessageCode] = useState(0);
  const [messageType, setMessageType] = useState("");
  const [targetId, setTargetId] = useState(0);
  const [sendMessage, setSendMessage] = useState("");
  const [groupsPartOf, setGroupsPartOf] = useState([]);
  const [uniqueUsers, setUniqueUsers] = useState([]);
  const [isEmojiPickerVisible, setEmojiPickerVisible] = useState(false);
  let messages = document.getElementById("messages");
  useEffect(() => {
    console.log("currentUserId", currentUserId);
    const fetchUserAndGroupData = async () => {
      try {
        const promises = [];
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followedUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followerUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groups`));
        const results = await Promise.all(promises);
        const usersIFollowResponse = results[0];
        const usersFollowMeResponse = results[1];
        const groupsPartOfResponse = results[2];
        if (!usersIFollowResponse.ok) {
          throw new Error('Failed to fetch usersIFollow list');
        }
        if (!usersFollowMeResponse.ok) {
          throw new Error('Failed to fetch usersFollowMe list');
        }
        if (!groupsPartOfResponse.ok) {
          throw new Error('Failed to fetch groupsPartOf list');
        }
        const usersIFollowData = await usersIFollowResponse.json();
        const usersFollowMeData = await usersFollowMeResponse.json();
        const groupsPartOfData = await groupsPartOfResponse.json();
        setGroupsPartOf(groupsPartOfData);
        let uniqueUsers = null;
        if (usersIFollowData != null && usersFollowMeData != null) {
          uniqueUsers = Array.from(new Set([...usersIFollowData, ...usersFollowMeData]));
        } else if (usersIFollowData == null) {
          uniqueUsers = usersFollowMeData;
        } else if (usersFollowMeData == null) {
          uniqueUsers = usersIFollowData;
        }
        setUniqueUsers(uniqueUsers);
        console.log("Unique Usernames:", uniqueUsers);
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
  const handleUserClick = async user => {
    const messagesResponse = await fetch(`http://localhost:8080/api/users/${currentUserId}/messages/${user.userId}`);
    if (!messagesResponse.ok) {
      throw new Error(`Failed to fetch messages between user ${currentUserId} and user ${user.userId}`);
    }
    const messages = await messagesResponse.json();
    console.log("Messages:", messages);
    let chatHistory = document.getElementById("chatHistory");
    // Clear the chat history
    chatHistory.innerHTML = "";

    // Add new messages to chat history
    messages.forEach(message => {
      const messageCard = createMessageCard(user, message);
      chatHistory.appendChild(messageCard);
    });
    setSelectedUser(user);
    setMessageCode(2);
    setMessageType("users");
    setTargetId(user.userId);
    setSelectedGroup(null); // Clear the selected group when selecting a user
    setChatboxVisible(true);
  };
  const handleGroupClick = group => {
    setSelectedGroup(group);
    setMessageCode(1);
    setMessageType("groups");
    setTargetId(group.groupId);
    setSelectedUser(null); // Clear the selected user when selecting a group
    setChatboxVisible(true);
  };
  const handleSubmit = e => {
    e.preventDefault();
    let bodymessage = {
      body: sendMessage,
      messageType: messageType,
      senderId: currentUserId,
      targetId: targetId
    };
    let obj = {
      code: messageCode,
      body: JSON.stringify(bodymessage)
    };
    socket.send(JSON.stringify(obj));
    setSendMessage("");
  };
  socket.onmessage = function (e) {
    let data = JSON.parse(e.data);
    let msg = JSON.parse(data.body).body;
    console.log("you received websocket message:", msg);
    let entry = document.createElement("li");
    entry.appendChild(document.createTextNode(msg));
    messages.appendChild(entry);
  };
  const createMessageCard = (user, message) => {
    const card = document.createElement("div");
    card.classList.add("card", "mb-3");
    const cardBody = document.createElement("div");
    cardBody.classList.add("card-body", "p-3");
    const userNameElement = document.createElement("h6");
    userNameElement.classList.add("fw-bold", "mb-1");
    userNameElement.textContent = message.senderUsername;
    const messageBodyElement = document.createElement("p");
    messageBodyElement.classList.add("mb-1");
    messageBodyElement.textContent = message.body;
    const sentAtElement = document.createElement("small");
    sentAtElement.classList.add("text-muted");
    sentAtElement.textContent = `Sent at ${formattedDate(message.createdAt)}`;
    cardBody.appendChild(userNameElement);
    cardBody.appendChild(messageBodyElement);
    cardBody.appendChild(sentAtElement);
    card.appendChild(cardBody);
    return card;
  };
  const messageStyle = {
    color: "orange"
  };

  // Function to handle opening/closing the emoji picker
  const toggleEmojiPicker = () => {
    setEmojiPickerVisible(!isEmojiPickerVisible);
  };
  const handleEmojiSelect = emoji => {
    // Get the current text in the textarea
    const messageTextarea = document.getElementById('message-textarea');
    const messageText = messageTextarea.value;

    // Append the emoji to the end of the text
    const updatedMessageText = messageText + emoji;

    // Update the message text in the textarea
    messageTextarea.value = updatedMessageText;
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "container"
  }, /*#__PURE__*/React.createElement("h1", null, "Chat"), /*#__PURE__*/React.createElement("h3", null, "Users"), uniqueUsers && uniqueUsers.length > 0 ? /*#__PURE__*/React.createElement("ul", {
    className: "list-group"
  }, uniqueUsers.map((user, index) => /*#__PURE__*/React.createElement("li", {
    key: index,
    className: "list-group-item"
  }, /*#__PURE__*/React.createElement("a", {
    href: "#",
    onClick: () => handleUserClick(user)
  }, user.username)))) : /*#__PURE__*/React.createElement("p", null, "You're not following/followed by any users"), /*#__PURE__*/React.createElement("h3", null, "Groups"), groupsPartOf && groupsPartOf.length > 0 ? /*#__PURE__*/React.createElement("ul", {
    className: "list-group"
  }, groupsPartOf.map((group, index) => /*#__PURE__*/React.createElement("li", {
    key: index,
    className: "list-group-item"
  }, /*#__PURE__*/React.createElement("a", {
    href: "#",
    onClick: () => handleGroupClick(group)
  }, group.title)))) : /*#__PURE__*/React.createElement("p", null, "You're not part of any groups"), /*#__PURE__*/React.createElement("ul", {
    id: "messages",
    style: {
      ...messageStyle,
      display: isChatboxVisible ? "block" : "none"
    }
  }, selectedUser && /*#__PURE__*/React.createElement("li", null, "Chat with ", selectedUser.username), selectedGroup && /*#__PURE__*/React.createElement("li", null, "Chat in ", selectedGroup.title)), /*#__PURE__*/React.createElement("div", {
    id: "chatHistory"
  }), /*#__PURE__*/React.createElement("form", {
    id: "chatbox",
    onSubmit: handleSubmit,
    style: {
      display: isChatboxVisible ? "block" : "none"
    }
  }, /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("textarea", {
    id: "message-textarea",
    className: "form-control",
    value: sendMessage,
    onChange: handleMessages,
    placeholder: "Type your message..."
  }), /*#__PURE__*/React.createElement("button", {
    onClick: toggleEmojiPicker
  }, "\uD83D\uDE0A"), isEmojiPickerVisible && /*#__PURE__*/React.createElement("div", {
    id: "emoji-picker",
    className: "emoji-picker"
  }, /*#__PURE__*/React.createElement("button", {
    onClick: () => handleEmojiSelect('😊')
  }, "\uD83D\uDE0A"), /*#__PURE__*/React.createElement("button", {
    onClick: () => handleEmojiSelect('😂')
  }, "\uD83D\uDE02"), /*#__PURE__*/React.createElement("button", {
    onClick: () => handleEmojiSelect('❤️')
  }, "\u2764\uFE0F"))), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary mt-2"
  }, "Send")));
}