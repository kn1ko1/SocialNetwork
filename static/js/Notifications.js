export const renderNotifications = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Notifications, null), pageContainer);
};
export function Notifications() {
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Notifications"));
}