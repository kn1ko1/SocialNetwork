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
      console.log("notifications:", data);
    }).catch(error => {
      console.error("Error fetching notifications data:", error);
    });
  };
  const handleNotificationResponse = notificationId => {
    // Filter out the notification with the given ID from notifications state
    const updatedNotifications = notifications.filter(notification => notification.notificationId !== notificationId);
    // Update notifications state with the filtered notifications
    setNotifications(updatedNotifications);
  };
  const notificationsStyle = {
    maxWidth: '1000px',
    background: 'linear-gradient(to bottom, #c7ddef, #ffffff)',
    // Light blue/grey to white gradient
    borderRadius: '10px',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
    // Optional: Add shadow for depth
    padding: '40px',
    margin: 'auto',
    marginBottom: '20px',
    // Adjust spacing between post cards
    border: '1px solid #ccc' // Add a thin border
  };
  return /*#__PURE__*/React.createElement("div", {
    style: notificationsStyle,
    className: "col-md-4"
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline',
      textAlign: 'center'
    }
  }, "Notifications"), notifications !== null && Object.keys(notifications).length > 0 ? /*#__PURE__*/React.createElement("ul", null, Object.values(notifications).map((notification, index) => /*#__PURE__*/React.createElement("li", {
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