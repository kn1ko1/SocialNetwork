const {
  useState,
  useEffect
} = React;
const App = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement(Login, null), /*#__PURE__*/React.createElement(Register, null), /*#__PURE__*/React.createElement(Home, null));
};
function Login(props) {
  const [usernameOrEmail, setUsernameOrEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirectVar, setRedirectVar] = useState(false);
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    const userToLogin = {
      usernameOrEmail,
      password
    };
    console.log(userToLogin);

    // Send user data to golang register function.
    const response = await fetch("http://localhost:8080/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      credentials: "include",
      body: JSON.stringify(userToLogin)
    });
    const validUser = await response.json();
    setRedirectVar(true);
    props.setName(validUser.first);
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
  }, "Sign in")), /*#__PURE__*/React.createElement("span", null, "Already have an account? \xA0")));
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

  // Redirect
  //const navigate = useNavigate();

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
    // Send user data to golang register function.
    const response = await fetch("http://localhost:8080/auth/registration", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(newUser)
    });
    console.log("dob", newUser.dob);
    await response.json();
    // let result = await response.json()
    // if (result.email === email) {
    setRedirectVar(true);
    // }
  };

  // if (redirectVar) {
  // 	return navigate("/login"); // This is still iffy!!! ????????????
  // }

  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("main", {
    className: "form-signin w-100 m-auto",
    style: {
      display: "block"
    }
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal"
  }, "Please register"), /*#__PURE__*/React.createElement("form", {
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
    className: "w-100 btn btn-lg btn-primary",
    type: "submit"
  }, "Register")), /*#__PURE__*/React.createElement("span", null, "Already have an account? \xA0")));
}
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
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);