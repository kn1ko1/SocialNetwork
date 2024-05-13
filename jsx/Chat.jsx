import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js"
import { fetchUsername } from "./components/shared/FetchUsername.js"
import { fetchGroupName } from "./components/shared/FetchGroupName.js"
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
	const [usersIFollow, setUsersIFollow] = useState([]);
	const [usersFollowMe, setUsersFollowMe] = useState([]);
	const [groupsPartOf, setGroupsPartOf] = useState([]);
	const [uniqueUsernames, setUniqueUsernames] = useState([]);

	let messages = document.getElementById("messages")

	useEffect(() => {
		console.log("currentUserId", currentUserId)
		const fetchUserAndGroupData = async () => {
			try {
				const promises = [];
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followerUserUsers`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groupUsers`));

				const results = await Promise.all(promises);

				const userUsersIFollowResponse = results[0]
				const userUsersFollowMeResponse = results[1]
				const groupsPartOfResponse = results[2]

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

				let usersIFollowUsernames = null
				// Extract usernames from userUsersIFollowData and usersFollowMeData
				if (userUsersIFollowData != null) {
					usersIFollowUsernames = await Promise.all(userUsersIFollowData.map(userUser => fetchUsername(userUser.subjectId)));

				}
				let usersFollowMeUsernames = null
				if (userUsersFollowMeData != null) {
					usersFollowMeUsernames = await Promise.all(userUsersFollowMeData.map(userFollower => fetchUsername(userFollower.subjectId)));
				}
				let groupsPartOfGroupNames = null
				if (userUsersIFollowData != null) {
					groupsPartOfGroupNames = await Promise.all(groupsPartOfData.map(group => fetchGroupName(group.groupId)));
				}
				// Update the state with the extracted usernames
				setUsersIFollow(usersIFollowUsernames);
				setUsersFollowMe(usersFollowMeUsernames);
				setGroupsPartOf(groupsPartOfGroupNames);

				let uniqueUsernames = null
				if (usersIFollowUsernames != null & usersFollowMeUsernames != null) {
					uniqueUsernames = Array.from(new Set([...usersIFollowUsernames, ...usersFollowMeUsernames]));
				} else if (usersIFollowUsernames == null) {
					uniqueUsernames = usersFollowMeUsernames
				} else if (usersFollowMeUsernames == null) {
					uniqueUsernames = usersIFollowUsernames
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


	const handleMessages = (e) => {
		setSendMessage(e.target.value)
	}

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
			{uniqueUsernames && uniqueUsernames.length > 0 ? (
				<ul>
					{uniqueUsernames.map((username, index) => (
						<li key={index}>
							<a href="#">{username}</a>
						</li>
					))}
				</ul>
			) : (
				<p>You're not following/followed by any users</p>
			)}
			<h3>Groups</h3>
			{groupsPartOf && groupsPartOf.length > 0 ? (
				<ul>
					{groupsPartOf.map((groupName, index) => (
						<li key={index}>
							<a href="#">{groupName}</a>
						</li>
					))}
				</ul>
			) : (
				<p>You're not part of any groups</p>
			)}
			<ul id="messages" style={messageStyle}></ul>
			<form id="chatbox" onSubmit={handleSubmit}>
				<textarea onChange={handleMessages}></textarea>
				<button type="submit" className="btn btn-primary">
					send
				</button>
			</form>
		</div>
	);
};	
