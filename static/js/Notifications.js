const {
  useState,
  useEffect
} = React;
import { getCurrentUserId } from "./shared/getCurrentUserId.js";
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
  }, /*#__PURE__*/React.createElement(GroupInvite, {
    notification: notification
  })))) : /*#__PURE__*/React.createElement("div", null, "No notifications"));
}
function GroupInvite({
  notification
}) {
  const [username, setUsername] = useState("");
  const [groupName, setGroupName] = useState("");
  useEffect(() => {
    fetchUsername();
    fetchGroupName();
  }, []);
  const fetchUsername = () => {
    fetch(`http://localhost:8080/api/users/${notification.senderId}`).then(response => response.json()).then(data => {
      setUsername(data.username);
    }).catch(error => {
      console.error("Error fetching notifications data:", error);
    });
  };
  const fetchGroupName = () => {
    fetch(`http://localhost:8080/api/groups/${notification.objectId}`).then(response => response.json()).then(data => {
      setGroupName(data.title);
    }).catch(error => {
      console.error("Error fetching notifications data:", error);
    });
  };
  const respondToNotification = (reply, notification) => {
    const data = {
      reply: reply,
      notification: notification
    };
    fetch(`http://localhost:8080/api/notifications/${notification.notificationId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }).then(response => response.json()).then(data => {
      // Handle success response
      console.log("Response sent successfully:", data);
    }).catch(error => {
      console.error("Error sending response:", error);
    });
  };
  return /*#__PURE__*/React.createElement("div", {
    id: "GroupInvite",
    className: "card",
    style: {
      maxWidth: "400px"
    }
  }, username, " invited you to join ", groupName, /*#__PURE__*/React.createElement("button", {
    onClick: () => respondToNotification("confirm", notification)
  }, "\u2713"), /*#__PURE__*/React.createElement("button", {
    onClick: () => respondToNotification("deny", notification)
  }, "\u2717"));
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
  }, "User ", notification.senderId, " has requested to join ", notification.objectId);
}