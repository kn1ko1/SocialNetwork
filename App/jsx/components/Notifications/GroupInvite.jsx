const { useState, useEffect } = React
import { fetchGroupById } from "../shared/FetchGroupById.js";
import { fetchUserById } from "../shared/FetchUserById.js";

import { notificationCardStyle } from "./NotificationCardStyle.js";
import { respondToNotification } from "./RespondToNotification.js";
import { websocketRespondToGroupNotification } from "./WebsocketRespondToGroupNotification.js";

export function GroupInvite({ notification, onNotificationResponse, socket }) {
    const [username, setUsername] = useState("");
    const [groupName, setGroupName] = useState("");

    useEffect(() => {
        fetchUserById(notification.senderId)
            .then(user => setUsername(user.username));
        fetchGroupById(notification.objectId)
            .then(group => setGroupName(group.title));
    }, [notification.senderId, notification.objectId]);

    const handleNotificationResponse = async (responseType) => {
        // Call the respondToNotification function to handle the response
        respondToNotification(responseType, notification);

        if (responseType == "confirm"){
            websocketRespondToGroupNotification(notification, socket)
        }
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
