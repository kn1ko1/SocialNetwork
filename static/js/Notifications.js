const {
  useState,
  useEffect
} = React;
import { getCurrentUserId } from "./shared/getCurrentUserId.js";
import { GroupInvite } from "./components/Notifications/GroupInvite.js";
import { GroupRequest } from "./components/Notifications/GroupRequest.js";
import { FollowRequest } from "./components/Notifications/FollowRequest.js";
import { EventInvite } from "./components/Notifications/EventInvite.js";
export const renderNotifications = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Notifications, null), pageContainer);
};
export function Notifications() {
  const {
    currentUserId
  } = getCurrentUserId();
  const [notifications, setNotifications] = useState(null);
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
  }, notification.notificationType === "groupInvite" && /*#__PURE__*/React.createElement(GroupInvite, {
    notification: notification
  }), notification.notificationType === "groupRequest" && /*#__PURE__*/React.createElement(GroupRequest, {
    notification: notification
  }), notification.notificationType === "eventInvite" && /*#__PURE__*/React.createElement(EventInvite, {
    notification: notification
  }), notification.notificationType === "followRequest" && /*#__PURE__*/React.createElement(FollowRequest, {
    notification: notification
  })))) : /*#__PURE__*/React.createElement("div", null, "No notifications"));
}