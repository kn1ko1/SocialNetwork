import { Login } from "./Login.js";
let socket = null;
export function initializeSocket() {
  if (!socket) {
    socket = new WebSocket("ws://localhost:8080/ws");
    socket.onopen = function (event) {
      console.log("WebSocket connection established.");
    };
  }
  return socket;
}
export function getSocket() {
  return socket;
}
const App = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement("div", {
    className: "nav-container"
  }), /*#__PURE__*/React.createElement("div", {
    className: "page-container"
  }, /*#__PURE__*/React.createElement(Login, null)));
};
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);