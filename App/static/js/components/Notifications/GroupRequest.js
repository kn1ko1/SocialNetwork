const {
  useState,
  useEffect
} = React;
import { fetchGroupById } from "../shared/FetchGroupById.js";
import { fetchUserById } from "../shared/FetchUserById.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
import { respondToNotification } from "./RespondToNotification.js";
import { websocketRespondToGroupNotification } from "./WebsocketRespondToGroupNotification.js";
export function GroupRequest({
  notification,
  onNotificationResponse,
  socket
}) {
  const [username, setUsername] = useState("");
  const [groupName, setGroupName] = useState("");
  useEffect(() => {
    fetchUserById(notification.senderId).then(user => setUsername(user.username));
    fetchGroupById(notification.objectId).then(group => setGroupName(group.title));
  }, [notification.senderId, notification.objectId]);
  const handleNotificationResponse = async responseType => {
    // Call the respondToNotification function to handle the response
    respondToNotification(responseType, notification);
    if (responseType == "confirm") {
      websocketRespondToGroupNotification(notification, socket);
    }
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
  }, username, " requested to join ", groupName), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("confirm")
  }, "\u2713"), /*#__PURE__*/React.createElement("button", {
    onClick: () => handleNotificationResponse("deny")
  }, "\u2717"))));
}