const {
  useState,
  useEffect
} = React;
import { fetchUsername } from "./FetchUsername.js";
import { fetchGroupName } from "./FetchGroupName.js";
import { respondToNotification } from "./RespondToNotification.js";
export function GroupRequest({
  notification
}) {
  const [username, setUsername] = useState("");
  const [groupName, setGroupName] = useState("");
  useEffect(() => {
    fetchUsername(notification.senderId).then(username => setUsername(username));
    fetchGroupName(notification.objectId).then(groupName => setGroupName(groupName));
  }, [notification.senderId, notification.objectId]);
  return /*#__PURE__*/React.createElement("div", {
    id: notification.notificationType,
    className: "card",
    style: {
      maxWidth: "400px"
    }
  }, username, " requested to join ", groupName, /*#__PURE__*/React.createElement("button", {
    onClick: () => respondToNotification("confirm", notification)
  }, "\u2713"), /*#__PURE__*/React.createElement("button", {
    onClick: () => respondToNotification("deny", notification)
  }, "\u2717"));
}