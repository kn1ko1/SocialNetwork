const {
  useState,
  useEffect
} = React;
import { fetchUserById } from "../shared/FetchUserById.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";
export function PrivateMessage({
  notification
}) {
  const [username, setUsername] = useState("");
  const {
    senderId,
    body
  } = notification;
  console.log("senderId, body:", senderId, body);
  useEffect(() => {
    fetchUserById(senderId).then(user => setUsername(user.username));
  }, [senderId]);
  return /*#__PURE__*/React.createElement("div", {
    id: "privateMessage",
    style: notificationCardStyle,
    className: "card"
  }, /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col"
  }, username, ":"), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, body)));
}