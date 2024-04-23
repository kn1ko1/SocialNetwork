import { Chat } from "./Chat.js";
import { Profile } from "./Profile.js";
import { Register } from "./Register.js";
import { FollowButton } from "./components/FollowButton.js";
import { GroupDetails } from "./GroupDetails.js";
import { getCurrentUserId } from "./shared/getCurrentUserId.js";
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
    currentUserId,
    isLoading,
    error
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

// PostForm component
// This component renders a form for creating a new post.
// It accepts a `groupId` prop to determine the group for the post.
function PostForm({
  groupId,
  followedUsers
}) {
  const [body, setBody] = useState("");
  const [privacy, setPrivacy] = useState("");
  const [selectedFile, setSelectedFile] = useState(null);
  const [selectedUserIds, setSelectedUserIds] = useState([]);
  const [showFollowedUsersList, setShowFollowedUsersList] = useState(false);
  const [followedUsersForAP, setFollowedUsersForAP] = useState(followedUsers || []);
  useEffect(() => {
    setFollowedUsersForAP(followedUsers);
  }, [followedUsers]);
  const handleCheckboxChange = e => {
    const userId = e.target.value;
    const isChecked = e.target.checked;
    if (isChecked) {
      setSelectedUserIds(prevSelectedUserIds => [...prevSelectedUserIds, userId]);
    } else {
      setSelectedUserIds(prevSelectedUserIds => prevSelectedUserIds.filter(id => id !== userId));
    }
  };

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
    if (privacy === "almost private") {
      groupId = -2; // Set groupId to -2 for almost private posts
      formData.append("almostPrivatePostUsers", JSON.stringify(selectedUserIds));
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
      setPrivacy("public");
      setSelectedFile(null);
      setSelectedUserIds([]);
      document.getElementById("postFormBody").value = "";
      setShowFollowedUsersList(false);
    } catch (error) {
      console.error("Error submitting post:", error);
    }
  };
  const handlePrivacyChange = e => {
    const newValue = e.target.value;
    setPrivacy(newValue);
    if (newValue === 'almost private') {
      setShowFollowedUsersList(true);
    } else {
      setShowFollowedUsersList(false);
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
  const followedUsersList = showFollowedUsersList ? followedUsersForAP !== null && followedUsersForAP.length > 0 ? /*#__PURE__*/React.createElement("ul", null, followedUsersForAP.map(followedUser => /*#__PURE__*/React.createElement("li", {
    key: followedUser.username
  }, /*#__PURE__*/React.createElement("label", null, /*#__PURE__*/React.createElement("input", {
    type: "checkbox",
    value: followedUser.userId,
    onChange: handleCheckboxChange
  }), followedUser.username)))) : /*#__PURE__*/React.createElement("p", {
    className: "text-muted"
  }, "No followed users") : null;
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
    onClick: handlePrivacyChange,
    className: "form-check-input"
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "post-public-status",
    className: "form-check-label"
  }, "Public")), /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "radio",
    id: "post-private-status",
    value: "private",
    name: "status",
    checked: privacy === "private",
    onClick: handlePrivacyChange,
    className: "form-check-input"
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "private-status",
    className: "form-check-label"
  }, "Private")), /*#__PURE__*/React.createElement("div", {
    className: "form-check"
  }, /*#__PURE__*/React.createElement("input", {
    required: true,
    type: "radio",
    id: "post-almostPrivate-status",
    value: "almost private",
    name: "status",
    checked: privacy === "almost private",
    onClick: handlePrivacyChange,
    className: "form-check-input"
  }), /*#__PURE__*/React.createElement("label", {
    htmlFor: "private-status",
    className: "form-check-label"
  }, "Almost Private"))), followedUsersList, /*#__PURE__*/React.createElement("button", {
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
  post,
  comments
}) {
  const [body, setBody] = useState("");
  const [selectedFile, setSelectedFile] = useState(null);
  const milliseconds = post.createdAt;
  const date = new Date(milliseconds);
  const formattedDate = date.toLocaleString();
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    const formData = new FormData();

    // Append form data
    formData.append("body", body);
    formData.append("postId", post.postId);
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
    const commentFileInput = document.getElementById(`commentFileInput${post.postId}`);
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
    src: post.imageURL,
    alt: "avatar",
    width: "60",
    height: "60"
  }), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("div", {
    className: "d-flex align-items-center mb-1"
  }, /*#__PURE__*/React.createElement("a", {
    className: "fw-bold text-primary mb-0 me-2",
    href: "#",
    onClick: () => renderProfile(post.userId)
  }, post.userId)), /*#__PURE__*/React.createElement("p", {
    className: "text-muted small mb-0"
  }, formattedDate))), !post.imageURL ? null : /*#__PURE__*/React.createElement("p", {
    className: "mt-3 mb-2 pb-1"
  }, /*#__PURE__*/React.createElement("img", {
    src: post.imageURL,
    className: "img-fluid"
  })), /*#__PURE__*/React.createElement("p", {
    className: "mt-3 mb-2 pb-1"
  }, post.body)), /*#__PURE__*/React.createElement("div", {
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
    id: `commentFileInput${post.postId}`,
    accept: "image/*",
    style: {
      display: "none"
    },
    onChange: handleFileChange
  }), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary btn-sm",
    onClick: submit
  }, "Post comment")), comments && comments.length > 0 && /*#__PURE__*/React.createElement("div", {
    className: "comments"
  }, /*#__PURE__*/React.createElement("h2", null, "Comments"), comments.map(comment => /*#__PURE__*/React.createElement(CommentCard, {
    key: comment.createdAt,
    comment: comment
  })))));
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
    className: "fw-bold text-primary mb-1",
    onClick: () => renderProfile(comment.userId)
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
const renderHome = () => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Home, null), pageContainer);
};

// Display information relating to homepage
function Home() {
  const {
    currentUserId,
    isLoading,
    error
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
  }, /*#__PURE__*/React.createElement("h2", null, "UserList"), userList !== null && userList.length > 0 ? userList.map((user, index) => /*#__PURE__*/React.createElement("div", {
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
    comments: almostPrivatePost.comments
  })) : /*#__PURE__*/React.createElement("p", null, "No almost private posts")), /*#__PURE__*/React.createElement("div", {
    className: "privatePosts"
  }, /*#__PURE__*/React.createElement("h2", null, "Private Posts"), privatePosts !== null && privatePosts.length > 0 ? privatePosts.map(privatePost => /*#__PURE__*/React.createElement(PostCard, {
    key: privatePost.createdAt,
    post: privatePost.post,
    comments: privatePost.comments
  })) : /*#__PURE__*/React.createElement("p", null, "No private posts")), /*#__PURE__*/React.createElement("div", {
    className: "publicPostsWithComments"
  }, /*#__PURE__*/React.createElement("h2", null, "Public Posts With Comments"), publicPostsWithComments !== null && publicPostsWithComments.length > 0 ? publicPostsWithComments.map((publicPostsWithComment, index) => /*#__PURE__*/React.createElement(PostCard, {
    key: index,
    post: publicPostsWithComment.post,
    comments: publicPostsWithComment.comments
  })) : /*#__PURE__*/React.createElement("p", null, "public posts")), /*#__PURE__*/React.createElement("div", {
    className: "userGroups"
  }, /*#__PURE__*/React.createElement("h2", null, "Groups"), /*#__PURE__*/React.createElement("ul", null, userGroups !== null && userGroups.map(userGroup => /*#__PURE__*/React.createElement("li", {
    key: userGroup.createdAt
  }, userGroup.Title, " ")))));
}
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);