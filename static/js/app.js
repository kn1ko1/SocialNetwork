const {
  useState,
  useEffect
} = React;
const App = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement(Login, null));
};
function Login(props) {
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
  return /*#__PURE__*/React.createElement("div", {
    className: "login-container"
  }, /*#__PURE__*/React.createElement("main", {
    className: "form-signin w-100 m-auto",
    style: {
      display: "block"
    }
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal login-text"
  }, "log in"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "floatingInput",
    className: "login-text"
  }, "Email address"), /*#__PURE__*/React.createElement("input", {
    type: "email",
    className: "form-control login-text",
    id: "floatingInput",
    placeholder: "name@example.com",
    onChange: e => setUsernameOrEmail(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "floatingPassword",
    className: "login-text"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    type: "password",
    className: "form-control login-text",
    id: "floatingPassword",
    placeholder: "Password",
    onChange: e => setPassword(e.target.value)
  })), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary login-button",
    type: "submit"
  }, "Sign in")), /*#__PURE__*/React.createElement("div", {
    className: "error-message"
  }), /*#__PURE__*/React.createElement("br", null), " ", /*#__PURE__*/React.createElement("span", {
    className: "login-text"
  }, "Don't have an account? \xA0"), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary login-button",
    onClick: renderRegister
  }, "Register")));
}
function Register(props) {
  const [email, setEmail] = useState("");
  const [encryptedPassword, setEncryptedPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [dob, setDob] = useState("");
  const [imageURL, setImageURL] = useState("");
  const [username, setUsername] = useState("");
  const [bio, setBio] = useState("");
  const [isPublic, setIsPublic] = useState("");
  const [redirectVar, setRedirectVar] = useState(false);
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
        throw new Error('Invalid credentials');
      }

      //takes response from backend and processes
      const data = await response.json();
      if (data.success) {
        setIsRegistered(true);
      } else {
        throw new Error('Invalid credentials');
      }
    } catch (error) {
      setError('Invalid credentials');
    }
  };

  //if credentials frontend succesfully create a new user then we render home
  if (isRegistered) {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), appContainer);
  }

  //this is the login button, when pressed will serve login form
  const renderLogin = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Login, null), appContainer);
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "login-container"
  }, " ", /*#__PURE__*/React.createElement("main", {
    className: "form-signin w-100 m-auto",
    style: {
      display: "block"
    }
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal"
  }, "register"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "email",
    className: "form-control",
    id: "floatingInput",
    placeholder: "name@example.com",
    onChange: e => setEmail(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "floatingInput"
  }, "Email address")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "password",
    className: "form-control reginput",
    id: "regpassword",
    placeholder: "Password",
    onChange: e => setEncryptedPassword(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "regpassword"
  }, "Password")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: "form-control reginput",
    id: "firstName",
    placeholder: "John",
    onChange: e => setFirstName(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "firstName"
  }, "First Name")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "text",
    className: "form-control reginput",
    id: "lastName",
    placeholder: "Doe",
    onChange: e => setLastName(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "lastName"
  }, "Last Name")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "date",
    className: "form-control reginput",
    id: "dob",
    placeholder: "16/01/1998",
    onChange: e => setDob(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "dob"
  }, "Date of Birth")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control reginput",
    id: "imageURL",
    placeholder: "https://...",
    onChange: e => setImageURL(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "imageURL"
  }, "ImageURL")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control reginput",
    id: "username",
    placeholder: "Johnny",
    onChange: e => setUsername(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "username"
  }, "Username")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-control reginput status"
  }, /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "radio",
    id: "public-status",
    value: "public",
    name: "status",
    checked: true,
    onClick: e => setIsPublic(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "public-status"
  }, "Public")), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "radio",
    id: "private-status",
    value: "private",
    name: "status",
    onClick: e => setIsPublic(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "private-status"
  }, "Private"))), /*#__PURE__*/React.createElement("label", {
    htmlFor: ""
  }, "Status")), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    className: "form-control reginput",
    name: "bio",
    placeholder: "About Me",
    id: "bio",
    cols: "30",
    rows: "10",
    onChange: e => setBio(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "about"
  }, "About me")), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary login-button",
    type: "submit"
  }, "Register")), /*#__PURE__*/React.createElement("div", {
    className: "error-message"
  }), /*#__PURE__*/React.createElement("br", null), " ", /*#__PURE__*/React.createElement("span", {
    className: "login-text"
  }, "Already have an account? \xA0"), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary login-button",
    onClick: renderLogin
  }, "Login")));
}

