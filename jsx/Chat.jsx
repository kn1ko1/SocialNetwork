import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js"
const { useState } = React

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

	let messages = document.getElementById("messages")

	useEffect(() => {
		const fetchUserAndGroupData = async () => {
			try {
				const promises = [];
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/followerUserUsers`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/groupUsers`));

				const results = await Promise.all(promises);

				const usersIFollow = results[0]
				const usersFollowMe = results[1]
				const groupsPartOf = results[2]

				if (!usersIFollow.ok) {
					throw new Error('Failed to fetch usersIFollow list');
				}
				if (!usersFollowMe.ok) {
					throw new Error('Failed to fetch usersFollowMe list');
				}
				if (!groupsPartOf.ok) {
					throw new Error('Failed to fetch groupsPartOf list');
				}

				const usersIFollowData = await usersIFollowResponse.json();
				const usersFollowMeData = await usersFollowMeResponse.json();
				const groupsPartOfData = await groupsPartOfResponse.json();

				setUsersIFollow(usersIFollowData);
				setUsersFollowMe(usersFollowMeData);
				setGroupsPartOf(groupsPartOfData);

				console.log("usersIFollowData:", usersIFollowData)
				console.log("usersFollowMeData:", usersFollowMeData)
				console.log("groupsPartOfData:", groupsPartOfData)

			} catch (error) {
				console.error('Error fetching possible chat options list:', error);
			}
		};

		fetchUserAndGroupData();
	}, []);

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
			<h3>Groups</h3>
			<ul id="messages" style={messageStyle}></ul>
			<form id="chatbox" onSubmit={handleSubmit}>
				<textarea onChange={handleMessages}></textarea>
				<button type="submit" className="btn btn-primary">
					send
				</button>
			</form>
		</div>
	)
}