const { useState, useEffect } = React
import { fetchUsername } from "../FetchUsername.js";
import { fetchGroupName } from "../FetchGroupName.js";
import { respondToNotification } from "../RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";



export function GroupRequest({ notification }) {
	const [username, setUsername] = useState("")
	const [groupName, setGroupName] = useState("")

	useEffect(() => {
		fetchUsername(notification.senderId)
			.then(username => setUsername(username));
		fetchGroupName(notification.objectId)
			.then(groupName => setGroupName(groupName));
	}, [notification.senderId, notification.objectId]);

	return (
		<div id={notification.notificationType} style={notificationCardStyle} className="card">
			<div className="row">
				<div className="col">
					{username} requested to join {groupName}
				</div>
				<div className="col-auto d-flex align-items-center"> {/* col-auto makes this column width fit its content */}
					<button onClick={() => respondToNotification("confirm", notification)}>&#10003;</button>
					<button onClick={() => respondToNotification("deny", notification)}>&#10007;</button>
				</div>
			</div>
		</div>
	);

}