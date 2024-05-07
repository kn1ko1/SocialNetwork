const { useState, useEffect } = React
import { fetchUsername } from "../FetchUsername.js";
import { respondToNotification } from "../RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";


export function FollowRequest({ notification }) {
	const [username, setUsername] = useState("")

	useEffect(() => {
		fetchUsername(notification.senderId)
			.then(username => setUsername(username));
	}, [notification.senderId]);

	return (
		<div id={notification.notificationType} style={notificationCardStyle} className="card">
			<div className="row">
				<div className="col">
					{username} has requested to follow you
				</div>
				<div className="col-auto d-flex align-items-center"> 
					<button onClick={() => respondToNotification("confirm", notification)}>&#10003;</button>
					<button onClick={() => respondToNotification("deny", notification)}>&#10007;</button>
				</div>
			</div>
		</div>
	);

}