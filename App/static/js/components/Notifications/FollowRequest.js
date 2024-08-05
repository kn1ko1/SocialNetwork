const {
  useState,
  useEffect
} = React;
import { fetchUserById } from "../shared/FetchUserById.js";
import { respondToNotification } from "./RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
export function FollowRequest({
  notification,
  onNotificationResponse
}) {
  const [username, setUsername] = useState("");
  useEffect(() => {
    fetchUserById(notification.senderId).then(user => setUsername(user.username));
  }, [notification.senderId]);
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
  }, username, " has requested to follow you"), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("confirm")
  }, "\u2713"), /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("deny")
  }, "\u2717"))));
}