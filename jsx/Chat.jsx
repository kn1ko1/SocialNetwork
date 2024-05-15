import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js"
const { useState, useEffect } = React

const GROUP_CHAT_MESSAGE = 1
const PRIVATE_MESSAGE = 2
const CREATE_EVENT = 3

export const renderChat = ({ socket }) => {

	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Chat socket={socket} />, pageContainer)
}

export function Chat({ socket }) {
	const { currentUserId } = getCurrentUserId()

	const [sendMessage, setSendMessage] = useState("")
	const [receiveMessage, setReceiveMessage] = useState("")
	const [groupsPartOf, setGroupsPartOf] = useState([]);
	const [uniqueUsers, setUniqueUsers] = useState([]);

	let messages = document.getElementById("messages")

	useEffect(() => {
		console.log("currentUserId", currentUserId)
		const fetchUserAndGroupData = async () => {
			try {
				const promises = [];
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followedUsers`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followerUsers`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groups`));

				const results = await Promise.all(promises);

				const usersIFollowResponse = results[0]
				const usersFollowMeResponse = results[1]
				const groupsPartOfResponse = results[2]

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
				// let usersIFollowUsernames = null
				// // Extract usernames from userUsersIFollowData and usersFollowMeData
				// if (userUsersIFollowData != null) {
				// 	usersIFollowUsernames = await Promise.all(userUsersIFollowData.map(userUser => fetchUsername(userUser.subjectId)));

				// }
				// let usersFollowMeUsernames = null
				// if (userUsersFollowMeData != null) {
				// 	usersFollowMeUsernames = await Promise.all(userUsersFollowMeData.map(userFollower => fetchUsername(userFollower.subjectId)));
				// }



				let uniqueUsers = null
				if (usersIFollowData != null & usersFollowMeData != null) {
					uniqueUsers = Array.from(new Set([...usersIFollowData, ...usersFollowMeData]));
				} else if (usersIFollowData == null) {
					uniqueUsers = usersFollowMeData
				} else if (usersFollowMeData == null) {
					uniqueUsers = usersIFollowData
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
		setSendMessage(e.target.value)
	}

	const [isChatboxVisible, setChatboxVisible] = useState(false);
	const [selectedUser, setSelectedUser] = useState(null);
	const [selectedGroup, setSelectedGroup] = useState(null);

	const handleUserClick = (user) => {
		setSelectedUser(user);
		setSelectedGroup(null); // Clear the selected group when selecting a user
		setChatboxVisible(true);
	};

	const handleGroupClick = (group) => {
		setSelectedGroup(group);
		setSelectedUser(null); // Clear the selected user when selecting a group
		setChatboxVisible(true);
	};

	const handleSubmit = (e) => {
		e.preventDefault()
		let bodymessage = {
			body: sendMessage,
			messageType: "group",
			senderId: currentUserId,
			targetId: 100,
		}
		let obj = { code: GROUP_CHAT_MESSAGE, body: JSON.stringify(bodymessage) }
		socket.send(JSON.stringify(obj))
		setSendMessage("")

		// Toggle the value of isChatboxVisible when a chat is selected
		setChatboxVisible(!isChatboxVisible);
	}

	socket.onmessage = function (e) {
		let data = JSON.parse(e.data)
		let msg = JSON.parse(data.body).body
		// setReceiveMessage(msg)
		// console.log("receiveMessage:", receiveMessage)
		let entry = document.createElement("li")
		entry.appendChild(document.createTextNode(msg))
		messages.appendChild(entry)
	}

	const messageStyle = {
		color: "orange",
	}

	return (
		<div>
			<h1>Chat</h1>
			<h3>Users</h3>
			{uniqueUsers && uniqueUsers.length > 0 ? (
				<ul>
					{uniqueUsers.map((user, index) => (
						<li key={index}>
							<a href="#" onClick={() => handleUserClick(user)}>{user.username}</a>
						</li>
					))}
				</ul>
			) : (
				<p>You're not following/followed by any users</p>
			)}
			<h3>Groups</h3>
			{groupsPartOf && groupsPartOf.length > 0 ? (
				<ul>
					{groupsPartOf.map((group, index) => (
						<li key={index}>
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
			<form id="chatbox" onSubmit={handleSubmit} style={{ display: isChatboxVisible ? "block" : "none" }}>
				<textarea onChange={handleMessages}></textarea>
				<button type="submit" className="btn btn-primary">
					send
				</button>
			</form>
		</div>
	);
};	
