import { PostFormGroup } from "./components/PostFormGroup.js";
import { EventForm } from "./components/EventForm.js";

const { useState, useEffect } = React

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
				for (let i = 0; i < eventsData.length; i++) {
					const milliseconds = eventsData[i].dateTime;
					const date = new Date(milliseconds);
					const formattedDate = date.toLocaleDateString();
					eventsData[i].dateTime = formattedDate;
				}
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


	// Function to add a new group user
	async function AddGroupUser({ groupId, userId }) {
		const requestData = {
			groupId: groupId,
			userId: userId
		};

		console.log('Request data:', requestData);

		try {
			const response = await fetch('http://localhost:8080/api/groupUsers', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(requestData)
			});

			console.log('Response:', response);


			if (response.ok) {
				// Handle success response
				console.log('Group user added successfully!');
			} else {
				// Handle error response
				console.error('Failed to add group user:', response.statusText);
			}
		} catch (error) {
			console.error('Error adding group user:', error);
		}
	}

	// const UserList = ({ userList }) => {
	const handleAddToGroup = (userId) => {
		console.log('Adding user to group with groupId:', group.groupId);
		console.log('User ID:', userId);
		AddGroupUser({ groupId: group.groupId, userId: userId }); // Call AddGroupUser function with groupId and userId
	};

	return (
		<div className="group-details">
			<h2>{group.title}</h2>
			<p>{group.description}</p>
			{/* <p>Members: {group.members}</p> */}
			<PostFormGroup group={group} />

			<EventForm group={group} />
			{/* Render userList here */}
			<div className="userList">
				<h2>UserList</h2>
				{userList !== null && userList.length > 0 ? (
					userList
						.filter(user => !groupMembers.some(member => member.userId === user.userId))
						.map((user, index) => (
							<div key={index}>
								<span>{user.username}</span>
								<button onClick={() => handleAddToGroup(user.userId)}>Add to Group</button>
							</div>
						))
				) : (
					<p>No Users?!</p>
				)}
			</div>
			{/* Render group members here */}
			<div className="groupMembers">
				<h2>Group Members</h2>
				{groupMembers !== null && groupMembers.length > 0 ? (
					groupMembers.map((member, index) => {
						// Find the user object corresponding to the member's userId
						const user = userList.find((user) => user.userId === member.userId);
						return (
							<div key={index}>
								{user ? user.username : 'Unknown User'}
							</div>
						);
					})
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
						<div key={index} className="row">
							<div className="col-3">
								<p className="d-inline-flex gap-1">
									<button
										className="btn btn-primary"
										type="button"
										data-bs-toggle="collapse"
										// Use unique ID for data-bs-target
										data-bs-target={`#collapseExample${index}`}
										aria-expanded="false"
										aria-controls={`collapseExample${index}`}
									>
										{event.title}
									</button>
								</p>
							</div>
							<div className="col-9">
								<div className="collapse" id={`collapseExample${index}`}>
									<div className="card card-body">
										{event.description} - {event.dateTime}
									</div>
								</div>
							</div>
						</div>
					))
				) : (
					<p>No Events</p>
				)}
			</div>

		</div>
	)
}