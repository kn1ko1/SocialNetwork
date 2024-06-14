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
  const handleChange = e => {
    const {
      name,
      value,
      type,
      checked
    } = e.target;
    setFormValues(prevValues => ({
      ...prevValues,
      [name]: type === "checkbox" ? checked : value
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
      if (!response.ok) {
        throw new Error("Invalid credentials");
      }
      const data = await response.json();
      if (data.success) {
        setIsRegistered(true);
      } else {
        throw new Error("Invalid credentials");
      }
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
    className: "container login-container"
  }, /*#__PURE__*/React.createElement("h1", {
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
    className: "form-control",
    id: "floatingInput",
    name: "email",
    placeholder: "name@example.com",
    onChange: handleChange
  })), /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "regpassword",
    className: "form-label"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "password",
    className: "form-control",
    id: "regpassword",
    name: "password",
    placeholder: "Password",
    onChange: handleChange
  }))), /*#__PURE__*/React.createElement("div", {
    className: "row mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "firstName",
    className: "form-label"
  }, "First Name"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: "form-control",
    id: "firstName",
    name: "firstName",
    placeholder: "John",
    onChange: handleChange
  })), /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "lastName",
    className: "form-label"
  }, "Last Name"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: "form-control",
    id: "lastName",
    name: "lastName",
    placeholder: "Doe",
    onChange: handleChange
  }))), /*#__PURE__*/React.createElement("div", {
    className: "row mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "dob",
    className: "form-label"
  }, "Date of Birth"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "date",
    className: "form-control",
    id: "dob",
    name: "dob",
    placeholder: "16/01/1998",
    onChange: handleChange
  })), /*#__PURE__*/React.createElement("div", {
    className: "col-md-6"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "username",
    className: "form-label"
  }, "Username"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "username",
    name: "username",
    placeholder: "Johnny",
    onChange: handleChange
  }))), /*#__PURE__*/React.createElement("div", {
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
    class: "btn btn-primary rounded",
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

  // 	<div className="container login-container">
  // 		<h1 className="h3 mb-3 fw-normal login-text">Register</h1>
  // 		<form onSubmit={submit}>
  // 			<div className="mb-3">
  // 				<label htmlFor="floatingInput">Email address</label>
  // 				<input
  // 					required
  // 					type="email"
  // 					className="form-control"
  // 					id="floatingInput"
  // 					name="email"
  // 					placeholder="name@example.com"
  // 					onChange={handleChange}
  // 				/>
  // 			</div>

  // 			<div className="mb-3">
  // 				<label htmlFor="regpassword">Password</label>
  // 				<input
  // 					required
  // 					type="password"
  // 					className="form-control reginput"
  // 					id="regpassword"
  // 					name="password"
  // 					placeholder="Password"
  // 					onChange={handleChange}
  // 				/>
  // 			</div>

  // 			<div className="mb-3">
  // 				<label htmlFor="firstName">First Name</label>
  // 				<input
  // 					required
  // 					type="text"
  // 					className="form-control reginput"
  // 					id="firstName"
  // 					name="firstName"
  // 					placeholder="John"
  // 					onChange={handleChange}
  // 				/>
  // 			</div>

  // 			<div className="mb-3">
  // 				<label htmlFor="lastName">Last Name</label>
  // 				<input
  // 					required
  // 					type="text"
  // 					className="form-control reginput"
  // 					id="lastName"
  // 					name="lastName"
  // 					placeholder="Doe"
  // 					onChange={handleChange}
  // 				/>
  // 			</div>

  // 			<div className="mb-3">
  // 				<label htmlFor="dob">Date of Birth</label>
  // 				<input
  // 					required
  // 					type="date"
  // 					className="form-control reginput"
  // 					id="dob"
  // 					name="dob"
  // 					placeholder="16/01/1998"
  // 					onChange={handleChange}
  // 				/>
  // 			</div>

  // 			<div className="mb-3">
  // 				<label htmlFor="image">Avatar Image (optional)</label>
  // 				<button
  // 					type="button"
  // 					className="btn btn-primary"
  // 					onClick={handleSelectFile}
  // 				>
  // 					Select File
  // 				</button>
  // 				<span>{selectedFile ? selectedFile.name : "No file selected"}</span>
  // 				<input
  // 					type="file"
  // 					id="fileInput"
  // 					accept="image/*"
  // 					style={{ display: "none" }}
  // 					onChange={handleFileChange}
  // 				/>
  // 			</div>

  // 			<div className="mb-3">
  // 				<label htmlFor="username">Username</label>
  // 				<input
  // 					type="text"
  // 					className="form-control reginput"
  // 					id="username"
  // 					name="username"
  // 					placeholder="Johnny"
  // 					onChange={handleChange}
  // 				/>
  // 			</div>

  // 			<div className="form-check">
  // 				<input
  // 					className="form-check-input"
  // 					type="radio"
  // 					id="public-status"
  // 					name="isPublic"
  // 					value={true}
  // 					checked={formValues.isPublic === true}
  // 					onChange={handleChange}
  // 				/>
  // 				<label className="form-check-label" htmlFor="public-status">
  // 					Public
  // 				</label>
  // 			</div>

  // 			<div className="form-check">
  // 				<input
  // 					className="form-check-input"
  // 					type="radio"
  // 					id="private-status"
  // 					name="isPublic"
  // 					value={false}
  // 					checked={formValues.isPublic === false}
  // 					onChange={handleChange}
  // 				/>
  // 				<label className="form-check-label" htmlFor="private-status">
  // 					Private
  // 				</label>
  // 			</div>

  // 			<div className="mb-3">
  // 				<label htmlFor="about">About me (optional)</label>
  // 				<input
  // 					type="text"
  // 					className="form-control reginput"
  // 					id="bio"
  // 					name="bio"
  // 					placeholder="About Me"
  // 					cols="30"
  // 					rows="10"
  // 					onChange={handleChange}
  // 				/>
  // 			</div>

  // 			<button className="btn btn-primary" type="submit">
  // 				Register
  // 			</button>
  // 		</form>
  // 		<div className="error-message"></div>
  // 		<br />
  // 		<div className="mb3">
  // 			<span className="login-text">Already have an account? &nbsp;</span>
  // 			<button type="submit" className="btn btn-primary" onClick={renderLogin}>
  // 				Log in
  // 			</button>
  // 		</div>
  // 	</div>
  // );
}