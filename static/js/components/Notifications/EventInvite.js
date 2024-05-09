const {
  useState,
  useEffect
} = React;
import { fetchUsername } from "../shared/FetchUsername.js";
import { respondToNotification } from "./RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
import { formattedDate } from "../shared/FormattedDate.js";
export function EventInvite({
  notification,
  onNotificationResponse
}) {
  const [username, setUsername] = useState("");
  const [event, setEvent] = useState({});
  const dateTime = formattedDate(event.dateTime);
  useEffect(() => {
    fetchUsername(notification.senderId).then(username => setUsername(username));
    fetchEvent(notification.objectId).then(event => setEvent(event));
  }, [notification.senderId, notification.objectId]);
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
  const handleNotificationResponse = async responseType => {
    // Call the respondToNotification function to handle the response
    respondToNotification(responseType, notification);
    // Call the parent component's callback to remove this notification
    onNotificationResponse(notification.notificationId);
  };
  return /*#__PURE__*/React.createElement("div", {
    id: notification.notificationType,
    style: notificationCardStyle,
    className: "card"
  }, /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col"
  }, username, " invited you to join ", event.title, " at ", dateTime, /*#__PURE__*/React.createElement("div", {
    id: "description"
  }, " \"", event.description, "\"")), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("confirm")
  }, "\u2713"), /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("deny")
  }, "\u2717"))));
}