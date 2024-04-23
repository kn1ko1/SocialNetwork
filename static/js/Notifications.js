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
  }, notification.notificationType))) : /*#__PURE__*/React.createElement("div", null, "No notifications"));
}