// Main post form, defaults to sending posts to public group (0)
function PostForm() {
  const [body, setBody] = useState("");
  const [imageURL, setImageURL] = useState("");
  const privacy = "public";
  let groupId = null;

  // Needs to be changed to get info from... cookie?
  const userId = Number(36);
  if (privacy === "public") {
    groupId = Number(0);
  }

  // Upon submitting:
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    // Reads info from returned HTML
    const postToSend = {
      body,
      groupId,
      imageURL,
      privacy,
      userId
    };
    console.log("Post being sent to backend: ", postToSend);

    // Send user data to golang api/PostHandler.go.
    await fetch("http://localhost:8080/api/posts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      credentials: "include",
      body: JSON.stringify(postToSend)
    });
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("main", {
    className: "postForm",
    style: {
      display: "block"
    }
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal"
  }, "Post Message Here"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "postFormBody",
    placeholder: "name@example.com",
    onChange: e => setBody(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "form-floating"
  }, /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "postFormImgURL",
    placeholder: "This/Will/Need/A/Button.gif",
    onChange: e => setImageURL(e.target.value)
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "postFormImgURL"
  }, "Image URL")), /*#__PURE__*/React.createElement("button", {
    className: "w-100 btn btn-lg btn-primary",
    type: "submit"
  }, "Submit"))));
}

// Display information relating to homepage
function Home() {
  const [users, setUsers] = useState([]);
  const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
  const [privatePosts, setPrivatePosts] = useState([]);
  const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
  const [userEvents, setUserEvents] = useState([]);
  const [userGroups, setUserGroups] = useState([]);
  const [userNotifications, setUserNotifications] = useState([]);
  useEffect(() => {
    fetch('http://localhost:8080/api/home').then(response => response.json()).then(data => {
      setUsers(data.allUsers);
      setAlmostPrivatePosts(data.almostPrivatePosts);
      setPrivatePosts(data.privatePosts);
      setPublicPostsWithComments(data.publicPostsWithComments);
      setUserEvents(data.userEvents);
      setUserGroups(data.userGroups);
      setUserNotifications(data.userNotifications);
    }).catch(error => {
      console.error('Error fetching data:', error);
    });
  }, []);
  return /*#__PURE__*/React.createElement("div", {
    className: "homePage"
  }, /*#__PURE__*/React.createElement("div", {
    className: "allUsersList"
  }, /*#__PURE__*/React.createElement("h2", null, "All Users"), /*#__PURE__*/React.createElement("ul", null, users.map(user => /*#__PURE__*/React.createElement("li", {
    key: user.userId
  }, user.username, " - ", user.email, " ")))), /*#__PURE__*/React.createElement("div", {
    className: "almostPrivatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Almost Private Posts"), /*#__PURE__*/React.createElement("ul", null, almostPrivatePosts !== null && almostPrivatePosts.map(almostPrivatePost => /*#__PURE__*/React.createElement("li", {
    key: almostPrivatePost.createdAt
  }, almostPrivatePost.body, " - ", almostPrivatePost.UserId)))), /*#__PURE__*/React.createElement("div", {
    className: "privatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Private Posts"), /*#__PURE__*/React.createElement("ul", null, privatePosts !== null && privatePosts.map(privatePost => /*#__PURE__*/React.createElement("li", {
    key: privatePost.createdAt
  }, privatePost.body, " - ", privatePost.UserId, " ")))), /*#__PURE__*/React.createElement("div", {
    className: "publicPostsWithComments"
  }, /*#__PURE__*/React.createElement("h2", null, "Public Posts"), /*#__PURE__*/React.createElement("ul", null, publicPostsWithComments !== null && publicPostsWithComments.map(publicPostsWithComment => /*#__PURE__*/React.createElement("li", {
    key: publicPostsWithComment.post.CreatedAt
  }, publicPostsWithComment.post.Body, " - ", publicPostsWithComment.post.UserId, " ")))), /*#__PURE__*/React.createElement("div", {
    className: "userEvents"
  }, /*#__PURE__*/React.createElement("h2", null, "Events"), /*#__PURE__*/React.createElement("ul", null, userEvents !== null && userEvents.map(userEvent => /*#__PURE__*/React.createElement("li", {
    key: userEvent.createdAt
  }, userEvent.Title, " ")))), /*#__PURE__*/React.createElement("div", {
    className: "userGroups"
  }, /*#__PURE__*/React.createElement("h2", null, "Groups"), /*#__PURE__*/React.createElement("ul", null, userGroups !== null && userGroups.map(userGroup => /*#__PURE__*/React.createElement("li", {
    key: userGroup.createdAt
  }, userGroup.Title, " ")))), /*#__PURE__*/React.createElement("div", {
    className: "userNotifications"
  }, /*#__PURE__*/React.createElement("h2", null, "Notifications"), /*#__PURE__*/React.createElement("ul", null, userNotifications !== null && userNotifications.map(userNotification => /*#__PURE__*/React.createElement("li", {
    key: userNotification.createdAt
  }, userNotification.NotificationType, " ")))));
}

// Elements related to the homepage clustered together for the login function to return.
// Needs work, but it's a start
const completedHomePage = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "completedHomePage"
  }, /*#__PURE__*/React.createElement(PostForm, null), /*#__PURE__*/React.createElement(Home, null));
};
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);