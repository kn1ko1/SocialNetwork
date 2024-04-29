import { useSocket } from "./shared/UserProvider.js";
import { PostFormGroup } from "./components/PostFormGroup.js";
import { EventForm } from "./components/EventForm.js";
import { GroupDetailsUserList } from "./components/GroupDetailsUserList.js";
import { PostCard } from "./components/PostCard.js";
import { GroupDetailsEvents } from "./components/GroupDetailsEvent.js";
import { Chat } from "./Chat.js";

const { useState, useEffect } = React

export function GroupDetails({ group }) {

	const {currentUserId} = useSocket()
	const [userList, setUserList] = useState([]);
	const [groupMembers, setGroupMembers] = useState([]);
	const [groupPosts, setGroupPosts] = useState([]);
	const [groupMessages, setGroupMessages] = useState([]);
	const [groupEvents, setGroupEvents] = useState([]);


	if (group.isMember) {
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
					if (eventsData != null) {
						for (let i = 0; i < eventsData.length; i++) {
							const milliseconds = eventsData[i].dateTime;
							const date = new Date(milliseconds);
							const formattedDate = date.toLocaleDateString();
							eventsData[i].dateTime = formattedDate;
						}
					}

					setUserList(userListData);
					setGroupMembers(groupMembersData);
					setGroupPosts(postsData);
					setGroupMessages(messagesData);
					setGroupEvents(eventsData);


				} catch (error) {
					console.error('Error fetching group data:', error);
				}
			};

			fetchGroupData();
		}, [group.groupId]);
	} else {

	}

	// Function to add a new group user
	async function AddGroupUser(userId, groupId) {
		const notificationtData = {
			notificationType: "groupInvite",
			objectId: groupId,
			senderId: currentUserId,
			status: "pending",
			targetId: userId,
		};

		console.log('notificationtData:', notificationtData);

		try {
			const response = await fetch('http://localhost:8080/api/notifications', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(notificationtData)
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

	const handleAddToGroup = (userId) => {
		console.log('Adding user to group with groupId:', group.groupId);
		console.log('User ID:', userId);
		AddGroupUser({ groupId: group.groupId, userId: userId }); // Call AddGroupUser function with groupId and userId
	};

	return (
		<div className="group-details">
			<h2>{group.title}</h2>
			<p>{group.description}</p>

			{group.isMember ? (
				<div id="groupData">
					<PostFormGroup group={group} />

					<EventForm group={group} />
					<div class="container text-center">
						<div class="row align-items-start">
							<div class="col-3">
								{/* Render user List here */}
								<GroupDetailsUserList
									userList={userList}
									groupId={group.groupId}
									groupMembers={groupMembers}
									AddGroupUser={AddGroupUser} />

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
								<GroupDetailsEvents groupEvents={groupEvents} />

							</div>
							<div class="col-6">
								{/* Render group posts here */}
								<div id="groupPosts">
									<h2>Posts</h2>
									{groupPosts !== null ? (
										groupPosts.map((post) => (
											<div key={post.createdAt}>
												<PostCard post={post} />
											</div>
										))
									) : (
										<div id="groupPosts">There are no posts in this groups yet</div>
									)}
								</div>
							</div>
							<div class="col-3">
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
								<Chat socket={socket} messageType={"group"} targetId={group.groupId}/>
							</div>
						</div>
					</div>





				</div>
			) : (
				<div>
					<div>You are not a member yet</div>
					<button onClick={() => handleAddToGroup(currentUserId)}>Request to join group</button>
				</div>
			)}





		</div>
	)
}