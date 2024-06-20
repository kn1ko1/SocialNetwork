const { useState, useEffect } = React
import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js"
import { GroupInvite } from "./components/Notifications/GroupInvite.js"
import { GroupRequest } from "./components/Notifications/GroupRequest.js"
import { FollowRequest } from "./components/Notifications/FollowRequest.js"
import { EventInvite } from "./components/Notifications/EventInvite.js"

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
				console.log("notifications:", data)
			})
			.catch((error) => {
				console.error("Error fetching notifications data:", error);
			});
	};

	const handleNotificationResponse = (notificationId) => {
		// Filter out the notification with the given ID from notifications state
		const updatedNotifications = notifications.filter(
			(notification) => notification.notificationId !== notificationId
		);
		// Update notifications state with the filtered notifications
		setNotifications(updatedNotifications);
	};


	const notificationsStyle = {
		maxWidth: '1000px',
		background: 'linear-gradient(to bottom, #c7ddef, #ffffff)', // Light blue/grey to white gradient
		borderRadius: '10px',
		boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', // Optional: Add shadow for depth
		padding: '40px',
		margin: 'auto',
		marginBottom: '20px', // Adjust spacing between post cards
		border: '1px solid #ccc', // Add a thin border
	  };

	return (
		<div style={notificationsStyle} className="col-md-4">
			<h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>Notifications</h2>
			{notifications !== null && Object.keys(notifications).length > 0 ? (
				<ul>
					{Object.values(notifications).map((notification, index) => (
						<li key={index}>
							{notification.notificationType === "groupInvite" && <GroupInvite
								notification={notification}
								onNotificationResponse={handleNotificationResponse}
							/>}
							{notification.notificationType === "groupRequest" && <GroupRequest
								notification={notification}
								onNotificationResponse={handleNotificationResponse}
							/>}
							{notification.notificationType === "eventInvite" && <EventInvite
								notification={notification}
								onNotificationResponse={handleNotificationResponse}
							/>}
							{notification.notificationType === "followRequest" && <FollowRequest
								notification={notification}
								onNotificationResponse={handleNotificationResponse}
							/>}
						</li>
					))}
				</ul>
			) : (
				<div>No notifications</div>
			)}

		</div>
	);
}



