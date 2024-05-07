const {
  useState,
  useEffect
} = React;
import { fetchUsername } from "../FetchUsername.js";
import { respondToNotification } from "../RespondToNotification.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
export function FollowRequest({
  notification
}) {
  const [username, setUsername] = useState("");
  useEffect(() => {
    fetchUsername(notification.senderId).then(username => setUsername(username));
  }, [notification.senderId]);
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
    onClick: () => respondToNotification("confirm", notification)
  }, "\u2713"), /*#__PURE__*/React.createElement("button", {
    onClick: () => respondToNotification("deny", notification)
  }, "\u2717"))));
}