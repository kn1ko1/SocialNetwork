const {
  useState,
  useEffect
} = React;
import { fetchGroupName } from "../shared/FetchGroupName.js";
import { fetchUsername } from "../shared/FetchUsername.js";
import { respondToNotification } from "./RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
export function GroupInvite({
  notification,
  onNotificationResponse
}) {
  const [username, setUsername] = useState("");
  const [groupName, setGroupName] = useState("");
  useEffect(() => {
    fetchUsername(notification.senderId).then(username => setUsername(username));
    fetchGroupName(notification.objectId).then(groupName => setGroupName(groupName));
  }, [notification.senderId, notification.objectId]);
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
  }, username, " invited you to join ", groupName), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("confirm")
  }, "\u2713"), /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("deny")
  }, "\u2717"))));
}