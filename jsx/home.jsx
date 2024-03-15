function Home() {
	const [users, setUsers] = useState([]);
	const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
	const [privatePosts, setPrivatePosts] = useState([]);
	const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
	const [userEvents, setUserEvents] = useState([]);
	const [userGroups, setUserGroups] = useState([]);
	const [userNotifications, setUserNotifications] = useState([]);

	useEffect(() => {
		fetch('http://localhost:8080/api/home')
			.then(response => response.json())
			.then(data => {
				setUsers(data.allUsers);
				setAlmostPrivatePosts(data.almostPrivatePosts)
				setPrivatePosts(data.privatePosts)
				setPublicPostsWithComments(data.publicPostsWithComments)
				setUserEvents(data.userEvents)
				setUserGroups(data.userGroups)
				setUserNotifications(data.userNotifications)
			})
			.catch(error => {
				console.error('Error fetching data:', error);
			});
	}, []);

	return (
		<div className="homePage">

			<div className="allUsersList">
				<h2>All Users</h2>
				<ul>
					{users.map(user => (
						<li key={user.userId}>
							{user.username} - {user.email} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="almostPrivatePosts">
				<h2>Almost Private Posts</h2>
				<ul>
					{almostPrivatePosts !== null && almostPrivatePosts.map(almostPrivatePost => (
						<li key={almostPrivatePost.createdAt}>
							{almostPrivatePost.body} - {almostPrivatePost.UserId}
							{/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="privatePosts">
				<h2>Private Posts</h2>
				<ul>
				{privatePosts !== null && privatePosts.map(privatePost => (
						<li key={privatePost.createdAt}>
							{privatePost.body} - {privatePost.UserId} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="publicPostsWithComments">
				<h2>Public Posts</h2>
				<ul>
				{publicPostsWithComments !== null && publicPostsWithComments.map(publicPostsWithComment => (
						<li key={publicPostsWithComment.post.CreatedAt}>
							{publicPostsWithComment.post.Body} - {publicPostsWithComment.post.UserId} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="userEvents">
				<h2>Events</h2>
				<ul>
				{userEvents !== null && userEvents.map(userEvent => (
						<li key={userEvent.createdAt}>
							{userEvent.Title} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="userGroups">
				<h2>Groups</h2>
				<ul>
				{userGroups !== null && userGroups.map(userGroup => (
						<li key={userGroup.createdAt}>
							{userGroup.Title} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="userNotifications">
				<h2>Notifications</h2>
				<ul>
				{userNotifications !== null && userNotifications.map(userNotification => (
						<li key={userNotification.createdAt}>
							{userNotification.NotificationType} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>
		</div>
	);
}

export default Home;
