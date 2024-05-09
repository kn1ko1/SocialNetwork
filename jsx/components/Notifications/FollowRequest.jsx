const { useState, useEffect } = React
import { fetchUsername } from "../shared/FetchUsername.js";
import { respondToNotification } from "./RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";


export function FollowRequest({ notification, onNotificationResponse }) {
	const [username, setUsername] = useState("");

	useEffect(() => {
		fetchUsername(notification.senderId)
			.then(username => setUsername(username));
	}, [notification.senderId]);

	const handleNotificationResponse = async (responseType) => {
		// Call the respondToNotification function to handle the response
		respondToNotification(responseType, notification);
		// Call the parent component's callback to remove this notification
		onNotificationResponse(notification.notificationId);
	};

	return (
		<div id={notification.notificationType} style={notificationCardStyle} className="card">
			<div className="row">
				<div className="col">
					{username} has requested to follow you
				</div>
				<div className="col-auto d-flex align-items-center">
					<button onClick={() => handleNotificationResponse("confirm")}>&#10003;</button>
					<button onClick={() => handleNotificationResponse("deny")}>&#10007;</button>
				</div>
			</div>
		</div>
	);
}