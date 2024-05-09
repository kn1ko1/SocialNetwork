const { useState, useEffect } = React
import { fetchGroupName } from "../shared/FetchGroupName.js";
import { fetchUsername } from "../shared/FetchUsername.js";
import { respondToNotification } from "./RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";

export function GroupInvite({ notification, onNotificationResponse }) {
    const [username, setUsername] = useState("");
    const [groupName, setGroupName] = useState("");

    useEffect(() => {
        fetchUsername(notification.senderId)
            .then(username => setUsername(username));
        fetchGroupName(notification.objectId)
            .then(groupName => setGroupName(groupName));
    }, [notification.senderId, notification.objectId]);

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
                    {username} invited you to join {groupName}
                </div>
                <div className="col-auto d-flex align-items-center">
                    <button onClick={() => handleNotificationResponse("confirm")}>&#10003;</button>
                    <button onClick={() => handleNotificationResponse("deny")}>&#10007;</button>
                </div>
            </div>
        </div>
    );
}
