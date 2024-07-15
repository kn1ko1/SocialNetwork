const {
  useState,
  useEffect
} = React;
import { fetchGroupById } from "../shared/FetchGroupById.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
export function GroupMessage({
  notification
}) {
  const [groupName, setGroupName] = useState("");
  const {
    targetId
  } = notification;
  useEffect(() => {
    fetchGroupById(targetId).then(group => setGroupName(group.title));
  }, [notification.messageId]);
  return /*#__PURE__*/React.createElement("div", {
    id: "GroupMessage",
    style: notificationCardStyle,
    className: "card"
  }, /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col"
  }, groupName, ", ", notification.senderUsername, ":"), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, notification.body)));
}