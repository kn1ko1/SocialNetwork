const {
  useState,
  useEffect
} = React;
import { useSocket } from "./app.js";
import { renderNavbar } from "./components/Navbar.js";
import { renderRegister } from "./Register.js";
import { renderHome } from "./Home.js";
export const renderLogin = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Login, null), pageContainer);
};
export function Login() {
  const {
    updateContext
  } = useSocket();
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
      const newSocket = new WebSocket("ws://localhost:8080/ws");
      newSocket.onopen = function (event) {
        console.log("WebSocket connection established.");
      };
      const fetchUserId = async () => {
        try {
          const response = await fetch("http://localhost:8080/api/userId", {
            credentials: "include"
          });
          if (response.ok) {
            const userId = await response.json();
            updateContext(newSocket, userId);
            renderHome();
            // Render the Navbar after successful login
            renderNavbar();
          } else {
            setErrorMessage("Failed to fetch userId");
            console.error("Response not okay:", response.status);
          }
        } catch (error) {
          setErrorMessage("Error fetching userId");
          console.error("Fetch error:", error);
        }
      };
      fetchUserId();
    }
  }, [isLoggedIn]);
  return /*#__PURE__*/React.createElement("div", {
    className: "container login-container"
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal login-text"
  }, "Log in"), /*#__PURE__*/React.createElement("form", {
    onSubmit: handleSubmit
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleInputEmail1",
    className: "form-label"
  }, "Email address"), /*#__PURE__*/React.createElement("input", {
    type: "email",
    className: "form-control form-control-lg",
    id: "exampleInputEmail1",
    "aria-describedby": "emailHelp",
    onChange: handleUsernameOrEmailChange
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleInputPassword1",
    className: "form-label"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    type: "password",
    className: "form-control form-control-lg",
    id: "exampleInputPassword1",
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