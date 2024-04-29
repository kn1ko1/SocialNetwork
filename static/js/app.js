import { UserProvider } from './shared/UserProvider.js';
import { Login } from "./Login.js";
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
ReactDOM.render( /*#__PURE__*/React.createElement(UserProvider, null, /*#__PURE__*/React.createElement(App, null)), root);