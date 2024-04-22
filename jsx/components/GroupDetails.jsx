export function GroupDetails({ group }) {

	const [userList, setUserList] = useState([]);
	const [groupMembers, setGroupMembers] = useState([]);
	const [groupPosts, setGroupPosts] = useState([]);
	const [groupMessages, setGroupMessages] = useState([]);
	const [groupEvents, setGroupEvents] = useState([]);


	useEffect(() => {
		const fetchGroupData = async () => {
			try {
				const promises = [];
				promises.push(fetch(`http://localhost:8080/api/users/transport`));
				promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/groupUsers`));
				promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/posts`));
				promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/messages`));
				promises.push(fetch(`http://localhost:8080/api/groups/${group.groupId}/events`));
				const results = await Promise.all(promises);

				const userListResponse = results[0]
				const groupMembersResponse = results[1]
				const postsResponse = results[2]
				const messagesResponse = results[3]
				const eventsResponse = results[4]
				if (!userListResponse.ok) {
					throw new Error('Failed to fetch user list');
				}
				if (!groupMembersResponse.ok) {
					throw new Error('Failed to fetch group members');
				}
				if (!postsResponse.ok) {
					throw new Error('Failed to fetch group posts');
				}
				if (!messagesResponse.ok) {
					throw new Error('Failed to fetch group messages');
				}
				if (!eventsResponse.ok) {
					throw new Error('Failed to fetch group eventsResponse');
				}
				const userListData = await userListResponse.json();
				const groupMembersData = await groupMembersResponse.json();
				const postsData = await postsResponse.json();
				const messagesData = await messagesResponse.json();
				const eventsData = await eventsResponse.json();
				setUserList(userListData);
				setGroupMembers(groupMembersData);
				setGroupPosts(postsData);
				setGroupMessages(messagesData);
				setGroupEvents(eventsData);
			} catch (error) {
				console.error('Error fetching group posts:', error);
			}
		};

		fetchGroupData();
	}, [group.groupId]);



	return (
		<div className="group-details">
			<h2>{group.title}</h2>
			<p>{group.description}</p>
			{/* <p>Members: {group.members}</p> */}
			<PostFormGroup groupId={group.groupId} />
			{/* Render userList here */}
			<div className="userList">
				<h2>UserList</h2>
				{userList !== null && userList.length > 0 ? (
					userList.map((user, index) => (
						<div key={index}>
							{user.username}
						</div>
					))
				) : (
					<p>No Users?!</p>
				)}
			</div>
			{/* Render group members here */}
			<div className="groupMembers">
				<h2>Group Members</h2>
				{groupMembers !== null && groupMembers.length > 1 ? (
					groupMembers.map((member, index) => (
						<div key={index}>
							{member.username}
						</div>
					))
				) : (
					<p>It's just you... Maybe you should invite someone?</p>
				)}
			</div>
			{/* Render group posts here */}
			<div id="groupPosts">
				<h2>Posts</h2>
				{groupPosts !== null ? (
					groupPosts.map((post) => (
						<li key={post.createdAt}>{post.body}</li>
					))
				) : (
					<div id="groupPosts">There are no posts in this groups yet</div>
				)}
			</div>
			{/* Render group Messages here */}
			<div className="groupMessages">
				<h2>Messages</h2>
				{groupMessages !== null && groupMessages.length > 0 ? (
					groupMessages.map((message, index) => (
						<div key={index}>
							{message.body}
						</div>
					))
				) : (
					<p>No Messages</p>
				)}
			</div>
			<div className="groupEvents">
				<h2>Events</h2>
				{groupEvents !== null && groupEvents.length > 0 ? (
					groupEvents.map((event, index) => (
						<div key={index}>
							{event.title}
						</div>
					))
				) : (
					<p>No Events</p>
				)}
			</div>
		</div>
	)
}