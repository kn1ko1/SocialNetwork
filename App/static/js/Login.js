const {
  useState,
  useEffect
} = React;
import { initializeSocket } from "./app.js";
import { renderNavbar } from "./components/shared/Navbar.js";
import { renderRegister } from "./Register.js";
import { renderHome } from "./Home.js";
export const renderLogin = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Login, null), pageContainer);
};
export function Login() {
  let socket = null;
  const [usernameOrEmail, setUsernameOrEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const handleUsernameOrEmailChange = e => {
    setUsernameOrEmail(e.target.value);
  };
  const handlePasswordChange = e => {
    setPassword(e.target.value);
  };
  const handleSubmit = async e => {
    e.preventDefault();
    const userToLogin = {
      usernameOrEmail,
      password
    };
    try {
      const response = await fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        credentials: "include",
        body: JSON.stringify(userToLogin)
      });
      if (!response.ok) {
        setErrorMessage("Invalid credentials");
        throw new Error("Invalid credentials");
      }
      const data = await response.json();
      if (data.success) {
        setIsLoggedIn(true);
        setErrorMessage("");
      } else {
        setErrorMessage("Invalid credentials");
        throw new Error("Invalid credentials");
      }
    } catch (error) {
      setErrorMessage("Invalid credentials");
    }
  };
  useEffect(() => {
    if (isLoggedIn) {
      socket = initializeSocket();
      renderNavbar({
        socket
      });
      renderHome({
        socket
      });
    }
  }, [isLoggedIn]);
  return /*#__PURE__*/React.createElement("div", {
    className: "container login-container",
    style: {
      minHeight: "100vh",
      display: "flex",
      justifyContent: "center",
      alignItems: "center"
    }
  }, /*#__PURE__*/React.createElement("div", {
    className: "logo-container"
  }, /*#__PURE__*/React.createElement("img", {
     /* Original image source, pre-docker */
    /* <img src="../sphere-logo.png" alt="Logo" className="logo" /> */
    src: "./uploads/images/sphere-logo.png",
    alt: "Logo",
    className: "logo"
  })), /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal login-text"
  }, "Log in"), /*#__PURE__*/React.createElement("form", {
    onSubmit: handleSubmit
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "usernameOrEmail",
    className: "form-label"
  }, "Email address or Username"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control form-control-lg",
    id: "usernameOrEmail",
    "aria-describedby": "emailHelp",
    onChange: handleUsernameOrEmailChange
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "password",
    className: "form-label"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    type: "password",
    className: "form-control form-control-lg",
    id: "password",
    onChange: handlePasswordChange
  })), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "Log in")), errorMessage && /*#__PURE__*/React.createElement("div", {
    className: "error-message"
  }, errorMessage), /*#__PURE__*/React.createElement("br", null), /*#__PURE__*/React.createElement("div", {
    className: "mb3"
  }, /*#__PURE__*/React.createElement("span", {
    className: "login-text"
  }, "Don't have an account? \xA0"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-primary",
    onClick: renderRegister
  }, "Register")));
}