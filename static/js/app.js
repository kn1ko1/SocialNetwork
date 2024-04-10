const {
  useState,
  useEffect
} = React;
let socket;
const App = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement(Login, null));
};

// // Const for getting userId in the frontend
const CurrentUserId = () => {
  const [userId, setUserId] = useState(null);
  useEffect(() => {
    const fetchUserId = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/userId', {
          credentials: 'include'
        });
        console.log("response: ", response);
        if (response.ok) {
          const userId = await response.json();
          setUserId(userId);
        } else {
          console.error('Failed to fetch userId');
        }
      } catch (error) {
        console.error('Error fetching userId:', error);
      }
    };
    fetchUserId();
  }, []);
  return userId; // Return only the userId state value
};
function Navbar() {
  const userId = CurrentUserId();
  const renderHome = () => {
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), appContainer);
  };
  const renderNotifications = () => {
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Notifications, null), appContainer);
  };
  const renderChat = () => {
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Chat, null), appContainer);
  };
  const renderGroup = () => {
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Group, null), appContainer);
  };
  const logout = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/logout", {
        method: "POST",
        credentials: "include"
      });
      if (response.ok) {
        socket.close();
        socket.addEventListener("close", event => {
          console.log("The connection has been closed successfully.");
        });
        const appContainer = document.querySelector(".app-container");
        ReactDOM.render( /*#__PURE__*/React.createElement(Login, null), appContainer);
        console.log("Logout successful!");
      } else {
        console.log("Failed to logout. Server response not OK.");
      }
    } catch (error) {
      console.error("An error occurred during logout:", error);
    }
  };
  return /*#__PURE__*/React.createElement("nav", {
    className: "navbar navbar-expand-md bg-body-tertiary"
  }, /*#__PURE__*/React.createElement("div", {
    className: "container-fluid"
  }, /*#__PURE__*/React.createElement("button", {
    className: "navbar-toggler",
    type: "button",
    "data-bs-toggle": "collapse",
    "data-bs-target": "#navbarSupportedContent",
    "aria-controls": "navbarSupportedContent",
    "aria-expanded": "false",
    "aria-label": "Toggle navigation"
  }, /*#__PURE__*/React.createElement("span", {
    className: "navbar-toggler-icon"
  })), /*#__PURE__*/React.createElement("div", {
    className: "collapse navbar-collapse",
    id: "navbarSupportedContent"
  }, /*#__PURE__*/React.createElement("ul", {
    className: "navbar-nav me-auto mx-auto mb-2 mb-lg-0"
  }, /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderProfile(userId, true)
  }, "PROFILE")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderHome
  }, "HOME")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderNotifications
  }, "NOTIFICATIONS")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderChat
  }, "CHAT")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: renderGroup
  }, "GROUP")), /*#__PURE__*/React.createElement("li", {
    className: "nav-item"
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: logout
  }, "LOGOUT"))))));
}
function Login() {
  const [usernameOrEmail, setUsernameOrEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
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
      const response = await fetch('http://localhost:8080/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        credentials: 'include',
        body: JSON.stringify(userToLogin)
      });
      if (!response.ok) {
        setErrorMessage('Invalid credentials');
        throw new Error('Invalid credentials');
      }
      const data = await response.json();
      if (data.success) {
        setIsLoggedIn(true);
        setErrorMessage('');
      } else {
        setErrorMessage('Invalid credentials');
        throw new Error('Invalid credentials');
      }
    } catch (error) {
      setErrorMessage('Invalid credentials');
    }
  };
  const renderRegister = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Register, null), appContainer);
  };
  if (isLoggedIn) {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), appContainer);
    socket = new WebSocket("ws://localhost:8080/ws");
    socket.onopen = function (event) {
      console.log("WebSocket connection established.");
    };
  }
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
function Register() {
  const [email, setEmail] = useState("");
  const [encryptedPassword, setEncryptedPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [dob, setDob] = useState("");
  const [imageURL, setImageURL] = useState("");
  const [username, setUsername] = useState("");
  const [bio, setBio] = useState("");
  const [isPublic, setIsPublic] = useState(true);
  const [isRegistered, setIsRegistered] = useState(false);
  const handleChange = e => {
    setIsPublic(e.target.value === "true");
  };
  //this is register button
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    // Create new user as JS object.
    const newUser = {
      email,
      encryptedPassword,
      firstName,
      lastName,
      dob,
      imageURL,
      username,
      bio,
      isPublic
    };
    try {
      // Send user data to backend
      const response = await fetch("http://localhost:8080/auth/registration", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(newUser)
      });
      if (!response.ok) {
        throw new Error("Invalid credentials");
      }

      //takes response from backend and processes
      const data = await response.json();
      if (data.success) {
        setIsRegistered(true);
      } else {
        throw new Error("Invalid credentials");
      }
    } catch (error) {
      throw new Error("Invalid credentials");
    }
  };

  //if credentials frontend succesfully create a new user then we render home
  if (isRegistered) {
    socket = new WebSocket("ws://localhost:8080/ws");
    socket.onopen = function (event) {
      console.log("WebSocket connection established.");
    };
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), appContainer);
  }

  //this is the login button, when pressed will serve login form
  const renderLogin = () => {
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Login, null), appContainer);
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "container login-container"
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal login-text"
  }, "register"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "floatingInput"
  }, "Email address"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "email",
    className: "form-control",
    id: "floatingInput",
    placeholder: "name@example.com",
    onChange: e => setEmail(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "regpassword"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "password",
    className: "form-control reginput",
    id: "regpassword",
    placeholder: "Password",
    onChange: e => setEncryptedPassword(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "firstName"
  }, "First Name"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: "form-control reginput",
    id: "firstName",
    placeholder: "John",
    onChange: e => setFirstName(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "lastName"
  }, "Last Name"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: "form-control reginput",
    id: "lastName",
    placeholder: "Doe",
    onChange: e => setLastName(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "dob"
  }, "Date of Birth"), /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "date",
    className: "form-control reginput",
    id: "dob",
    placeholder: "16/01/1998",
    onChange: e => setDob(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "imageURL"
  }, "ImageURL"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control reginput",
    id: "imageURL",
    placeholder: "https://...",
    onChange: e => setImageURL(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "username"
  }, "Username"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control reginput",
    id: "username",
    placeholder: "Johnny",
    onChange: e => setUsername(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    className: "form-check-input",
    type: "radio",
    id: "public-status",
    value: true,
    name: "status",
    checked: isPublic === true,
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
    value: false,
    name: "status",
    checked: isPublic === false,
    onChange: handleChange
  }), /*#__PURE__*/React.createElement("label", {
    className: "form-check-label",
    htmlFor: "private-status"
  }, "Private")), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "about"
  }, "About me"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control reginput",
    id: "bio",
    placeholder: "About Me",
    cols: "30",
    rows: "10",
    onChange: e => setBio(e.target.value)
  })), /*#__PURE__*/React.createElement("button", {
    className: "btn btn-primary",
    type: "submit"
  }, "Register")), /*#__PURE__*/React.createElement("div", {
    className: "error-message"
  }), /*#__PURE__*/React.createElement("br", null), " ", /*#__PURE__*/React.createElement("div", {
    className: "mb3"
  }, /*#__PURE__*/React.createElement("span", {
    className: "login-text"
  }, "Already have an account? \xA0"), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary",
    onClick: renderLogin
  }, "Log in")));
}
const renderProfile = (userId, isEditable) => {
  const appContainer = document.querySelector(".app-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Profile, {
    userId: userId,
    isEditable: isEditable
  }), appContainer);
};
function Profile({
  userId,
  isEditable
}) {
  const [profileUserData, setProfileUserData] = useState({});
  const [userPostData, setUserPostData] = useState([]);
  const [userFollowerData, setUserFollowerData] = useState([]);
  const [userFollowsData, setUserFollowsData] = useState([]);
  const [isPublicValue, setIsPublicValue] = useState(null);
  const fetchProfileData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/profile/${userId}`, {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json"
        }
      });
      if (!response.ok) {
        throw new Error("Failed to fetch profile data");
      }
      const data = await response.json();
      setProfileUserData(data.profileUserData);
      setUserPostData(data.userPostData || []);
      setUserFollowerData(data.userFollowerData || []);
      setUserFollowsData(data.userFollowsData || []);
      setIsPublicValue(data.profileUserData.isPublic);
    } catch (error) {
      console.error("Error fetching profile data:", error);
    }
  };
  useEffect(() => {
    fetchProfileData();
  }, []);

  // useEffect(() => {
  // 	// This effect will re-render the component whenever isPublicValue changes
  // }, [isPublicValue]);

  const handlePrivacyChange = event => {
    const newPrivacySetting = JSON.parse(event.target.value);

    // Update the database with the new privacy status
    fetch("http://localhost:8080/api/profile/privacy", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        userId: profileUserData.userId,
        isPublic: newPrivacySetting
      })
    }).then(response => {
      if (!response.ok) {
        throw new Error("Failed to update privacy status");
      }

      // Update the local state with the new privacy setting
      setIsPublicValue(newPrivacySetting);
    }).catch(error => {
      console.error("Error updating privacy status:", error);
    });
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("div", {
    id: "profileData"
  }, /*#__PURE__*/React.createElement("h2", null, profileUserData.username, "'s Profile"), /*#__PURE__*/React.createElement("div", {
    id: "myProfileData"
  }), isEditable ? /*#__PURE__*/React.createElement("div", {
    id: "isPublicToggle"
  }, /*#__PURE__*/React.createElement("label", null, /*#__PURE__*/React.createElement("input", {
    type: "radio",
    value: true,
    checked: isPublicValue === true // Check if isPublicValue is true
    ,
    onChange: handlePrivacyChange
  }), "Public"), /*#__PURE__*/React.createElement("label", null, /*#__PURE__*/React.createElement("input", {
    type: "radio",
    value: false,
    checked: isPublicValue === false // Check if isPublicValue is false
    ,
    onChange: handlePrivacyChange
  }), "Private")) : /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Privacy:"), " ", isPublicValue ? "Public" : "Private"), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "User ID:"), " ", profileUserData.userId), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Username:"), " ", profileUserData.username), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Email:"), " ", profileUserData.email), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "First Name:"), " ", profileUserData.firstName), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Last Name:"), " ", profileUserData.lastName), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Date of Birth:"), " ", new Date(profileUserData.dob).toLocaleDateString()), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Bio:"), " ", profileUserData.bio), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Image URL:"), " ", profileUserData.imageURL), /*#__PURE__*/React.createElement("h2", null, profileUserData.username, "'s Posts"), /*#__PURE__*/React.createElement("div", {
    id: "myPostsData"
  }, userPostData.map(post => /*#__PURE__*/React.createElement("div", {
    key: post.postId
  }, /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Post ID:"), " ", post.postId), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Created At:"), " ", post.createdAt), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Body:"), " ", post.body), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Image URL:"), " ", post.imageURL)))), /*#__PURE__*/React.createElement("h2", null, profileUserData.username, "'s Followers"), /*#__PURE__*/React.createElement("div", {
    id: "myFollowersData"
  }, userFollowerData && userFollowerData.map(follower => /*#__PURE__*/React.createElement("p", {
    key: follower.userId
  }, follower.userId))), /*#__PURE__*/React.createElement("h2", null, profileUserData.username, "'s Followed"), /*#__PURE__*/React.createElement("div", {
    id: "usersIFollowData"
  }, userFollowsData && userFollowsData.map(user => /*#__PURE__*/React.createElement("p", {
    key: user.userId
  }, user.userId)))));
}
function Chat() {
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("h1", null, "Chat"));
}
function Group() {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [error, setError] = useState(null);
  const handleSubmit = async e => {
    e.preventDefault();
    const groupData = new FormData();
    groupData.append('group-title', title);
    groupData.append('group-description', description);
    try {
      const response = await fetch('http://localhost:8080/api/groups', {
        method: 'POST',
        credentials: 'include',
        body: groupData
      });
      if (!response.ok) {
        throw new Error('Failed to create group');
      }

      // Reset form fields on successful submission
      setTitle('');
      setDescription('');
      setError(null);
    } catch (error) {
      setError('Failed to create group. Please try again.');
      console.error('Error creating group:', error);
    }
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("form", {
    onSubmit: handleSubmit
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "titleInput",
    className: "form-label"
  }, "Title"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "titleInput",
    value: title,
    onChange: e => setTitle(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "descriptionInput",
    className: "form-label"
  }, "Description"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "descriptionInput",
    value: description,
    onChange: e => setDescription(e.target.value)
  })), error && /*#__PURE__*/React.createElement("div", {
    className: "alert alert-danger"
  }, error), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "Create")));
}
function Notifications() {
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("h1", null, "Notifications"));
}

// PostForm component
// This component renders a form for creating a new post.
// It accepts a `groupId` prop to determine the group for the post.
function PostForm({
  groupId
}) {
  const [body, setBody] = useState("");
  const [privacy, setPrivacy] = useState("");
  const [selectedFile, setSelectedFile] = useState(null);

  // Handler for form submission
  const submit = async e => {
    e.preventDefault(); // Prevent page reload

    const formData = new FormData();

    // Append form data
    formData.append("body", body);
    formData.append("privacy", privacy);
    if (privacy === "private") {
      groupId = -1; // Set groupId to -1 for private posts
    }
    formData.append("groupId", groupId);
    if (selectedFile) {
      formData.append("image", selectedFile);
    }
    console.log("Form data being sent to backend: ", formData);
    try {
      // Send user data to the server
      await fetch("http://localhost:8080/api/posts", {
        method: "POST",
        credentials: "include",
        body: formData
      });

      // Reset form fields after successful submission
      setBody("");
      setPrivacy("");
      setSelectedFile(null);
      document.getElementById("postFormBody").value = "";
    } catch (error) {
      console.error("Error submitting post:", error);
    }
  };

  // Handler for file selection
  const handleFileChange = e => {
    setSelectedFile(e.target.files[0]);
  };
  const handleSelectFile = () => {
    const fileInput = document.getElementById("fileInput");
    fileInput.click();
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("main", {
    className: "postForm container",
    style: {
      maxWidth: "400px"
    }
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal"
  }, "Post Message Here"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-floating mb-3"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "postFormBody",
    placeholder: "Type your post here...",
    onChange: e => setBody(e.target.value)
  })), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-primary",
    onClick: handleSelectFile
  }, "Select File"), /*#__PURE__*/React.createElement("span", null, selectedFile ? selectedFile.name : "No file selected"), /*#__PURE__*/React.createElement("input", {
    type: "file",
    id: "fileInput",
    accept: "image/*",
    style: {
      display: "none"
    },
    onChange: handleFileChange
  })), /*#__PURE__*/React.createElement("br", null), " ", /*#__PURE__*/React.createElement("div", {
    className: "form-floating mb-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "radio",
    id: "post-public-status",
    value: "public",
    name: "status",
    checked: privacy === "public",
    onClick: e => setPrivacy(e.target.value),
    className: "form-check-input"
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "post-public-status",
    className: "form-check-label"
  }, "Public")), /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "radio",
    id: "private-status",
    value: "private",
    name: "status",
    checked: privacy === "private",
    onClick: e => setPrivacy(e.target.value),
    className: "form-check-input"
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "private-status",
    className: "form-check-label"
  }, "Private"))), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary",
    type: "submit"
  }, "Submit"))));
}
const postCardStyle = {
  maxWidth: '600px',
  background: 'linear-gradient(to bottom, #c7ddef, #ffffff)',
  // Light blue/grey to white gradient
  borderRadius: '10px',
  boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
  // Optional: Add shadow for depth
  padding: '20px',
  margin: 'auto',
  marginBottom: '20px' // Adjust spacing between post cards
};
function FollowButton({
  userId,
  isFollowed
}) {
  const [isFollowing, setIsFollowing] = useState(isFollowed);
  const handleFollowToggle = async () => {
    if (isFollowing) {
      // If already following, unfollow the user
      await handleUnfollow(userId);
    } else {
      // If not following, follow the user
      await handleFollow(userId);
    }
    // Toggle the local follow state
    setIsFollowing(!isFollowing);
  };
  const handleFollow = async subjectId => {
    try {
      const response = await fetch(`http://localhost:8080/api/user/userUser/${subjectId}`, {
        method: "POST",
        credentials: "include"
      });
      if (response.ok) {
        console.log("Successfully followed the user.");
        return true; // Return true if the follow request is successful
      } else {
        console.error("Failed to follow the user.");
      }
    } catch (error) {
      console.error("Error following the user:", error);
    }
    return false; // Return false if the follow request fails
  };
  const handleUnfollow = async subjectId => {
    try {
      const response = await fetch(`http://localhost:8080/api/user/userUser/${subjectId}`, {
        method: "DELETE",
        credentials: "include"
      });
      if (response.ok) {
        console.log("Successfully unfollowed the user.");
        return true; // Return true if the follow request is successful
      } else {
        console.error("Failed to unfollow the user.");
      }
    } catch (error) {
      console.error("Error following the user:", error);
    }
    return false; // Return false if the follow request fails
  };
  return /*#__PURE__*/React.createElement("button", {
    className: "btn btn-primary btn-sm",
    onClick: handleFollowToggle
  }, isFollowing ? 'Unfollow' : 'Follow');
}
function PostCard({
  post
}) {
  const [isFollowing, setIsFollowing] = useState(false);
  const [body, setBody] = useState("");
  const [selectedFile, setSelectedFile] = useState(null);
  const milliseconds = post.post.createdAt;
  const date = new Date(milliseconds);
  const formattedDate = date.toLocaleString();
  const handleFollowClick = async () => {
    const followSuccess = await handleFollow(post.post.userId);
    setIsFollowing(followSuccess);
  };
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    const formData = new FormData();

    // Append form data
    formData.append("body", body);
    formData.append("postId", post.post.postId);
    if (selectedFile) {
      formData.append("image", selectedFile);
    }
    console.log("Form data being sent to backend: ", formData);

    // Send user data to golang api/PostHandler.go.
    await fetch("http://localhost:8080/api/comments", {
      method: "POST",
      credentials: "include",
      body: formData
    });

    // Reset the form fields to their default state
    setBody("");
    setSelectedFile(null);
    document.getElementById("commentTextArea").value = "";
  };

  // Function to handle file selection
  const handleFileChange = e => {
    setSelectedFile(e.target.files[0]);
    // const file = e.target.files[0];
  };
  const handleSelectFile = () => {
    const commentFileInput = document.getElementById(`commentFileInput${post.post.postId}`);
    commentFileInput.click();
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "card",
    style: postCardStyle
  }, /*#__PURE__*/React.createElement("div", {
    className: "card-body"
  }, /*#__PURE__*/React.createElement("div", {
    className: "d-flex flex-start align-items-center"
  }, /*#__PURE__*/React.createElement("img", {
    className: "rounded-circle shadow-1-strong me-3",
    src: post.post.imageURL,
    alt: "avatar",
    width: "60",
    height: "60"
  }), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("div", {
    className: "d-flex align-items-center mb-1"
  }, /*#__PURE__*/React.createElement("a", {
    className: "fw-bold text-primary mb-0 me-2",
    href: "#",
    onClick: () => renderProfile(post.post.userId)
  }, post.post.userId), /*#__PURE__*/React.createElement("button", {
    className: "btn btn-primary btn-sm",
    onClick: handleFollowClick,
    disabled: isFollowing
  }, isFollowing ? "Following" : "Follow")), /*#__PURE__*/React.createElement("p", {
    className: "text-muted small mb-0"
  }, formattedDate))), !post.post.imageURL ? null : /*#__PURE__*/React.createElement("p", {
    className: "mt-3 mb-2 pb-1"
  }, /*#__PURE__*/React.createElement("img", {
    src: post.post.imageURL,
    className: "img-fluid"
  })), /*#__PURE__*/React.createElement("p", {
    className: "mt-3 mb-2 pb-1"
  }, post.post.body)), /*#__PURE__*/React.createElement("div", {
    className: "card-footer py-3 border-0",
    style: {
      backgroundColor: "#f8f9fa"
    }
  }, /*#__PURE__*/React.createElement("div", {
    className: "d-flex flex-start w-100"
  }, /*#__PURE__*/React.createElement("img", {
    className: "rounded-circle shadow-1-strong me-3",
    src: post.avatar,
    alt: "avatar",
    width: "40",
    height: "40"
  }), /*#__PURE__*/React.createElement("div", {
    className: "form-outline w-100"
  }, /*#__PURE__*/React.createElement("textarea", {
    className: "form-control",
    id: "commentTextArea",
    rows: "4",
    style: {
      background: "#fff"
    },
    onChange: e => setBody(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    className: "form-label",
    htmlFor: "textAreaExample"
  }, "Message"))), /*#__PURE__*/React.createElement("div", {
    className: "float-end mt-2 pt-1"
  }, /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-primary",
    onClick: handleSelectFile
  }, "Select File"), /*#__PURE__*/React.createElement("span", null, selectedFile ? selectedFile.name : "No file selected"), /*#__PURE__*/React.createElement("input", {
    type: "file",
    id: `commentFileInput${post.post.postId}`,
    accept: "image/*",
    style: {
      display: "none"
    },
    onChange: handleFileChange
  }), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary btn-sm",
    onClick: submit
  }, "Post comment")), /*#__PURE__*/React.createElement("div", {
    className: "comments"
  }, /*#__PURE__*/React.createElement("h2", null, "Comments"), post.comments !== null && post.comments.length > 0 ? post.comments.map(comment => /*#__PURE__*/React.createElement(CommentCard, {
    key: comment.createdAt,
    comment: comment
  })) : /*#__PURE__*/React.createElement("p", {
    className: "text-muted"
  }, "No comments"))));
}
function CommentCard({
  comment
}) {
  const formattedDate = new Date(comment.createdAt).toLocaleString();
  return /*#__PURE__*/React.createElement("div", {
    className: "card mt-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "d-flex flex-start align-items-center"
  }, /*#__PURE__*/React.createElement("img", {
    className: "rounded-circle shadow-1-strong me-3",
    src: comment.imageURL,
    alt: "avatar",
    width: "60",
    height: "60"
  }), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h6", {
    className: "fw-bold text-primary mb-1"
  }, comment.userId), /*#__PURE__*/React.createElement("p", {
    className: "text-muted small mb-0"
  }, formattedDate))), comment.imageURL && /*#__PURE__*/React.createElement("div", {
    className: "mt-3 mb-2 pb-1"
  }, /*#__PURE__*/React.createElement("img", {
    src: comment.imageURL,
    className: "img-fluid",
    alt: "comment"
  })), /*#__PURE__*/React.createElement("div", {
    className: "card-body"
  }, /*#__PURE__*/React.createElement("p", {
    className: "card-text"
  }, comment.body)));
}

