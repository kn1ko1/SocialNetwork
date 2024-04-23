import { Chat } from "./Chat.js";
import { Profile } from "./Profile.js";
import { Register } from "./Register.js";
import { FollowButton } from "./components/FollowButton.js";
import { GroupDetails } from "./GroupDetails.js";
import { getCurrentUserId } from "./shared/getCurrentUserId.js";
import { PostForm } from "./components/PostForm.js";
import { PostCard } from "./components/PostCard.js";
const {
  useState,
  useEffect
} = React;
let socket;
const App = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement("div", {
    className: "nav-container"
  }), /*#__PURE__*/React.createElement("div", {
    className: "page-container"
  }, /*#__PURE__*/React.createElement(Login, null)));
};
const renderNavbar = () => {
  const navContainer = document.querySelector(".nav-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Navbar, null), navContainer);
};
function Navbar() {
  const {
    currentUserId
  } = getCurrentUserId();
  const logout = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/logout", {
        method: "POST",
        credentials: "include"
      });
      console.log(response);
      if (response.ok) {
        socket.close();
        socket.addEventListener("close", event => {
          console.log("The connection has been closed successfully.");
        });
        renderLogin();
        const navContainer = document.querySelector(".nav-container");
        ReactDOM.render(null, navContainer);
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
    onClick: () => renderProfile(currentUserId, true)
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
const renderLogin = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Login, null), pageContainer);
};
function Login() {
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
      renderNavbar();
      renderHome();
      socket = new WebSocket("ws://localhost:8080/ws");
      socket.onopen = function (event) {
        console.log("WebSocket connection established.");
      };
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
const renderRegister = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Register, null), pageContainer);
};
const renderProfile = (userId, isEditable) => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Profile, {
    userId: userId,
    isEditable: isEditable
  }), pageContainer);
};
const renderChat = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Chat, {
    socket: socket
  }), pageContainer);
};
const renderGroup = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Group, null), pageContainer);
};
function Group() {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [groupData, setGroupData] = useState([]);
  const [selectedGroup, setSelectedGroup] = useState(null);
  //const [showGroupDetails, setShowGroupDetails] = useState(false);

  const fetchGroupData = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/groups", {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json"
        }
      });
      if (!response.ok) {
        throw new Error("Failed to fetch group data");
      }
      const data = await response.json();
      setGroupData(data);
    } catch (error) {
      console.error("Error fetching group data:", error);
    }
  };
  useEffect(() => {
    fetchGroupData();
  }, []);
  const create = async e => {
    e.preventDefault(); // prevent reload.

    const groupData = new FormData();

    // Append form data
    groupData.append("group-title", title);
    groupData.append("group-description", description);
    console.log("Group data being sent to backend:", title);
    console.log("Group data being sent to backend:", description);

    // Send user data to golang api/PostHandler.go.
    await fetch("http://localhost:8080/api/groups", {
      method: "POST",
      credentials: "include",
      body: groupData
    });
    setTitle("");
    setDescription("");
    document.getElementById("exampleTitle").value = "";
    document.getElementById("exampleDescription").value = "";
    fetchGroupData();
  };
  const handleGroupClick = group => {
    setSelectedGroup(group);
    //setShowGroupDetails(true);
  };
  const handleGoBack = () => {
    setSelectedGroup(null);
    setShowGroupDetails(false); // Update showGroupDetails to false when going back
  };
  return /*#__PURE__*/React.createElement("div", null, selectedGroup ? /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("button", {
    onClick: () => setSelectedGroup(null)
  }, "Go Back"), /*#__PURE__*/React.createElement(GroupDetails, {
    group: selectedGroup
  })) : /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("form", {
    onSubmit: create,
    className: "container",
    style: {
      maxWidth: "400px"
    }
  }, /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleTitle",
    className: "form-label"
  }, "Title"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "exampleTitle",
    "aria-describedby": "emailHelp",
    value: title,
    onChange: e => setTitle(e.target.value)
  })), /*#__PURE__*/React.createElement("div", {
    className: "mb-3"
  }, /*#__PURE__*/React.createElement("label", {
    htmlFor: "exampleInputPassword1",
    className: "form-label"
  }, "Description"), /*#__PURE__*/React.createElement("input", {
    type: "text",
    className: "form-control",
    id: "exampleDescription",
    value: description,
    onChange: e => setDescription(e.target.value)
  })), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "Create")), /*#__PURE__*/React.createElement("div", {
    id: "groupData"
  }, groupData !== null ? groupData.map(group => /*#__PURE__*/React.createElement("div", {
    key: group.title,
    onClick: () => handleGroupClick(group)
  }, /*#__PURE__*/React.createElement("h3", null, group.title), /*#__PURE__*/React.createElement("p", null, group.description))) : /*#__PURE__*/React.createElement("div", {
    id: "noGroupsError"
  }, "There are no created groups yet"))));
}
const renderNotifications = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Notifications, null), pageContainer);
};
function Notifications() {
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Notifications"));
}
const renderHome = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), pageContainer);
};

