const { useState, useEffect } = React
import { getCurrentUserId } from "./shared/getCurrentUserId.js"

export const renderNotifications = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Notifications />, pageContainer)
}

export function Notifications() {
	const { currentUserId } = getCurrentUserId()
	const [notifications, setNotifications] = useState(null);


	useEffect(() => {
		if (currentUserId !== null) {
			fetchNotifications();
		}
	}, [currentUserId]);

	const fetchNotifications = () => {
		fetch(`http://localhost:8080/api/users/${currentUserId}/notifications`)
			.then((response) => response.json())
			.then((data) => {
				setNotifications(data);
			})
			.catch((error) => {
				console.error("Error fetching notifications data:", error);
			});
	};


	return (
		<div>
			<h1>Notifications</h1>
			{notifications !== null && Object.keys(notifications).length > 0 ? (
				<ul>
					{Object.values(notifications).map((notification, index) => (
						<li key={index}><GroupInvite notification={notification} /></li>
					))}
				</ul>
			) : (
				<div>No notifications</div>
			)}

		</div>
	);
}

function GroupInvite({ notification }) {
	const [username, setUsername] = useState("")
	const [groupName, setGroupName] = useState("")


	useEffect(() => {
		fetchUsername()
		fetchGroupName()
	}, []);

	const fetchUsername = () => {
		fetch(`http://localhost:8080/api/users/${notification.senderId}`)
			.then((response) => response.json())
			.then((data) => {
				setUsername(data.username);
			})
			.catch((error) => {
				console.error("Error fetching notifications data:", error);
			});
	};
	const fetchGroupName = () => {
		fetch(`http://localhost:8080/api/groups/${notification.objectId}`)
			.then((response) => response.json())
			.then((data) => {
				setGroupName(data.title);
			})
			.catch((error) => {
				console.error("Error fetching notifications data:", error);
			});
	};

	const respondToNotification = (reply, notification) => {
		const data = {
			reply: reply,
			notification: notification
		};
	
		fetch(`http://localhost:8080/api/notifications/${notification.notificationId}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		})
		.then(response => response.json())
		.then(data => {
			// Handle success response
			console.log("Response sent successfully:", data);
		})
		.catch(error => {
			console.error("Error sending response:", error);
		});
	};
	
	return (
		<div id="GroupInvite" className="card" style={{ maxWidth: "400px" }}>
			{username} invited you to join {groupName}
			<button onClick={() => respondToNotification("confirm", notification)}>&#10003;</button>
			<button onClick={() => respondToNotification("deny", notification)}>&#10007;</button>
		</div>
	);
	
}

function GroupRequest({ notification }) {
	return (
		<div id="GroupRequest" className="card" style={{ maxWidth: "400px" }}>
			User {notification.senderId} has requested to join {notification.objectId}

		</div>
	)
}