const {
  useState,
  useEffect
} = React;
//import ReactDOM from "react-dom";
import { getCurrentUserId } from "./GetCurrentUserId.js";
import { renderProfile } from "../../Profile.js";
import { renderHome } from "../../Home.js";
//import { renderNotifications } from "../../Notifications.js";
import { renderChat } from "../../Chat.js";
import { renderGroup } from "../../Group.js";
import { renderLogin } from "../../Login.js";
import { NotificationPopUp } from "../Notifications/NotificationPopUp.js";
import { GroupInvite } from "../../components/Notifications/GroupInvite.js";
import { GroupRequest } from "../../components/Notifications/GroupRequest.js";
import { FollowRequest } from "../../components/Notifications/FollowRequest.js";
import { EventInvite } from "../../components/Notifications/EventInvite.js";
export const renderNavbar = ({
  socket
}) => {
  const navContainer = document.querySelector(".nav-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Navbar, {
    socket: socket
  }), navContainer);
};
export const Navbar = ({
  socket
}) => {
  const {
    currentUserId,
    isLoading,
    error
  } = getCurrentUserId();
  const [notifications, setNotifications] = useState(null);
  const [notificationData, setNotificationData] = useState(null);
  const [username, setUsername] = useState("");
  useEffect(() => {
    if (currentUserId) {
      fetchUsername(currentUserId).then(username => setUsername(username)).catch(error => console.error("Error fetching username:", error));
    }
  }, [currentUserId]);
  useEffect(() => {
    const handleSocketMessage = e => {
      let data = JSON.parse(e.data);
      setNotificationData(data);
    };
    socket.addEventListener("message", handleSocketMessage);
    return () => {
      socket.removeEventListener("message", handleSocketMessage);
    };
  }, [socket]);
  const fetchNotifications = () => {
    fetch(`http://localhost:8080/api/users/${currentUserId}/notifications`).then(response => {
      if (!response.ok) {
        throw new Error("Failed to fetch notifications");
      }
      return response.json();
    }).then(data => {
      setNotifications(data);
    }).catch(error => {
      console.error("Error fetching notifications:", error);
    });
  };
  const handleNotificationResponse = notificationId => {
    const updatedNotifications = notifications.filter(notification => notification.notificationId !== notificationId);
    setNotifications(updatedNotifications);
  };
  const fetchUsername = async userId => {
    if (!userId) {
      throw new Error("Invalid userId");
    }
    try {
      const response = await fetch(`http://localhost:8080/api/users/${userId}`);
      if (!response.ok) {
        throw new Error("Error fetching user data");
      }
      const data = await response.json();
      return data.username;
    } catch (error) {
      console.error("Error fetching user data:", error);
      throw error;
    }
  };
  const logout = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/logout", {
        method: "POST",
        credentials: "include"
      });
      if (response.ok) {
        socket.close();
        socket.addEventListener("close", event => {
          console.log("The connection has been closed successfully.");
        });
        renderLogin();
        const navContainer = document.querySelector(".nav-container");
        ReactDOM.render(null, navContainer);
        console.log("Logout successful!");
      } else {
        console.log("Failed to logout. Server response not OK.");
      }
    } catch (error) {
      console.error("An error occurred during logout:", error);
    }
  };
  if (isLoading) {
    return /*#__PURE__*/React.createElement("div", null, "Loading...");
  }
  if (error) {
    return /*#__PURE__*/React.createElement("div", null, "Error: ", error);
  }
  return /*#__PURE__*/React.createElement("nav", {
    className: "navbar navbar-expand-md bg-body-tertiary"
  }, /*#__PURE__*/React.createElement("div", {
    className: "container-fluid"
  }, /*#__PURE__*/React.createElement("button", {
    className: "navbar-toggler",
    type: "button",
    "data-bs-toggle": "collapse",
    "data-bs-target": "#navbarSupportedContent",
    "aria-controls": "navbarSupportedContent",
    "aria-expanded": "false",
    "aria-label": "Toggle navigation"
  }, /*#__PURE__*/React.createElement("span", {
    className: "navbar-toggler-icon"
  })), /*#__PURE__*/React.createElement("div", {
    className: "collapse navbar-collapse",
    id: "navbarSupportedContent"
  }, notificationData && /*#__PURE__*/React.createElement(NotificationPopUp, {
    data: notificationData,
    onClose: () => setNotificationData(null)
  }), /*#__PURE__*/React.createElement("div", {
    className: "navbar-logo"
  }, /*#__PURE__*/React.createElement("img", {
    src: "../../static/sphere-logo.png",
    alt: "Logo",
    className: "logo",
    style: {
      width: "60px",
      height: "auto"
    }
  }), username && /*#__PURE__*/React.createElement("span", {
    className: "ms-2"
  }, "Welcome, ", username)), /*#__PURE__*/React.createElement("ul", {
    className: "navbar-nav me-auto mx-auto mb-2 mb-lg-0"
  }, /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderHome({
      socket
    })
  }, "HOME")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderProfile(socket, currentUserId, true)
  }, "PROFILE")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item dropdown"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link dropdown-toggle",
    href: "#",
    id: "notificationsDropdown",
    role: "button",
    "data-bs-toggle": "dropdown",
    "aria-expanded": "false",
    onClick: () => fetchNotifications()
  }, "NOTIFICATIONS"), /*#__PURE__*/React.createElement("ul", {
    className: "dropdown-menu",
    "aria-labelledby": "notificationsDropdown",
    style: {
      minWidth: '500px'
    }
  }, notifications !== null && Object.keys(notifications).length > 0 ? Object.values(notifications).map((notification, index) => /*#__PURE__*/React.createElement("li", {
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
  }))) : /*#__PURE__*/React.createElement("li", null, "No notifications"))), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderChat({
      socket
    })
  }, "CHAT")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderGroup({
      socket
    })
  }, "GROUP")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => logout({
      socket
    })
  }, "LOGOUT"))))));
};

// export default Navbar;