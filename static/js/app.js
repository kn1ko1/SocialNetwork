const {
  useState,
  useEffect
} = React;
const App = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement(Login, null));
};

// function Dummy(props) {
// 	return (
// 		<div className="container-fluid text-center">
// 		<div className="row mb-2">
// 			<div className="col-lg-3">
// 				<h1>Welcome</h1>
// 			</div>
// 			<div className="col-lg-6">
// 				<h1>Welcome</h1>
// 			</div>
// 			<div className="col-lg-3">
// 				<h1>Welcome</h1>
// 			</div>
// 		</div>
// 		<div className="row">
// 			<div className="col-6">
// 				<h1>Welcome</h1>
// 			</div>
// 			<div className="col-6">
// 				<h1>Welcome</h1>
// 			</div>
// 		</div>
// 	</div>
// 	)
// }

function Navbar() {
  const renderHome = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), appContainer);
  };
  const renderProfile = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Profile, null), appContainer);
  };
  const renderNotifications = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Notifications, null), appContainer);
  };
  const renderChat = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Chat, null), appContainer);
  };
  const renderGroup = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Group, null), appContainer);
  };

  // const logout = async () => {

  // 	const response = await fetch("http://localhost:8080/auth/logout", {
  // 		method: "POST",
  // 		credentials: "include",
  // 	});
  // 	const appContainer = document.querySelector('.app-container');
  // 	ReactDOM.render(<Login />, appContainer);

  // 	const cookieHeader = response.headers.get('set-cookie');
  // 	if (cookieHeader) {
  // 		document.cookie = cookieHeader;
  // 		console.log("Logout successful!");
  // 	} else {
  // 		console.log("Failed to logout");
  // 	}
  // };

  const logout = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/logout", {
        method: "POST",
        credentials: "include"
      });
      const cookieHeader = response.headers.get('Set-Cookie');
      console.log("this is cookieHeader;", cookieHeader);
      if (response.ok) {
        const appContainer = document.querySelector('.app-container');
        ReactDOM.render( /*#__PURE__*/React.createElement(Login, null), appContainer);

        // if (cookieHeader) {
        // 	// document.cookie = cookieHeader;
        //       console.log("HERE?")
        // 	  // Check if the protocol is not HTTPS before setting the Secure attribute
        // 	  const secureAttribute = window.location.protocol !== 'https:' ? '' : '; Secure';
        // 	    // Add SameSite=None attribute to the cookie
        // 		const sameSiteAttribute = '; SameSite=None';
        // 	  document.cookie = `${cookieHeader}${secureAttribute}${sameSiteAttribute}`;

        console.log("Logout successful!");
        // } else {
        // 	console.log("Failed to set cookie after logout");
        //}
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
  const [usernameOrEmail, setUsernameOrEmail] = useState("");
  const [password, setPassword] = useState("");
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
    //ReactDOM.render(<Home />, appContainer);
    ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), appContainer);
  }

  //this is the register button, when pressed will serve registration form
  const renderRegister = () => {
    const appContainer = document.querySelector('.app-container');
    ReactDOM.render( /*#__PURE__*/React.createElement(Register, null), appContainer);
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "container login-container"
  }, /*#__PURE__*/React.createElement("h1", {
    className: "h3 mb-3 fw-normal login-text"
  }, "log in"), /*#__PURE__*/React.createElement("form", {
    onSubmit: submit
  }, /*#__PURE__*/React.createElement("div", {
    class: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    for: "exampleInputEmail1",
    class: "form-label"
  }, "Email address"), /*#__PURE__*/React.createElement("input", {
    type: "email",
    className: "form-control form-control-lg",
    id: "exampleInputEmail1",
    "aria-describedby": "emailHelp",
    onChange: e => setUsernameOrEmail(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    class: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    for: "exampleInputPassword1",
    class: "form-label"
  }, "Password"), /*#__PURE__*/React.createElement("input", {
    type: "password",
    className: "form-control form-control-lg",
    id: "exampleInputPassword1",
    onChange: e => setPassword(e.target.value)
  })), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    class: "btn btn-primary"
  }, "Log in")), /*#__PURE__*/React.createElement("div", {
    className: "error-message"
  }), /*#__PURE__*/React.createElement("br", null), " ", /*#__PURE__*/React.createElement("div", {
    className: "mb3"
  }, /*#__PURE__*/React.createElement("span", {
    className: "login-text"
  }, "Don't have an account? \xA0"), /*#__PURE__*/React.createElement("button", {
    type: "submit",
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
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement(Navbar, null), /*#__PURE__*/React.createElement("h1", null, "Profile"));
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

  // Needs to be changed to get info from... cookie?
  // Upon submitting:
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    const formData = new FormData();

    // Append form data
    formData.append('body', body);
    formData.append('privacy', privacy);
    formData.append('groupId', groupId);
    if (selectedFile) {
      formData.append('image', selectedFile);
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
    document.getElementById('postFormBody').value = "";
  };

  // Function to handle file selection
  const handleFileChange = e => {
    setSelectedFile(e.target.files[0]);
    // const file = e.target.files[0];
  };
  const handleSelectFile = () => {
    const fileInput = document.getElementById('fileInput');
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
  }, "Select File"), /*#__PURE__*/React.createElement("span", null, selectedFile ? selectedFile.name : 'No file selected'), /*#__PURE__*/React.createElement("input", {
    type: "file",
    id: "fileInput",
    accept: "image/*",
    style: {
      display: 'none'
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
function PostCard({
  post
}) {
  console.log("post", post);
  console.log("post.post.userId", post.post.userId);
  return /*#__PURE__*/React.createElement("div", {
    className: "card"
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
  }, post.post.createdAt))), /*#__PURE__*/React.createElement("p", {
    className: "mt-3 mb-4 pb-2"
  }, post.post.body), /*#__PURE__*/React.createElement("div", {
    className: "small d-flex justify-content-start"
  }, /*#__PURE__*/React.createElement("a", {
    href: "#!",
    className: "d-flex align-items-center me-3"
  }, /*#__PURE__*/React.createElement("i", {
    className: "far fa-thumbs-up me-2"
  }), /*#__PURE__*/React.createElement("p", {
    className: "mb-0"
  }, "Like")), /*#__PURE__*/React.createElement("a", {
    href: "#!",
    className: "d-flex align-items-center me-3"
  }, /*#__PURE__*/React.createElement("i", {
    className: "far fa-comment-dots me-2"
  }), /*#__PURE__*/React.createElement("p", {
    className: "mb-0"
  }, "Comment")), /*#__PURE__*/React.createElement("a", {
    href: "#!",
    className: "d-flex align-items-center me-3"
  }, /*#__PURE__*/React.createElement("i", {
    className: "fas fa-share me-2"
  }), /*#__PURE__*/React.createElement("p", {
    className: "mb-0"
  }, "Share")))), /*#__PURE__*/React.createElement("div", {
    className: "card-footer py-3 border-0",
    style: {
      backgroundColor: '#f8f9fa'
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
    id: "textAreaExample",
    rows: "4",
    style: {
      background: '#fff'
    }
  }), /*#__PURE__*/React.createElement("label", {
    className: "form-label",
    htmlFor: "textAreaExample"
  }, "Message"))), /*#__PURE__*/React.createElement("div", {
    className: "float-end mt-2 pt-1"
  }, /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-primary btn-sm"
  }, "Post comment"), /*#__PURE__*/React.createElement("button", {
    type: "button",
    className: "btn btn-outline-primary btn-sm"
  }, "Cancel"))));
}

