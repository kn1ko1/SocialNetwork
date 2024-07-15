const { useState, useEffect } = React;
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


export const renderNavbar = ({ socket }) => {
  const navContainer = document.querySelector(".nav-container");
  ReactDOM.render(<Navbar socket={socket} />, navContainer);
};
export const Navbar = ({ socket }) => {
  const { currentUserId, isLoading, error } = getCurrentUserId();
  const [notifications, setNotifications] = useState(null);
  const [notificationData, setNotificationData] = useState(null);
  const [username, setUsername] = useState("");

  useEffect(() => {
    if (currentUserId) {
      fetchUsername(currentUserId)
        .then((username) => setUsername(username))
        .catch((error) => console.error("Error fetching username:", error));
    }
  }, [currentUserId]);

  useEffect(() => {
    const handleSocketMessage = (e) => {
      let data = JSON.parse(e.data);
      setNotificationData(data);
    };

    socket.addEventListener("message", handleSocketMessage);

    return () => {
      socket.removeEventListener("message", handleSocketMessage);
    };
  }, [socket]);

  const fetchNotifications = () => {
    fetch(`http://localhost:8080/api/users/${currentUserId}/notifications`)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch notifications");
        }
        return response.json();
      })
      .then((data) => {
        setNotifications(data);
      })
      .catch((error) => {
        console.error("Error fetching notifications:", error);
      });
  };

  const handleNotificationResponse = (notificationId) => {
    const updatedNotifications = notifications.filter(
      (notification) => notification.notificationId !== notificationId
    );
    setNotifications(updatedNotifications);
  };

  const fetchUsername = async (userId) => {
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
        credentials: "include",
      });

      if (response.ok) {
        socket.close();
        socket.addEventListener("close", (event) => {
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
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <nav className="navbar navbar-expand-md bg-body-tertiary">
      <div className="container-fluid">
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarSupportedContent">
          {notificationData && (
            <NotificationPopUp
              data={notificationData}
              onClose={() => setNotificationData(null)}
              socket={socket}
            />
          )}

          <div className="navbar-logo">
            <img
              src="../../static/sphere-logo.png"
              alt="Logo"
              className="logo"
              style={{ width: "60px", height: "auto" }}
            />
            {username && <span className="ms-2">Welcome, {username}</span>}
          </div>

          <ul className="navbar-nav me-auto mx-auto mb-2 mb-lg-0">
            <li className="nav-item">
              <a
                className="nav-link"
                href="#"
                onClick={() => renderHome({ socket })}
              >
                HOME
              </a>
            </li>

            <li className="nav-item">
              <a
                className="nav-link"
                href="#"
                onClick={() => renderProfile(socket, currentUserId, true)}
              >
                PROFILE
              </a>
            </li>

            <li className="nav-item dropdown">
              <a
                className="nav-link dropdown-toggle"
                href="#"
                id="notificationsDropdown"
                role="button"
                data-bs-toggle="dropdown"
                aria-expanded="false"
                onClick={() => fetchNotifications()}
              >
                NOTIFICATIONS
              </a>
              <ul className="dropdown-menu" aria-labelledby="notificationsDropdown" style={{ minWidth: '500px' }}>
                {notifications !== null && Object.keys(notifications).length > 0 ? (
                  Object.values(notifications).map((notification, index) => (
                    <li key={index}>
                      {notification.notificationType === "groupInvite" && <GroupInvite
                        notification={notification}
                        onNotificationResponse={handleNotificationResponse}
                        socket={socket}
                      />}
                      {notification.notificationType === "groupRequest" && <GroupRequest
                        notification={notification}
                        onNotificationResponse={handleNotificationResponse}
                        socket={socket}
                      />}
                      {notification.notificationType === "eventInvite" && <EventInvite
                        notification={notification}
                        onNotificationResponse={handleNotificationResponse}
                      />}
                      {notification.notificationType === "followRequest" && <FollowRequest
                        notification={notification}
                        onNotificationResponse={handleNotificationResponse}
                      />}
                    </li>
                  ))
                ) : (
                  <li>No notifications</li>
                )}
              </ul>
            </li>

            <li className="nav-item">
              <a
                className="nav-link"
                href="#"
                onClick={() => renderChat({ socket })}
              >
                CHAT
              </a>
            </li>

            <li className="nav-item">
              <a
                className="nav-link"
                href="#"
                onClick={() => renderGroup({ socket })}
              >
                GROUP
              </a>
            </li>

            <li className="nav-item">
              <a
                className="nav-link"
                href="#"
                onClick={() => logout({ socket })}
              >
                LOGOUT
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};

// export default Navbar;
