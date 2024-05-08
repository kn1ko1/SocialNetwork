const {
  useState,
  useEffect
} = React;
import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { PostForm } from "./components/Home/PostForm.js";
import { PostCard } from "./components/shared/PostCard.js";
import { FollowButton } from "./components/shared/FollowButton.js";
import { renderProfile } from "./Profile.js";
import { Chat } from "./Chat.js";
export const renderHome = ({
  socket
}) => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Home, {
    socket: socket
  }), pageContainer);
};

// Display information relating to homepage
export function Home({
  socket
}) {
  const {
    currentUserId
  } = getCurrentUserId();
  const [userList, setUserList] = useState([]);
  const [followedUsers, setFollowedUsers] = useState([]);
  const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
  const [privatePosts, setPrivatePosts] = useState([]);
  const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
  const [userGroups, setUserGroups] = useState([]);
  const [userList2, setUserList2] = useState([]);
  const [followedUsersList, setFollowedUsersList] = useState([]);
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const promises = [];
        promises.push(fetch(`http://localhost:8080/api/users`));
        promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers`));
        const results = await Promise.all(promises);
        const userListResponse = results[0];
        const followedUserListResponse = results[1];
        if (!userListResponse.ok) {
          throw new Error('Failed to fetch user list');
        }
        if (!followedUserListResponse.ok) {
          throw new Error('Failed to fetch followed users list');
        }
        const userListData = await userListResponse.json();
        const followedUsersList = await followedUserListResponse.json();
        setUserList2(userListData);
        // setGroupMembers(groupMembersData);
        console.log("UserListData2", userListData);
        console.log("followedUserList", followedUsersList);
        const filteredFollowedUsers = userListData.filter(user => followedUsersList.some(followedUser => followedUser.subjectId === user.userId));
        setFollowedUsersList(filteredFollowedUsers);
        console.log("filteredFollowedUsers", filteredFollowedUsers);
      } catch (error) {
        console.error('Error fetching group data:', error);
      }
    };
    if (currentUserId !== null) {
      fetchUserData();
    }
  }, [currentUserId]);
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
  }, /*#__PURE__*/React.createElement(PostForm, {
    groupId: 0,
    followedUsers: followedUsersList
  }), /*#__PURE__*/React.createElement("div", {
    class: "container text-center"
  }, /*#__PURE__*/React.createElement("div", {
    class: "row align-items-start"
  }, /*#__PURE__*/React.createElement("div", {
    class: "col-3"
  }, /*#__PURE__*/React.createElement("div", {
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
  }))) : /*#__PURE__*/React.createElement("p", null, "No Users?!"))), /*#__PURE__*/React.createElement("div", {
    class: "col-6"
  }, /*#__PURE__*/React.createElement("div", {
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
  }, userGroup.title))))), /*#__PURE__*/React.createElement("div", {
    class: "col-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "card"
  }, /*#__PURE__*/React.createElement(Chat, {
    socket: {
      socket
    }
  }))))));
}