// Display information relating to homepage
function Home() {
  const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
  const [privatePosts, setPrivatePosts] = useState([]);
  const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
  const [userGroups, setUserGroups] = useState([]);
  useEffect(() => {
    fetch('http://localhost:8080/api/home').then(response => response.json()).then(data => {
      setAlmostPrivatePosts(data.almostPrivatePosts);
      setPrivatePosts(data.privatePosts);
      setPublicPostsWithComments(data.publicPostsWithComments);
      setUserGroups(data.userGroups);
    }).catch(error => {
      console.error('Error fetching data:', error);
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
  }, /*#__PURE__*/React.createElement("h2", null, "Public Posts With Comments"), publicPostsWithComments !== null && publicPostsWithComments.length > 0 ? publicPostsWithComments.map(publicPostsWithComments => /*#__PURE__*/React.createElement(PostCard, {
    key: publicPostsWithComments.createdAt,
    post: publicPostsWithComments
  })) : /*#__PURE__*/React.createElement("p", null, "public posts")), /*#__PURE__*/React.createElement("div", {
    className: "userGroups"
  }, /*#__PURE__*/React.createElement("h2", null, "Groups"), /*#__PURE__*/React.createElement("ul", null, userGroups !== null && userGroups.map(userGroup => /*#__PURE__*/React.createElement("li", {
    key: userGroup.createdAt
  }, userGroup.Title, " ")))));
}
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);