// Display information relating to homepage
function Home() {
  const {
    currentUserId
  } = getCurrentUserId();
  const [userList, setUserList] = useState([]);
  const [followedUsers, setFollowedUsers] = useState([]);
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
  useEffect(() => {
    // Filter userList to get only the followed users
    const filteredFollowedUsers = userList.filter(user => user.isFollowed === true);

    // Set the filtered list to followedUsers state
    setFollowedUsers(filteredFollowedUsers);
  }, [userList]);
  return /*#__PURE__*/React.createElement("main", {
    className: "homePage"
  }, /*#__PURE__*/React.createElement(PostForm, {
    groupId: 0,
    followedUsers: followedUsers
  }), /*#__PURE__*/React.createElement("div", {
    className: "userList"
  }, /*#__PURE__*/React.createElement("h2", null, "UserList"), userList !== null && userList.length > 0 ? userList
  // Filter out the current user
  .filter(user => user.userId !== currentUserId).map((user, index) => /*#__PURE__*/React.createElement("div", {
    key: index
  }, /*#__PURE__*/React.createElement("a", {
    className: "nav-link",
    href: "#",
    onClick: () => renderProfile(user.userId)
  }, user.username), /*#__PURE__*/React.createElement(FollowButton, {
    followerId: currentUserId,
    subjectId: user.userId,
    isFollowed: user.isFollowed
  }))) : /*#__PURE__*/React.createElement("p", null, "No Users?!")), /*#__PURE__*/React.createElement("div", {
    className: "almostPrivatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Almost Private Posts"), almostPrivatePosts !== null && almostPrivatePosts.length > 0 ? almostPrivatePosts.map(almostPrivatePost => /*#__PURE__*/React.createElement(PostCard, {
    key: almostPrivatePost.createdAt,
    post: almostPrivatePost.post,
    comments: almostPrivatePost.comments,
    showCommentForm: true
  })) : /*#__PURE__*/React.createElement("p", null, "No almost private posts")), /*#__PURE__*/React.createElement("div", {
    className: "privatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Private Posts"), privatePosts !== null && privatePosts.length > 0 ? privatePosts.map(privatePost => /*#__PURE__*/React.createElement(PostCard, {
    key: privatePost.createdAt,
    post: privatePost.post,
    comments: privatePost.comments,
    showCommentForm: true
  })) : /*#__PURE__*/React.createElement("p", null, "No private posts")), /*#__PURE__*/React.createElement("div", {
    className: "publicPostsWithComments"
  }, /*#__PURE__*/React.createElement("h2", null, "Public Posts With Comments"), publicPostsWithComments !== null && publicPostsWithComments.length > 0 ? publicPostsWithComments.map((publicPostsWithComment, index) => /*#__PURE__*/React.createElement(PostCard, {
    key: index,
    post: publicPostsWithComment.post,
    comments: publicPostsWithComment.comments,
    showCommentForm: true
  })) : /*#__PURE__*/React.createElement("p", null, "public posts")), /*#__PURE__*/React.createElement("div", {
    className: "userGroups"
  }, /*#__PURE__*/React.createElement("h2", null, "Groups"), /*#__PURE__*/React.createElement("ul", null, userGroups !== null && userGroups.map(userGroup => /*#__PURE__*/React.createElement("li", {
    key: userGroup.createdAt
  }, userGroup.Title, " ")))));
}
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);