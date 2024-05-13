import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js";
import { PostFormGroup } from "./components/GroupDetail/PostFormGroup.js";
import { EventForm } from "./components/GroupDetail/EventForm.js";
import { GroupDetailsUserList } from "./components/GroupDetail/GroupDetailsUserList.js";
import { PostCard } from "./components/shared/PostCard.js";
import { GroupDetailsEvents } from "./components/GroupDetail/GroupDetailsEvent.js";
import { fetchCommentsForPosts } from "./components/shared/FetchCommentsForPosts.js";

const { useState, useEffect } = React

export const renderGroupDetails = (group) => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<GroupDetails group={group} />, pageContainer)
}

export function GroupDetails({ group }) {

	const { currentUserId } = getCurrentUserId()
	const [userList, setUserList] = useState([]);
	const [groupMembers, setGroupMembers] = useState([]);
	const [groupPosts, setGroupPosts] = useState([]);
	const [groupMessages, setGroupMessages] = useState([]);
	const [groupEvents, setGroupEvents] = useState([]);


	if (group.isMember) {
		useEffect(() => {

			fetchGroupData(group.groupId);
		}, [group.groupId]);
	}

	const fetchGroupData = async (groupId) => {
		try {
			const promises = [];
			promises.push(fetch(`http://localhost:8080/api/users/transport`));
			promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/groupUsers`));
			promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/posts`));
			promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/messages`));
			promises.push(fetch(`http://localhost:8080/api/groups/${groupId}/events`));
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
			if (postsData != null) {
				const postsWithComments = await fetchCommentsForPosts(postsData)
				setGroupPosts(postsWithComments);
			} else {
				setGroupPosts(null)
			}
			setGroupPosts(postsWithComments);
			setGroupMessages(messagesData);
			setGroupEvents(eventsData);


		} catch (error) {
			console.error('Error fetching group data:', error);
		}
	};

	async function AddGroupUser(userId, groupId, notificationType) {
		const notificationtData = {
			notificationType: notificationType,
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

	return (
		<div className="group-details">
			<h2>{group.title}</h2>
			<p>{group.description}</p>

			{group.isMember ? (
				<div id="groupData">
					<PostFormGroup group={group} fetchFunc={() => fetchGroupData(group.groupId)} />

					<EventForm group={group} />
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
					{/* Render group posts here */}
					<div id="groupPosts">
						<h2>Posts</h2>
						{groupPosts !== null ? (
							groupPosts.map((post) => (
								<li key={post.createdAt}>
									<PostCard
										post={post}
										comments={post.comments}
										showCommentForm={true}
										fetchFunc={() => fetchGroupData(group.groupId)} />

								</li>
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
					<GroupDetailsEvents groupEvents={groupEvents} />

				</div>
			) : (
				<div>
					<div>You are not a member yet</div>
					<button onClick={() => AddGroupUser(group.creatorId, group.groupId, "groupRequest")}>Request to join group</button>
				</div>
			)}





		</div>
	)
}