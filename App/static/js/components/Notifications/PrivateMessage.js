import { notificationCardStyle } from "./NotificationCardStyle.jsx";
export function PrivateMessage({
  notification
}) {
  return /*#__PURE__*/React.createElement("div", {
    id: "privateMessage",
    style: notificationCardStyle,
    className: "card"
  }, /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col"
  }, notification.senderUsername, ":"), /*#__PURE__*/React.createElement("div", {
    className: "col-auto d-flex align-items-center"
  }, notification.body)));
}