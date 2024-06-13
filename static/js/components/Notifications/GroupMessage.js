const {
  useState,
  useEffect
} = React;
import { fetchUserById } from "../shared/FetchUserById.js";
import { fetchGroupById } from "../shared/FetchGroupById.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
export function GroupMessage({
  notification
}) {
  const [username, setUsername] = useState("");
  const [groupName, setGroupName] = useState("");
  const {
    senderId,
    targetId,
    body
  } = notification;
  console.log("senderId, body:", senderId, body);
  useEffect(() => {
    fetchUserById(senderId).then(user => setUsername(user.username));
    fetchGroupById(targetId).then(group => setGroupName(group.title));
  }, [senderId]);
  return /*#__PURE__*/React.createElement("div", {
    id: "GroupMessage",
    style: notificationCardStyle,
    className: "card"
  }, /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col"
  }, groupName, ", ", username, ":"), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, body)));
}