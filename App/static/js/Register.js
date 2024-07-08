const {
  useState
} = React;
import { initializeSocket } from "./app.js";
import { renderNavbar } from "./components/shared/Navbar.js";
import { renderHome } from "./Home.js";
import { renderLogin } from "./Login.js";
export const renderRegister = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Register, null), pageContainer);
};
export function Register() {
  const [formValues, setFormValues] = useState({
    email: "",
    password: "",
    firstName: "",
    lastName: "",
    dob: "",
    username: "",
    bio: "",
    isPublic: true
  });
  const [selectedFile, setSelectedFile] = useState(null);
  const [isRegistered, setIsRegistered] = useState(false);
  const [errors, setErrors] = useState({});
  const validate = () => {
    const errors = {};
    const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    const dobDate = new Date(formValues.dob);
    const currentDate = new Date();
    if (!emailPattern.test(formValues.email)) {
      errors.email = "Please enter a valid email address following the format: email@example.com";
    }
    if (!formValues.password) {
      errors.password = "Password is required.";
    }
    if (!formValues.firstName) {
      errors.firstName = "First name is required.";
    }
    if (!formValues.lastName) {
      errors.lastName = "Last name is required.";
    }
    if (isNaN(dobDate.getTime()) || dobDate >= currentDate) {
      errors.dob = "Please enter a valid Date of Birth in the past.";
    }
    if (formValues.username.length < 4 || formValues.username.length > 50) {
      errors.username = "Username must be between 4 and 50 characters.";
    }
    return errors;
  };
  const handleChange = e => {
    const {
      name,
      value,
      type,
      checked
    } = e.target;
    setFormValues(prevValues => ({
      ...prevValues,
      [name]: type === "checkbox" ? checked : type === "radio" ? JSON.parse(value) : value
    }));
  };
  const handleFileChange = e => {
    setSelectedFile(e.target.files[0]);
  };
  const handleSelectFile = () => {
    const fileInput = document.getElementById("fileInput");
    fileInput.click();
  };
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    const validationErrors = validate();
    if (Object.keys(validationErrors).length > 0) {
      setErrors(validationErrors);
      return;
    }
    try {
      const formData = new FormData();
      Object.keys(formValues).forEach(key => {
        formData.append(key, formValues[key]);
      });
      if (selectedFile) {
        formData.append("image", selectedFile);
      }
      console.log("formData:", formData);
      // Send user data to backend
      const response = await fetch("http://localhost:8080/auth/registration", {
        method: "POST",
        body: formData
      });
      console.log("Response:", response);
      if (!response.ok) {
        setErrors({
          server: data.message || "Registration failed"
        });
        throw new Error("Invalid credentials");
      }
      const data = await response.json();
      console.log(data);
      if (data.success) {
        setIsRegistered(true);
      } else {
        console.log(1);
        const validationErrors = {};
        if (data.errorField == "Email") {
          validationErrors.email = data.errorMessage;
        }
        if (data.errorField == "Username") {
          validationErrors.username = data.errorMessage;
        }
        if (Object.keys(validationErrors).length > 0) {
          setErrors(validationErrors);
          return;
        }
        return;
      }

      // } else {
      // 	setErrors({ server: "Invalid credentials" });
      // 	throw new Error("Invalid credentials");
      // }
    } catch (error) {
      console.error("Registration error:", error);
    }
  };
  if (isRegistered) {
    const socket = initializeSocket();
    renderNavbar({
      socket
    });
    renderHome({
      socket
    });
  }
  return /*#__PURE__*/React.createElement("div", {
    className: "container login-container",
    style: {
      minHeight: "100vh",
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
      paddingBottom: "20px"
    }
  }, /*#__PURE__*/React.createElement("div", {
    className: "logo-container"
  }, /*#__PURE__*/React.createElement("img", {
    src: "../static/sphere-logo.png",
    alt: "Logo",
    className: "logo"
  })), /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal login-text"
  }, "Register"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "row mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "floatingInput",
    className: "form-label"
  }, "Email address"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "email",
    className: `form-control ${errors.email ? "is-invalid" : ""}`,
    id: "floatingInput",
    name: "email",
    placeholder: "name@example.com",
    onChange: handleChange
  }), errors.email && /*#__PURE__*/React.createElement("div", {
    className: "invalid-feedback"
  }, errors.email)), /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "regpassword",
    className: "form-label"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "password",
    className: `form-control ${errors.password ? "is-invalid" : ""}`,
    id: "regpassword",
    name: "password",
    placeholder: "Password",
    onChange: handleChange
  }), errors.password && /*#__PURE__*/React.createElement("div", {
    className: "invalid-feedback"
  }, errors.password))), /*#__PURE__*/React.createElement("div", {
    className: "row mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "firstName",
    className: "form-label"
  }, "First Name"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: `form-control ${errors.firstName ? "is-invalid" : ""}`,
    id: "firstName",
    name: "firstName",
    placeholder: "John",
    onChange: handleChange
  }), errors.firstName && /*#__PURE__*/React.createElement("div", {
    className: "invalid-feedback"
  }, errors.firstName)), /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "lastName",
    className: "form-label"
  }, "Last Name"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: `form-control ${errors.lastName ? "is-invalid" : ""}`,
    id: "lastName",
    name: "lastName",
    placeholder: "Doe",
    onChange: handleChange
  }), errors.lastName && /*#__PURE__*/React.createElement("div", {
    className: "invalid-feedback"
  }, errors.lastName))), /*#__PURE__*/React.createElement("div", {
    className: "row mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "dob",
    className: "form-label"
  }, "Date of Birth"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "date",
    className: `form-control ${errors.dob ? "is-invalid" : ""}`,
    id: "dob",
    name: "dob",
    placeholder: "16/01/1998",
    onChange: handleChange
  }), errors.dob && /*#__PURE__*/React.createElement("div", {
    className: "invalid-feedback"
  }, errors.dob)), /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "username",
    className: "form-label"
  }, "Username"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: `form-control ${errors.username ? "is-invalid" : ""}`,
    id: "username",
    name: "username",
    placeholder: "Johnny",
    onChange: handleChange
  }), errors.username && /*#__PURE__*/React.createElement("div", {
    className: "invalid-feedback"
  }, errors.username))), /*#__PURE__*/React.createElement("div", {
    className: "row mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "image",
    className: "form-label"
  }, "Avatar Image (optional)"), /*#__PURE__*/React.createElement("div", {
    className: "input-group"
  }, /*#__PURE__*/React.createElement("input", {
    type: "file",
    id: "fileInput",
    accept: "image/*",
    style: {
      display: "none"
    },
    onChange: handleFileChange
  }), /*#__PURE__*/React.createElement("button", {
    className: "btn btn-primary rounded",
    type: "button",
    onClick: handleSelectFile
  }, "Select File"), /*#__PURE__*/React.createElement("span", {
    className: "input-group-text",
    style: {
      backgroundColor: "transparent",
      border: "none"
    }
  }, selectedFile ? selectedFile.name : "No file selected"))), /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    className: "form-label"
  }, "Profile Visibility"), /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    className: "form-check-input",
    type: "radio",
    id: "public-status",
    name: "isPublic",
    value: true,
    checked: formValues.isPublic === true,
    onChange: handleChange
  }), /*#__PURE__*/React.createElement("label", {
    className: "form-check-label",
    htmlFor: "public-status"
  }, "Public")), /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    className: "form-check-input",
    type: "radio",
    id: "private-status",
    name: "isPublic",
    value: false,
    checked: formValues.isPublic === false,
    onChange: handleChange
  }), /*#__PURE__*/React.createElement("label", {
    className: "form-check-label",
    htmlFor: "private-status"
  }, "Private"))))), /*#__PURE__*/React.createElement("div", {
    className: "row mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-12"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "bio",
    className: "form-label"
  }, "About me (optional)"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "bio",
    name: "bio",
    placeholder: "About Me",
    onChange: handleChange
  }))), /*#__PURE__*/React.createElement("button", {
    className: "btn btn-primary",
    type: "submit"
  }, "Register")), /*#__PURE__*/React.createElement("div", {
    className: "error-message"
  }), /*#__PURE__*/React.createElement("br", null), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("span", {
    className: "login-text"
  }, "Already have an account? \xA0"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-primary",
    onClick: renderLogin
  }, "Log in")));
}