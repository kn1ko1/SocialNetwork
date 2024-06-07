import { EventInvite } from "./EventInvite.js";
import { GroupRequest } from "./GroupRequest.js";
import { GroupInvite } from "./GroupInvite.js";
const codeToHeaderText = {
  1: "Group Chat Message",
  2: "Private Message",
  3: "Create Event",
  4: "Group Request",
  5: "Group Invite",
  6: "Event Invite"
};
export const NotificationPopUp = ({
  data,
  onClose
}) => {
  try {
    const notification = JSON.parse(data.body);
    const code = parseInt(data.code, 10);
    console.log("socket message notification:", notification);

    // Get the header text based on the code
    const headerText = codeToHeaderText[code] || "Notification";
    return /*#__PURE__*/React.createElement("div", {
      id: "notificationPopup"
    }, /*#__PURE__*/React.createElement("div", {
      className: "toast show position-fixed bottom-0 end-0 p-3 m-3",
      style: {
        zIndex: 1000
      }
    }, /*#__PURE__*/React.createElement("div", {
      className: "toast-header"
    }, /*#__PURE__*/React.createElement("strong", {
      className: "me-auto"
    }, headerText), /*#__PURE__*/React.createElement("button", {
      type: "button",
      className: "btn-close",
      "aria-label": "Close",
      onClick: onClose
    })), /*#__PURE__*/React.createElement("div", {
      className: "toast-body"
    }, code === 4 || code === 5 || code === 6 ? /*#__PURE__*/React.createElement(React.Fragment, null, code === 4 && /*#__PURE__*/React.createElement(GroupRequest, {
      notification: notification,
      onNotificationResponse: onClose
    }), code === 5 && /*#__PURE__*/React.createElement(GroupInvite, {
      notification: notification,
      onNotificationResponse: onClose
    }), code === 6 && /*#__PURE__*/React.createElement(EventInvite, {
      notification: notification,
      onNotificationResponse: onClose
    })) : notification)));
  } catch (error) {
    console.error("Error processing notification data:", error);
    return null;
  }
};