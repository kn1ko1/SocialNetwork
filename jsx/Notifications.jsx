const { useState, useEffect } = React
import { useSocket } from "./shared/UserProvider.js"
export const renderNotifications = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Notifications />, pageContainer)
}

export function Notifications() {
	const { currentUserId } = useSocket();
	const [notifications, setNotifications] = useState({})


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
						
							<li key={index}><GroupInvite notification={notification}/></li>
					

					))}
				</ul>
			) : (
				<div>No notifications</div>
			)}
		</div>
	);
}

function GroupInvite({ notification }) {
	return (
		<div id="GroupInvite" className="card" style={{ maxWidth: "400px" }}>
			User {notification.senderId} invited you to join Group {notification.objectId}
		</div>
	)
}

function GroupRequest({ notification }) {
	return (
		<div id="GroupRequest" className="card" style={{ maxWidth: "400px" }}>
			User {notification.senderId} has requested to join Group {notification.objectId}
		</div>
	)
}