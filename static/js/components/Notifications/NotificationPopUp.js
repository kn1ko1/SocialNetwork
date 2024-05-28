export const NotificationPopUp = ({
  message,
  onClose
}) => {
  return /*#__PURE__*/React.createElement("div", {
    className: "toast show position-fixed bottom-0 end-0 p-3 m-3",
    style: {
      zIndex: 1000
    }
  }, /*#__PURE__*/React.createElement("div", {
    className: "toast-header"
  }, /*#__PURE__*/React.createElement("strong", {
    className: "me-auto"
  }, "Notification"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn-close",
    "aria-label": "Close",
    onClick: onClose
  })), /*#__PURE__*/React.createElement("div", {
    className: "toast-body"
  }, message));
};