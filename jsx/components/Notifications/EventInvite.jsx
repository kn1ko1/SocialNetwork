const { useState, useEffect } = React
import { fetchUserById } from "../shared/FetchUserById.js";
import { respondToNotification } from "./RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
import { formattedDate } from "../shared/FormattedDate.js";

export function EventInvite({ notification, onNotificationResponse }) {
    const [username, setUsername] = useState("")
    const [event, setEvent] = useState({})

    const dateTime = formattedDate(event.dateTime)
    useEffect(() => {
        fetchUserById(notification.senderId)
            .then(user => setUsername(user.username));
        fetchEvent(notification.objectId)
            .then(event => setEvent(event));
    }, [notification.senderId, notification.objectId]);
    console.log("notification", notification)
    const fetchEvent = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/events/${notification.objectId}`);
            const data = await response.json();
            return data;
        } catch (error) {
            console.error("Error fetching event:", error);
            return null;
        }
    };

    const handleNotificationResponse = async (responseType) => {
        // Call the respondToNotification function to handle the response
        respondToNotification(responseType, notification);
        // Call the parent component's callback to remove this notification
        if (onNotificationResponse != null) {
            onNotificationResponse(notification.notificationId);
        }

    };

    return (
        <div id={notification.notificationType} style={notificationCardStyle} className="card">
            <div className="row">
                <div className="col">
                    {username} invited you to join {event.title} at {dateTime}
                    <div id="description"> "{event.description}"</div>
                </div>
                <div className="col-auto d-flex align-items-center">
                    <button onClick={() => handleNotificationResponse("confirm")}>&#10003;</button>
                    <button onClick={() => handleNotificationResponse("deny")}>&#10007;</button>
                </div>
            </div>
        </div>
    );


}