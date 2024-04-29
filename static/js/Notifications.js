const {
  useState,
  useEffect
} = React;
import { useSocket } from "./shared/UserProvider.js";
export const renderNotifications = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Notifications, null), pageContainer);
};
export function Notifications() {
  const {
    currentUserId
  } = useSocket();
  const [notifications, setNotifications] = useState({});
  useEffect(() => {
    if (currentUserId !== null) {
      fetchNotifications();
    }
  }, [currentUserId]);
  const fetchNotifications = () => {
    fetch(`http://localhost:8080/api/users/${currentUserId}/notifications`).then(response => response.json()).then(data => {
      setNotifications(data);
    }).catch(error => {
      console.error("Error fetching notifications data:", error);
    });
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Notifications"), notifications !== null && Object.keys(notifications).length > 0 ? /*#__PURE__*/React.createElement("ul", null, Object.values(notifications).map((notification, index) => /*#__PURE__*/React.createElement("li", {
    key: index
  }, /*#__PURE__*/React.createElement(GroupInvite, {
    notification: notification
  })))) : /*#__PURE__*/React.createElement("div", null, "No notifications"));
}
function GroupInvite({
  notification
}) {
  return /*#__PURE__*/React.createElement("div", {
    id: "GroupInvite",
    className: "card",
    style: {
      maxWidth: "400px"
    }
  }, "User ", notification.senderId, " invited you to join Group ", notification.objectId);
}
function GroupRequest({
  notification
}) {
  return /*#__PURE__*/React.createElement("div", {
    id: "GroupRequest",
    className: "card",
    style: {
      maxWidth: "400px"
    }
  }, "User ", notification.senderId, " has requested to join Group ", notification.objectId);
}