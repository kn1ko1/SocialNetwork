const {
  useState
} = React;
const App = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement(Login, null), /*#__PURE__*/React.createElement(Register, null), /*#__PURE__*/React.createElement(Home, null), /*#__PURE__*/React.createElement(Profile, null));
};
function Login(props) {
  const [usernameOrEmail, setUsernameOrEmail] = useState("");
  const [password, setPassword] = useState("");
  // const [redirectVar, setRedirectVar] = useState(false);
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
    // setRedirectVar(true);
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
function Home(props) {
  return /*#__PURE__*/React.createElement("main", null, /*#__PURE__*/React.createElement("div", {
    className: "contentContainer"
  }, props.name ? /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement(ProfileImgContainer, {
    name: props.name,
    user: props.user,
    imageURL: props.imageURL
  }), /*#__PURE__*/React.createElement(GroupContainer, {
    groups: props.groups,
    socket: props.socket
  }), /*#__PURE__*/React.createElement(PostForm, {
    imageURL: props.imageURL
  }), /*#__PURE__*/React.createElement(RightSide, {
    openConnection: props.openConnection,
    fetchRequestData: props.fetchRequestData
  }), /*#__PURE__*/React.createElement(GetChat, null)) : /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement("p", null, "You are not logged in"))));
}
function Profile(props) {
  const [status, setStatus] = useState("");
  const [privatePosts, setPrivatePosts] = useState([]);

  // Update status to props.user.status.
  // useEffect(() => {
  //   setStatus(props.user.status);
  // }, [props.user.status]);

  const sendStatusToBackend = async data => {
    console.log(data);
    await fetch("http://localhost:8080/update-user-status", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      credentials: "include",
      body: JSON.stringify(data)
    });
  };
  const updateUserStatus = async ev => {
    let buttonClicked = ev.target.getAttribute("data-type");
    if (buttonClicked === "private") {
      sendStatusToBackend({
        user: props.user.email,
        setStatus: "private"
      });
      setStatus("private");
    } else if (buttonClicked === "public") {
      // update on backend if user is not already public
      sendStatusToBackend({
        user: props.user.email,
        setStatus: "public"
      });
      setStatus("public");
    }
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "profileContainer"
  }, "name=", props.name, "user=", props.user, "imageURL=", props.imageURL, "socket=", props.socket, "currentUser=", props.currentUser, "fetchUsersData=", props.fetchUsersData, "update=", props.update, "setUpdate=", props.setUpdate, /*#__PURE__*/React.createElement("div", {
    className: "formContainer"
  }, /*#__PURE__*/React.createElement("div", {
    className: "smallAvatar"
  }, /*#__PURE__*/React.createElement("img", {
    src: props.imageURL,
    alt: "profile photo"
  })), /*#__PURE__*/React.createElement("div", {
    className: "profile-page-title"
  }, props.name, "'s Posts")), props.currentUser === undefined ? /*#__PURE__*/React.createElement("div", {
    id: "set-public-private",
    className: "privacyButtons",
    style: {
      width: "100%",
      backgroundColor: "white",
      justifyContent: "space-evenly",
      alignItems: "center"
    }
  }, /*#__PURE__*/React.createElement(React.Fragment, null, /*#__PURE__*/React.createElement("button", {
    className: "postType",
    onClick: updateUserStatus,
    "data-type": "private",
    disabled: status === "private" ? true : false,
    style: {
      backgroundColor: status === "private" ? "rgba(129, 25, 41, 0.55)" : "rgb(148, 28, 47)"
    }
  }, "Set Private"), /*#__PURE__*/React.createElement("button", {
    className: "postType",
    onClick: updateUserStatus,
    "data-type": "public",
    disabled: status === "public" ? true : false,
    style: {
      backgroundColor: status === "public" ? "rgba(129, 25, 41, 0.55)" : "rgb(148, 28, 47)"
    }
  }, "Set Public"))) : /*#__PURE__*/React.createElement("div", {
    id: "set-public-private",
    className: "privacyButtons",
    style: {
      width: "100%",
      backgroundColor: "rgba(250, 250, 250, 0.5)"
    }
  }));
}
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);