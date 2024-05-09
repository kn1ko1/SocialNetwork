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
			})
			.catch((error) => {
				console.error("Error fetching notifications data:", error);
			});console.log("notifications:", data)
	};

	const handleNotificationResponse = (notificationId) => {
		console.log("notificationId", notificationId)
		// Filter out the notification with the given ID from notifications state
		const updatedNotifications = notifications.filter(
			(notification) => notification.notificationId !== notificationId
		);
		// Update notifications state with the filtered notifications
		setNotifications(updatedNotifications);
	};

	return (
		<div>
			<h1>Notifications</h1>
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



