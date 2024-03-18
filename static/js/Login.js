import React, { useState } from 'react';
export default function Login(props) {
  const [usernameOrEmail, setUsernameOrEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirectVar, setRedirectVar] = useState(false);
  const [error, setError] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const errorMessage = document.querySelector(".error-message");

  //this is the sign in button
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    //this is user input 
    const userToLogin = {
      usernameOrEmail,
      password
    };
    try {
      //check credentials with backend
      const response = await fetch('http://localhost:8080/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        credentials: 'include',
        body: JSON.stringify(userToLogin)
      });
      if (!response.ok) {
        errorMessage.innerHTML = 'Invalid credentials';
        throw new Error('Invalid credentials');
      }

      //takes response from backend and processes
      const data = await response.json();
      if (data.success) {
        setIsLoggedIn(true);
      } else {
        errorMessage.innerHTML = 'Invalid credentials';
        throw new Error('Invalid credentials');
      }
    } catch (error) {
      errorMessage.innerHTML = 'Invalid credentials';
      setError('Invalid credentials');
    }
  };

  //if credentials frontend match backend then we render home
  if (isLoggedIn) {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render(completedHomePage(), appContainer);
  }

  //this is the register button, when pressed will serve registration form
  const renderRegister = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Register, null), appContainer);
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("main", {
    className: "form-signin w-100 m-auto",
    style: {
      display: "block"
    }
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal"
  }, "Please sign in"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    type: "email",
    className: "form-control",
    id: "floatingInput",
    placeholder: "name@example.com",
    onChange: e => setUsernameOrEmail(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "floatingInput"
  }, "Email address")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    type: "password",
    className: "form-control",
    id: "floatingPassword",
    placeholder: "Password",
    onChange: e => setPassword(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "floatingPassword"
  }, "Password")), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary",
    type: "submit"
  }, "Sign in")), /*#__PURE__*/React.createElement("div", {
    className: "error-message"
  }), /*#__PURE__*/React.createElement("span", null, "Don't have an account? \xA0"), /*#__PURE__*/React.createElement("button", {
    className: "btn btn-link",
    onClick: renderRegister
  }, "Register")));
}