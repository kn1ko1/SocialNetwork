import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { PostCard } from "./components/shared/PostCard.js";
import { FollowButton } from "./components/shared/FollowButton.js";
const {
  useState,
  useEffect
} = React;
export const renderProfile = (userId, isEditable) => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Profile, {
    userId: userId,
    isEditable: isEditable
  }), pageContainer);
};
export function Profile({
  userId,
  isEditable
}) {
  const {
    currentUserId
  } = getCurrentUserId();
  const [profileUserData, setProfileUserData] = useState({});
  const [userPostData, setUserPostData] = useState([]);
  const [userFollowerData, setUserFollowerData] = useState([]);
  const [userFollowsData, setUserFollowsData] = useState([]);
  const [isPublicValue, setIsPublicValue] = useState(null);
  const [isFollowed, setIsFollowed] = useState(false);
  useEffect(() => {
    fetchProfileData();
  }, [userId]);
  useEffect(() => {
    if (!isEditable && currentUserId) {
      checkIfFollowed(currentUserId);
    }
  }, [profileUserData]);
  const fetchProfileData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/profile/${userId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json"
        }
      });
      if (!response.ok) {
        throw new Error(`Failed to fetch profile data: ${response.status} ${response.statusText}`);
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
  const checkIfFollowed = async currentUserId => {
    try {
      const response = await fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers/${userId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json"
        }
      });
      if (response.ok) {
        const updatedProfileUserData = {
          ...profileUserData,
          isFollowed: true
        };
        setProfileUserData(updatedProfileUserData);
        console.log("updatedProfileUserData", updatedProfileUserData);
        setIsFollowed(true);
        console.log("checkIfFollowed.  isFollowed", isFollowed);
        console.log("response", response);
      } else if (response.status === 404) {
        const updatedProfileUserData = {
          ...profileUserData,
          isFollowed: false
        };
        console.log("updatedProfileUserData", updatedProfileUserData);
        setProfileUserData(updatedProfileUserData);
        setIsFollowed(false);
        console.log("checkIfFollowed.  isFollowed", isFollowed);
      } else {
        console.error("Error fetching user user data:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching user user data:", error);
    }
  };
  const handlePrivacyChange = event => {
    const newPrivacySetting = JSON.parse(event.target.value);
    setIsPublicValue(newPrivacySetting);
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
    }).catch(error => {
      console.error("Error updating privacy status:", error);
      setIsPublicValue(!newPrivacySetting);
    });
  };
  const postCardStyle = {
    maxWidth: '1000px',
    background: 'linear-gradient(to bottom, #c7ddef, #ffffff)',
    // Light blue/grey to white gradient
    borderRadius: '10px',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
    // Optional: Add shadow for depth
    padding: '40px',
    margin: 'auto',
    marginBottom: '20px',
    // Adjust spacing between post cards
    border: '1px solid #ccc' // Add a thin border
  };
  return /*#__PURE__*/React.createElement("div", {
    className: "container",
    style: postCardStyle
  }, /*#__PURE__*/React.createElement("div", {
    className: "row"
  }, /*#__PURE__*/React.createElement("div", {
    className: "col-md-4"
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline',
      textAlign: 'center'
    }
  }, profileUserData.username, "'s Profile"), /*#__PURE__*/React.createElement("br", null), !isEditable && /*#__PURE__*/React.createElement("div", {
    className: "d-flex justify-content-center align-items-center"
  }, /*#__PURE__*/React.createElement(FollowButton, {
    followerId: currentUserId,
    user: profileUserData
  })), isPublicValue || isEditable || isFollowed ? /*#__PURE__*/React.createElement(React.Fragment, null, isEditable ? /*#__PURE__*/React.createElement("div", {
    id: "isPublicToggle"
  }, /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("h4", {
    style: {
      fontSize: '14px'
    }
  }, "Toggle to change profile privacy setting"), /*#__PURE__*/React.createElement("strong", null, "Privacy:"), /*#__PURE__*/React.createElement("label", null, /*#__PURE__*/React.createElement("input", {
    type: "radio",
    value: true,
    checked: isPublicValue === true,
    onChange: handlePrivacyChange
  }), "Public"), /*#__PURE__*/React.createElement("label", null, /*#__PURE__*/React.createElement("input", {
    type: "radio",
    value: false,
    checked: isPublicValue === false,
    onChange: handlePrivacyChange
  }), "Private"))) : /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Privacy:"), " ", isPublicValue ? "Public" : "Private"), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "User ID:"), " ", profileUserData.userId), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Username:"), " ", profileUserData.username), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Email:"), " ", profileUserData.email), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "First Name:"), " ", profileUserData.firstName), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Last Name:"), " ", profileUserData.lastName), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Date of Birth:"), " ", new Date(profileUserData.dob).toLocaleDateString()), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Bio:"), " ", profileUserData.bio), /*#__PURE__*/React.createElement("p", null, /*#__PURE__*/React.createElement("strong", null, "Image URL:"), " ", profileUserData.imageURL)) : /*#__PURE__*/React.createElement("p", null, "This profile is private.")), /*#__PURE__*/React.createElement("div", {
    className: "col-md-4"
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline',
      textAlign: 'center'
    }
  }, profileUserData.username, "'s Posts"), /*#__PURE__*/React.createElement("div", {
    id: "myPostsData"
  }, userPostData.map(post => /*#__PURE__*/React.createElement("div", {
    key: post.postId
  }, /*#__PURE__*/React.createElement(PostCard, {
    post: post,
    showCommentForm: false
  }))))), /*#__PURE__*/React.createElement("div", {
    className: "col-md-4"
  }, /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline',
      textAlign: 'center'
    }
  }, profileUserData.username, "'s Followers"), /*#__PURE__*/React.createElement("div", {
    id: "myFollowersData"
  }, userFollowerData && userFollowerData.map(follower => /*#__PURE__*/React.createElement("p", {
    key: follower.username
  }, follower.username))), /*#__PURE__*/React.createElement("h2", {
    style: {
      textDecoration: 'underline',
      textAlign: 'center'
    }
  }, "Users ", profileUserData.username, " Follows"), /*#__PURE__*/React.createElement("div", {
    id: "usersIFollowData"
  }, userFollowsData && userFollowsData.map(user => /*#__PURE__*/React.createElement("p", {
    key: user.username
  }, user.username))))));
}