// Display information relating to homepage
function Home() {
  const [userList, setUserList] = useState([]);
  const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
  const [privatePosts, setPrivatePosts] = useState([]);
  const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
  const [userGroups, setUserGroups] = useState([]);
  useEffect(() => {
    fetch("http://localhost:8080/api/home").then(response => response.json()).then(data => {
      setUserList(data.userList);
      setAlmostPrivatePosts(data.almostPrivatePosts);
      setPrivatePosts(data.privatePosts);
      setPublicPostsWithComments(data.publicPostsWithComments);
      setUserGroups(data.userGroups);
    }).catch(error => {
      console.error("Error fetching data:", error);
    });
  }, []);
  return /*#__PURE__*/React.createElement("main", {
    className: "homePage"
  }, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement(PostForm, {
    groupId: 0
  }), /*#__PURE__*/React.createElement("div", {
    className: "userList"
  }, /*#__PURE__*/React.createElement("h2", null, "UserList"), userList !== null && userList.length > 0 ? userList.map((user, index) => /*#__PURE__*/React.createElement("div", {
    key: index
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderProfile(user.userId)
  }, user.username), /*#__PURE__*/React.createElement(FollowButton, {
    userId: user.userId,
    isFollowed: user.isFollowed
  }))) : /*#__PURE__*/React.createElement("p", null, "No Users?!")), /*#__PURE__*/React.createElement("div", {
    className: "almostPrivatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Almost Private Posts"), almostPrivatePosts !== null && almostPrivatePosts.length > 0 ? almostPrivatePosts.map(almostPrivatePost => /*#__PURE__*/React.createElement(PostCard, {
    key: almostPrivatePost.createdAt,
    post: almostPrivatePost
  })) : /*#__PURE__*/React.createElement("p", null, "No almost private posts")), /*#__PURE__*/React.createElement("div", {
    className: "privatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Private Posts"), privatePosts !== null && privatePosts.length > 0 ? privatePosts.map(privatePost => /*#__PURE__*/React.createElement(PostCard, {
    key: privatePost.createdAt,
    post: privatePost
  })) : /*#__PURE__*/React.createElement("p", null, "No private posts")), /*#__PURE__*/React.createElement("div", {
    className: "publicPostsWithComments"
  }, /*#__PURE__*/React.createElement("h2", null, "Public Posts With Comments"), publicPostsWithComments !== null && publicPostsWithComments.length > 0 ? publicPostsWithComments.map((publicPostsWithComment, index) => /*#__PURE__*/React.createElement(PostCard, {
    key: index,
    post: publicPostsWithComment
  })) : /*#__PURE__*/React.createElement("p", null, "public posts")), /*#__PURE__*/React.createElement("div", {
    className: "userGroups"
  }, /*#__PURE__*/React.createElement("h2", null, "Groups"), /*#__PURE__*/React.createElement("ul", null, userGroups !== null && userGroups.map(userGroup => /*#__PURE__*/React.createElement("li", {
    key: userGroup.createdAt
  }, userGroup.Title, " ")))));
}
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);