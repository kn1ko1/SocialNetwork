import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { formattedDate } from "./components/shared/FormattedDate.js";
const {
  useState,
  useEffect
} = React;
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
  const [currentUser, setCurrentUser] = useState({});
  const [messageCode, setMessageCode] = useState(0);
  const [messageType, setMessageType] = useState("");
  const [targetId, setTargetId] = useState(0);
  const [sendMessage, setSendMessage] = useState("");
  const [groupsPartOf, setGroupsPartOf] = useState([]);
  const [uniqueUsers, setUniqueUsers] = useState([]);
  const [isChatboxVisible, setChatboxVisible] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const [selectedGroup, setSelectedGroup] = useState(null);
  const [isEmojiPickerVisible, setEmojiPickerVisible] = useState(false);
  useEffect(() => {
    console.log("currentUserId", currentUserId);
    const fetchUserAndGroupData = async () => {
      try {
        const promises = [];
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followedUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followerUsers`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groups`));
        const results = await Promise.all(promises);
        const currentUserResponse = results[0];
        const usersIFollowResponse = results[1];
        const usersFollowMeResponse = results[2];
        const groupsPartOfResponse = results[3];
        if (!currentUserResponse.ok) {
          throw new Error("Failed to fetch current user");
        }
        if (!usersIFollowResponse.ok) {
          throw new Error("Failed to fetch usersIFollow list");
        }
        if (!usersFollowMeResponse.ok) {
          throw new Error("Failed to fetch usersFollowMe list");
        }
        if (!groupsPartOfResponse.ok) {
          throw new Error("Failed to fetch groupsPartOf list");
        }
        const currentUser = await currentUserResponse.json();
        const usersIFollowData = await usersIFollowResponse.json();
        const usersFollowMeData = await usersFollowMeResponse.json();
        const groupsPartOfData = await groupsPartOfResponse.json();
        setCurrentUser(currentUser);
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
        console.error("Error fetching possible chat options list:", error);
      }
    };
    if (currentUserId !== null) {
      fetchUserAndGroupData();
    }
  }, [currentUserId]);
  const handleMessages = e => {
    setSendMessage(e.target.value);
  };
  const handleUserClick = async user => {
    setSelectedUser(user);
    setMessageCode(2);
    setMessageType("users");
    setTargetId(user.userId);
    setSelectedGroup(null); // Clear the selected group when selecting a user
    setChatboxVisible(true);
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
      const messageCard = createMessageCard(message);
      chatHistory.appendChild(messageCard);
    });
  };
  const handleGroupClick = async group => {
    setSelectedGroup(group);
    setMessageCode(1);
    setMessageType("groups");
    setTargetId(group.groupId);
    setSelectedUser(null); // Clear the selected user when selecting a group
    setChatboxVisible(true);
    const messagesResponse = await fetch(`http://localhost:8080/api/groups/${group.groupId}/messages`);
    if (!messagesResponse.ok) {
      throw new Error(`Failed to fetch messages for group ${group.groupId}`);
    }
    const messages = await messagesResponse.json();
    console.log("Messages:", messages);
    let chatHistory = document.getElementById("chatHistory");
    // Clear the chat history
    chatHistory.innerHTML = "";

    // Add new messages to chat history
    messages.forEach(message => {
      const messageCard = createMessageCard(message);
      chatHistory.appendChild(messageCard);
    });
  };
  const handleSubmit = e => {
    e.preventDefault();
    let currentTimeInMilliseconds = new Date().getTime();
    let submitMessage = {
      body: sendMessage,
      createdAt: currentTimeInMilliseconds,
      messageType: messageType,
      senderId: currentUserId,
      senderUsername: currentUser.username,
      targetId: targetId,
      updatedAt: currentTimeInMilliseconds
    };
    let obj = {
      code: messageCode,
      body: JSON.stringify(submitMessage)
    };
    socket.send(JSON.stringify(obj));
    setSendMessage("");
    let chatHistory = document.getElementById("chatHistory");
    const messageCard = createMessageCard(submitMessage);
    if (chatHistory.childNodes.length > 0) {
      chatHistory.prepend(messageCard); // Add to the start
    } else {
      chatHistory.appendChild(messageCard); // Add normally if length is 0
    }
  };

  // When refceiving a websocket message in chat:
  socket.onmessage = function (e) {
    let data = JSON.parse(e.data);
    let message = JSON.parse(data.body);
    console.log("you received websocket message:", message);
    let chatHistory = document.getElementById("chatHistory");
    if (data.code == 1 && selectedGroup.groupId == message.targetId || data.code == 2 && selectedUser.username == message.senderUsername) {
      const messageCard = createMessageCard(message);
      if (chatHistory.childNodes.length > 0) {
        chatHistory.prepend(messageCard); // Add to the start
      } else {
        chatHistory.appendChild(messageCard); // Add normally if length is 0
      }
    }
  };
  const createMessageCard = message => {
    const card = document.createElement("div");
    card.classList.add("card", "mb-3");
    card.style.backgroundColor = "transparent";
    const cardBody = document.createElement("div");
    cardBody.classList.add("card-body", "p-3");
    const userNameElement = document.createElement("h6");
    userNameElement.classList.add("fw-bold", "mb-1");
    userNameElement.textContent = message.senderUsername;
    const messageBodyElement = document.createElement("p");
    messageBodyElement.classList.add("mb-1");
    messageBodyElement.textContent = message.body;

    // Add the following CSS styles
    messageBodyElement.style.backgroundColor = "#f0f0f0"; // Adjust the background color as needed
    messageBodyElement.style.padding = "10px"; // Adjust the padding as needed
    messageBodyElement.style.borderRadius = "10px"; // Adjust the border radius as needed
    messageBodyElement.style.wordWrap = "break-word"; // Enable word wrapping

    const sentAtElement = document.createElement("small");
    sentAtElement.classList.add("text-muted");
    sentAtElement.textContent = `Sent at ${formattedDate(message.createdAt)}`;
    if (message.senderUsername === currentUser.username) {
      cardBody.classList.add("me-auto");
      //console.log("did we get here 1?")
    } else {
      cardBody.classList.add("ms-auto");
      userNameElement.classList.add("text-end");
      messageBodyElement.classList.add("text-end");
      //console.log("did we get here 2?")
    }
    cardBody.appendChild(userNameElement);
    cardBody.appendChild(messageBodyElement);
    cardBody.appendChild(sentAtElement);
    card.appendChild(cardBody);
    return card;
  };

  // Function to handle opening/closing the emoji picker
  const toggleEmojiPicker = () => {
    setEmojiPickerVisible(!isEmojiPickerVisible);
  };
  const handleEmojiSelect = emoji => {
    // Get the current text in the textarea
    const messageTextarea = document.getElementById("message-textarea");
    const messageText = messageTextarea.value;

    // if (messageText === "") {
    //   messageText + " "
    // }

    // Append the emoji to the end of the text
    const updatedMessageText = messageText + emoji;

    // Update the message text in the textarea
    messageTextarea.value = updatedMessageText;
    setSendMessage(updatedMessageText);
  };
  const messageStyle = {
    color: "black",
    marginBottom: "20px",
    textDecoration: "underline",
    textAlign: "center",
    backgroundColor: "linear-gradient(to bottom, #c7ddef, #ffffff)",
    // Light gray background
    padding: "10px",
    // Inner spacing
    borderRadius: "10px",
    // Rounded corners
    margin: "5px 0"
  };
  const chatStyle = {
    maxWidth: "1300px",
    background: "linear-gradient(to bottom, #c7ddef, #ffffff)",
    // Light blue/grey to white gradient
    borderRadius: "10px",
    boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
    // Optional: Add shadow for depth
    padding: "40px",
    margin: "auto",
    marginBottom: "20px",
    // Adjust spacing between post cards
    border: "1px solid #ccc" // Add a thin border
  };
  const opaqueStyle = {
    backgroundColor: "rgba(255, 255, 255, 0.25)",
    // Adjust the opacity here
    maxWidth: "1300px",
    borderRadius: "10px",
    boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
    // Optional: Add shadow for depth
    padding: "40px",
    margin: "auto",
    marginBottom: "20px" // Adjust spacing between post cards
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "container-fluid"
  }, /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-4",
    style: {
      ...opaqueStyle,
      height: "100vh",
      overflowY: "auto"
    }
  }, /*#__PURE__*/React.createElement("div", {
    style: chatStyle
  }, /*#__PURE__*/React.createElement("h3", {
    style: {
      textDecoration: "underline",
      textAlign: "center"
    }
  }, "Users"), uniqueUsers && uniqueUsers.length > 0 ? /*#__PURE__*/React.createElement("ul", {
    className: "list-group"
  }, uniqueUsers.map((user, index) => /*#__PURE__*/React.createElement("li", {
    key: index,
    className: "list-group-item"
  }, /*#__PURE__*/React.createElement("a", {
    href: "#",
    onClick: () => handleUserClick(user)
  }, user.username)))) : /*#__PURE__*/React.createElement("p", null, "You're not following/followed by any users")), /*#__PURE__*/React.createElement("div", {
    style: chatStyle
  }, /*#__PURE__*/React.createElement("h3", {
    style: {
      textDecoration: "underline",
      textAlign: "center"
    }
  }, "Groups"), groupsPartOf && groupsPartOf.length > 0 ? /*#__PURE__*/React.createElement("ul", {
    className: "list-group"
  }, groupsPartOf.map((group, index) => /*#__PURE__*/React.createElement("li", {
    key: index,
    className: "list-group-item"
  }, /*#__PURE__*/React.createElement("a", {
    href: "#",
    onClick: () => handleGroupClick(group)
  }, group.title)))) : /*#__PURE__*/React.createElement("p", null, "You're not part of any groups"))), /*#__PURE__*/React.createElement("div", {
    className: `col-md-7 ${isChatboxVisible ? "d-block" : "d-none"}`,
    style: {
      ...opaqueStyle,
      height: "100vh",
      overflowY: "auto"
    }
  }, /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn-close",
    "aria-label": "Close",
    onClick: () => setChatboxVisible(false)
  }), /*#__PURE__*/React.createElement("h3", {
    id: "messages",
    style: {
      ...messageStyle,
      display: isChatboxVisible ? "block" : "none"
    }
  }, selectedUser && /*#__PURE__*/React.createElement("h3", null, "Chat with ", selectedUser.username), selectedGroup && /*#__PURE__*/React.createElement("h3", null, "Chat in ", selectedGroup.title)), /*#__PURE__*/React.createElement("form", {
    id: "chatbox",
    onSubmit: handleSubmit
  }, /*#__PURE__*/React.createElement("textarea", {
    id: "message-textarea",
    className: "form-control",
    value: sendMessage,
    onChange: handleMessages,
    placeholder: "Type your message..."
  }), /*#__PURE__*/React.createElement("div", {
    style: {
      display: "flex",
      flexDirection: "column",
      alignItems: "flex-start"
    }
  }, /*#__PURE__*/React.createElement("button", {
    type: "button",
    onClick: toggleEmojiPicker,
    style: {
      marginBottom: "10px"
    }
  }, "\uD83D\uDE0A"), isEmojiPickerVisible && /*#__PURE__*/React.createElement("div", {
    id: "emoji-picker",
    className: "emoji-picker",
    style: {
      marginBottom: "10px",
      marginTop: "10px"
    }
  }, /*#__PURE__*/React.createElement("button", {
    type: "button",
    onClick: () => handleEmojiSelect("😊")
  }, "\uD83D\uDE0A"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    onClick: () => handleEmojiSelect("😂")
  }, "\uD83D\uDE02"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    onClick: () => handleEmojiSelect("❤️")
  }, "\u2764\uFE0F"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    onClick: () => handleEmojiSelect("👍")
  }, "\uD83D\uDC4D"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    onClick: () => handleEmojiSelect("😢")
  }, "\uD83D\uDE22"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    onClick: () => handleEmojiSelect("😍")
  }, "\uD83D\uDE0D")), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary mt-2"
  }, "Send"))), /*#__PURE__*/React.createElement("div", {
    id: "chatHistory"
  }))));
}