import { useSocket } from "./shared/UserProvider.js"
const { useState } = React

const GROUP_CHAT_MESSAGE = 1
const PRIVATE_MESSAGE = 2
const CREATE_EVENT = 3

export const renderChat = (messageType, targetId) => {

	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Chat messageType={messageType} targetId={targetId}/>, pageContainer)
}

export function Chat(messageType, targetId) {
	const { socket, currentUserId } = useSocket();

	const [sendMessage, setSendMessage] = useState("")
	const [receiveMessage, setReceiveMessage] = useState("")

	let messages = document.getElementById("messages")

	const handleMessages = (e) => {
		setSendMessage(e.target.value)
	}

	const handleSubmit = (e) => {
		e.preventDefault()
		let bodymessage = {
			body: sendMessage,
			messageType: messageType,
			senderId: currentUserId,
			targetId: targetId,
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