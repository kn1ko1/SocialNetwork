const {
  useState,
  useEffect
} = React;
import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
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
    console.log("notifications:", data);
  };
  const handleNotificationResponse = notificationId => {
    console.log("notificationId", notificationId);
    // Filter out the notification with the given ID from notifications state
    const updatedNotifications = notifications.filter(notification => notification.notificationId !== notificationId);
    // Update notifications state with the filtered notifications
    setNotifications(updatedNotifications);
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Notifications"), notifications !== null && Object.keys(notifications).length > 0 ? /*#__PURE__*/React.createElement("ul", null, Object.values(notifications).map((notification, index) => /*#__PURE__*/React.createElement("li", {
    key: index
  }, notification.notificationType === "groupInvite" && /*#__PURE__*/React.createElement(GroupInvite, {
    notification: notification,
    onNotificationResponse: handleNotificationResponse
  }), notification.notificationType === "groupRequest" && /*#__PURE__*/React.createElement(GroupRequest, {
    notification: notification,
    onNotificationResponse: handleNotificationResponse
  }), notification.notificationType === "eventInvite" && /*#__PURE__*/React.createElement(EventInvite, {
    notification: notification,
    onNotificationResponse: handleNotificationResponse
  }), notification.notificationType === "followRequest" && /*#__PURE__*/React.createElement(FollowRequest, {
    notification: notification,
    onNotificationResponse: handleNotificationResponse
  })))) : /*#__PURE__*/React.createElement("div", null, "No notifications"));
}