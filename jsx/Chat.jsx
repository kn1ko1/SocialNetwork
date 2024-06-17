import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js"
import { formattedDate } from "./components/shared/FormattedDate.js";
const { useState, useEffect } = React

// const GROUP_CHAT_MESSAGE = 1;
// const PRIVATE_MESSAGE = 2;

export const renderChat = ({ socket }) => {
    const pageContainer = document.querySelector(".page-container");
    ReactDOM.render(<Chat socket={socket} />, pageContainer);
}

export function Chat({ socket }) {
    const { currentUserId } = getCurrentUserId();
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
                    throw new Error('Failed to fetch current user');
                }
                if (!usersIFollowResponse.ok) {
                    throw new Error('Failed to fetch usersIFollow list');
                }
                if (!usersFollowMeResponse.ok) {
                    throw new Error('Failed to fetch usersFollowMe list');
                }
                if (!groupsPartOfResponse.ok) {
                    throw new Error('Failed to fetch groupsPartOf list');
                }

                const currentUser = await currentUserResponse.json();
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

    const handleMessages = (e) => {
        setSendMessage(e.target.value);
    }


    const handleUserClick = async (user) => {
        const messagesResponse = await fetch(`http://localhost:8080/api/users/${currentUserId}/messages/${user.userId}`)
        if (!messagesResponse.ok) {
            throw new Error(`Failed to fetch messages between user ${currentUserId} and user ${user.userId}`);
        }
        const messages = await messagesResponse.json();
        console.log("Messages:", messages)

        let chatHistory = document.getElementById("chatHistory")
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

    const handleGroupClick = (group) => {
        setSelectedGroup(group);
        setMessageCode(1);
        setMessageType("groups");
        setTargetId(group.groupId);
        setSelectedUser(null); // Clear the selected user when selecting a group
        setChatboxVisible(true);
    };

    const handleSubmit = (e) => {
        console.log("currentUserId in the submit is", currentUserId)
        e.preventDefault();
        let bodymessage = {
            body: sendMessage,
            messageType: messageType,
            senderId: currentUserId,
            targetId: targetId,
        }
        let obj = { code: messageCode, body: JSON.stringify(bodymessage) }
        socket.send(JSON.stringify(obj));
        setSendMessage("");
    }

    socket.onmessage = function (e) {
        let data = JSON.parse(e.data);
        let message = JSON.parse(data.body).body;
        console.log("you received websocket message:", message);
        let chatHistory = document.getElementById("chatHistory");

        const messageCard = createMessageCard(user, message);

        // Insert the new message at the beginning of chatHistory
        if (chatHistory.firstChild) {
            chatHistory.insertBefore(messageCard, chatHistory.firstChild);
        } else {
            chatHistory.appendChild(messageCard);
        }
    }

    const createMessageCard = (user, message) => {
        const card = document.createElement("div");
        card.classList.add("card", "mb-3");

        const cardBody = document.createElement("div");
        cardBody.classList.add("card-body", "p-3");

        const userNameElement = document.createElement("h6");
        userNameElement.classList.add("fw-bold", "mb-1");
        userNameElement.textContent = message.senderUsername

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
        color: "orange",
    }

    // Function to handle opening/closing the emoji picker
    const toggleEmojiPicker = () => {
        setEmojiPickerVisible(!isEmojiPickerVisible);
    };

    const handleEmojiSelect = (emoji) => {
        // Get the current text in the textarea
        const messageTextarea = document.getElementById('message-textarea');
        const messageText = messageTextarea.value;

        // Append the emoji to the end of the text
        const updatedMessageText = messageText + emoji;

        // Update the message text in the textarea
        messageTextarea.value = updatedMessageText;
    };

    return (
        <div className="container">
            <h1>Chat</h1>
            <h3>Users</h3>
            {uniqueUsers && uniqueUsers.length > 0 ? (
                <ul className="list-group">
                    {uniqueUsers.map((user, index) => (
                        <li key={index} className="list-group-item">
                            <a href="#" onClick={() => handleUserClick(user)}>{user.username}</a>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>You're not following/followed by any users</p>
            )}
            <h3>Groups</h3>
            {groupsPartOf && groupsPartOf.length > 0 ? (
                <ul className="list-group">
                    {groupsPartOf.map((group, index) => (
                        <li key={index} className="list-group-item">
                            <a href="#" onClick={() => handleGroupClick(group)}>{group.title}</a>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>You're not part of any groups</p>
            )}
            <ul id="messages" style={{ ...messageStyle, display: isChatboxVisible ? "block" : "none" }}>
                {selectedUser && <li>Chat with {selectedUser.username}</li>}
                {selectedGroup && <li>Chat in {selectedGroup.title}</li>}
            </ul>
            <div id="chatHistory"></div>
            <form id="chatbox" onSubmit={handleSubmit} style={{ display: isChatboxVisible ? "block" : "none" }}>
                {/* Message input */}
                <div>
                    <textarea
                        id="message-textarea"
                        className="form-control"
                        value={sendMessage}
                        onChange={handleMessages}
                        placeholder="Type your message..."
                    ></textarea>
                    {/* Emoji button */}
                    <button onClick={toggleEmojiPicker}>😊</button>
                    {/* Emoji picker */}
                    {isEmojiPickerVisible && (
                        <div id="emoji-picker" className="emoji-picker">
                            <button onClick={() => handleEmojiSelect('😊')}>😊</button>
                            <button onClick={() => handleEmojiSelect('😂')}>😂</button>
                            <button onClick={() => handleEmojiSelect('❤️')}>❤️</button>
                            {/* Add more emoji buttons as needed */}
                        </div>
                    )}
                </div>
                {/* Send button */}
                <button type="submit" className="btn btn-primary mt-2">Send</button>
            </form>
        </div>
    );
}

