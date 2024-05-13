const {
  useState,
  useEffect
} = React;
// PostForm component
// This component renders a form for creating a new post.
// It accepts a `groupId` prop to determine the group for the post
// and followedUsers so that you can assign users to Almost Private Posts
export function PostForm({
  groupId,
  followedUsers,
  fetchUserPostData
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
    // Really we should do something clever with websockets and updating useStates, but this is much easier
    fetchUserPostData();
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