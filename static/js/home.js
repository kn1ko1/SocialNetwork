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
export default Home;