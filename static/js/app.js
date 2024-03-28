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
function Navbar() {
  const renderHome = () => {
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), appContainer);
  };
  const renderProfile = () => {
    const appContainer = document.querySelector(".app-container");
    ReactDOM.render( /*#__PURE__*/React.createElement(Profile, null), appContainer);
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
    onClick: renderProfile
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
  const [isPublic, setIsPublic] = useState("public");
  const [isRegistered, setIsRegistered] = useState(false);

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
    value: "public",
    name: "status",
    checked: isPublic === "public",
    onChange: e => setIsPublic(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    className: "form-check-label",
    htmlFor: "public-status"
  }, "Public")), /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    className: "form-check-input",
    type: "radio",
    id: "private-status",
    value: "private",
    name: "status",
    checked: isPublic === "private",
    onChange: e => setIsPublic(e.target.value)
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
function Profile() {
  const [profileUserData, setProfileUserData] = useState({});
  const [userPostData, setUserPostData] = useState([]);
  const [userFollowerData, setUserFollowerData] = useState([]);
  const [userFollowsData, setUserFollowsData] = useState([]);
  const [privacySetting, setPrivacySetting] = useState('');
  useEffect(() => {
    fetch("http://localhost:8080/api/profile", {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json"
      }
    }).then(response => {
      if (!response.ok) {
        throw new Error("Failed to fetch profile data");
      }
      return response.json();
    }).then(data => {
      // Update the HTML elements with the profile data under each category
      // const myProfileDataElement = document.getElementById('myProfileData');
      // myProfileDataElement.innerHTML = JSON.stringify(data.profileUserData, null, 2);

      // Set privacy setting based on isPublic value (0 or 1)
      setPrivacySetting(data.profileUserData.isPublic === 1 ? 'public' : 'private');
      setProfileUserData(data.profileUserData);
      setUserPostData(data.userPostData);
      console.log("userPostData", data.userPostData);
      setUserFollowerData(data.userFollowerData);
      console.log("userFollowerData", data.userFollowerData);
      setUserFollowsData(data.userFollowsData);
      console.log("userFollowsData", data.userFollowsData);

      // const myPostsDataElement = document.getElementById('myPostsData');
      // myPostsDataElement.innerHTML = JSON.stringify(data.userPostData, null, 2);

      // const myFollowersDataElement = document.getElementById('myFollowersData');
      // myFollowersDataElement.innerHTML = JSON.stringify(data.userFollowerData, null, 2);

      // const usersIFollowDataElement = document.getElementById('usersIFollowData');
      // usersIFollowDataElement.innerHTML = JSON.stringify(data.userFollowsData, null, 2);
    }).catch(error => {
      console.error("Error fetching profile data:", error);
    });
  }, []);
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("div", {
    id: "profileData"
  }, /*#__PURE__*/React.createElement("h2", null, "My Profile"), /*#__PURE__*/React.createElement("div", {
    id: "myProfileData"
  }), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("label", null, /*#__PURE__*/React.createElement("input", {
    type: "radio",
    value: "public",
    checked: privacySetting === 'public'
    // onChange={handlePrivacyChange}
  }), "Public"), /*#__PURE__*/React.createElement("label", null, /*#__PURE__*/React.createElement("input", {
    type: "radio",
    value: "private",
    checked: privacySetting === 'private'
    // onChange={handlePrivacyChange}
  }), "Private")), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "User ID:"), " ", profileUserData.userId), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Username:"), " ", profileUserData.username), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Email:"), " ", profileUserData.email), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "First Name:"), " ", profileUserData.firstName), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Last Name:"), " ", profileUserData.lastName), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Date of Birth:"), " ", new Date(profileUserData.dob).toLocaleDateString()), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Bio:"), " ", profileUserData.bio), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Image URL:"), " ", profileUserData.imageURL), /*#__PURE__*/React.createElement("h2", null, "My Posts"), /*#__PURE__*/React.createElement("div", {
    id: "myPostsData"
  }, userPostData.map(post => /*#__PURE__*/React.createElement("div", {
    key: post.postId
  }, /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Post ID:"), " ", post.postId), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Created At:"), " ", post.createdAt), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Body:"), " ", post.body), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Image URL:"), " ", post.imageURL)))), /*#__PURE__*/React.createElement("h2", null, "My Followers"), /*#__PURE__*/React.createElement("div", {
    id: "myFollowersData"
  }, userFollowerData && userFollowerData.map(follower => /*#__PURE__*/React.createElement("p", {
    key: follower.userId
  }, follower.userId))), /*#__PURE__*/React.createElement("h2", null, "Users I Follow"), /*#__PURE__*/React.createElement("div", {
    id: "usersIFollowData"
  }, userFollowsData && userFollowsData.map(user => /*#__PURE__*/React.createElement("p", {
    key: user.userId
  }, user.userId)))));
}
function Chat() {
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("h1", null, "Chat"));
}
function Group() {
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("h1", null, "Group"));
}
function Notifications() {
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("h1", null, "Notifications"));
}

// Main post form, defaults to sending posts to public group (0)
function PostForm({
  groupId
}) {
  const [body, setBody] = useState("");
  const [privacy, setPrivacy] = useState("");
  const [selectedFile, setSelectedFile] = useState(null);

  // Upon submitting:
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    const formData = new FormData();

    // Append form data
    formData.append('body', body);
    formData.append('privacy', privacy);
    if (privacy === "private") {
      groupId = -1;
    }
    formData.append('groupId', groupId);
    if (selectedFile) {
      formData.append("image", selectedFile);
    }
    console.log("Form data being sent to backend: ", formData);

    // Send user data to golang api/PostHandler.go.
    await fetch("http://localhost:8080/api/posts", {
      method: "POST",
      credentials: "include",
      body: formData
    });

    // Reset the form fields to their default state
    setBody("");
    setPrivacy("");
    setSelectedFile(null);
    document.getElementById("postFormBody").value = "";
  };

  // Function to handle file selection
  const handleFileChange = e => {
    setSelectedFile(e.target.files[0]);
    // const file = e.target.files[0];
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
function PostCard({
  post
}) {
  const [body, setBody] = useState("");
  const [selectedFile, setSelectedFile] = useState(null);
  const milliseconds = post.post.createdAt;
  const date = new Date(milliseconds);
  const formattedDate = date.toLocaleString();
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
  }), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h6", {
    className: "fw-bold text-primary mb-1"
  }, post.post.userId), /*#__PURE__*/React.createElement("p", {
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
  const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
  const [privatePosts, setPrivatePosts] = useState([]);
  const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
  const [userGroups, setUserGroups] = useState([]);
  useEffect(() => {
    fetch("http://localhost:8080/api/home").then(response => response.json()).then(data => {
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
    className: "almostPrivatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Almost Private Posts"), almostPrivatePosts !== null && almostPrivatePosts.length > 0 ? almostPrivatePosts.map(almostPrivatePost => /*#__PURE__*/React.createElement(PostCard, {
    key: almostPrivatePost.createdAt,
    post: almostPrivatePost
  })) : /*#__PURE__*/React.createElement("p", null, "No almost private posts")), /*#__PURE__*/React.createElement("div", {
    className: "privatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Almost Private Posts"), privatePosts !== null && privatePosts.length > 0 ? privatePosts.map(privatePost => /*#__PURE__*/React.createElement(PostCard, {
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