import { getCurrentUserId } from "../shared/getCurrentUserId.js";
import { getSocket } from "../shared/socket.js";
import { renderProfile } from "../Profile.js";
import { renderHome } from "../Home.js";
import { renderNotifications } from "../Notifications.js";
import { renderChat } from "../Chat.js";
import { renderGroup } from "../Group.js";
export const renderNavbar = () => {
  const navContainer = document.querySelector(".nav-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Navbar, null), navContainer);
};
export function Navbar() {
  const {
    currentUserId
  } = getCurrentUserId();
  const logout = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/logout", {
        method: "POST",
        credentials: "include"
      });
      console.log(response);
      if (response.ok) {
        socket = getSocket();
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
  }, /*#__PURE__*/React.createElement("ul", {
    className: "navbar-nav me-auto mx-auto mb-2 mb-lg-0"
  }, /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderProfile(currentUserId, true)
  }, "PROFILE")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderHome
  }, "HOME")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderNotifications
  }, "NOTIFICATIONS")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderChat
  }, "CHAT")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderGroup
  }, "GROUP")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: logout
  }, "LOGOUT"))))));
}