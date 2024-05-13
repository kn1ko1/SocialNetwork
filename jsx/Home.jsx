const { useState, useEffect } = React
import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js"
import { formattedDate } from "./components/shared/FormattedDate.js"
import { PostForm } from "./components/Home/PostForm.js"
import { PostCard } from "./components/shared/PostCard.js"
import { FollowButton } from "./components/shared/FollowButton.js"
import { renderProfile } from "./Profile.js"

export const renderHome = ({ socket }) => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Home socket={socket} />, pageContainer)
}

// Display information relating to homepage
export function Home({ socket }) {
	const { currentUserId } = getCurrentUserId()
	const [userList, setUserList] = useState([])
	const [followedUsers, setFollowedUsers] = useState([]);
	const [almostPrivatePosts, setAlmostPrivatePosts] = useState([])
	const [privatePosts, setPrivatePosts] = useState([])
	const [publicPostsWithComments, setPublicPostsWithComments] = useState([])
	const [userGroups, setUserGroups] = useState([])
	const [userList2, setUserList2] = useState([])
	const [followedUsersList, setFollowedUsersList] = useState([])
	const [userEvents, setUserEvents] = useState([])




	useEffect(() => {
		const fetchUserData = async () => {
			try {
				const promises = [];
				promises.push(fetch(`http://localhost:8080/api/users`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers`));
				promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/events`));
				// promises.push(fetch(`http://localhost:8080/api/users/${currentUserId}/posts`));

				const results = await Promise.all(promises);

				const userListResponse = results[0];
				const followedUserListResponse = results[1];
				const userEventsResponse = results[2];
				// const allPostsResponse = results[3];

				if (!userListResponse.ok) {
					throw new Error('Failed to fetch user list');
				}
				if (!followedUserListResponse.ok) {
					throw new Error('Failed to fetch followed users list');
				}
				if (!userEventsResponse.ok) {
					throw new Error('Failed to fetch users events');
				}
				// if (!allPostsResponse.ok) {
				// 	throw new Error('Failed to fetch all posts available to user');
				// }

				const userListData = await userListResponse.json();
				const followedUsersListData = await followedUserListResponse.json();
				const userEventsData = await userEventsResponse.json();


				let filteredFollowedUsers = null;
				let updatedUserListData = userListData;

				if (followedUsersListData != null) {
					filteredFollowedUsers = userListData.filter(user =>
						followedUsersListData.some(followedUser => followedUser.subjectId === user.userId)
					);

					// Add isFollowed property to each user where userId matches subjectId
					updatedUserListData = userListData.map(user => ({
						...user,
						isFollowed: followedUsersListData.some(followedUser => followedUser.subjectId === user.userId)
					}));
				} else {
					// If followedUsersList is null, set isFollowed to false for every user
					updatedUserListData = userListData.map(user => ({
						...user,
						isFollowed: false
					}));
				}

				setUserList2(updatedUserListData);
				console.log("updatedUserListData", updatedUserListData)
				setFollowedUsersList(filteredFollowedUsers);
				console.log("filteredFollowedUsers", filteredFollowedUsers)
				setUserEvents(userEventsData);
				console.log("userEventsData", userEventsData)
			} catch (error) {
				console.error('Error fetching group data:', error);
			}
		};
		if (currentUserId != null) {
			fetchUserData()
		}

	}, [currentUserId])

	useEffect(() => {
		fetch("http://localhost:8080/api/home")
			.then((response) => response.json())
			.then((data) => {
				setUserList(data.userList)
				setAlmostPrivatePosts(data.almostPrivatePosts)
				setPrivatePosts(data.privatePosts)
				setPublicPostsWithComments(data.publicPostsWithComments)
				setUserGroups(data.userGroups)
			})
			.catch((error) => {
				console.error("Error fetching data:", error)
			})
	}, [])

	return (
		<main className="homePage">
			<PostForm groupId={0} followedUsers={followedUsersList} />

			<div class="container text-center">
				<div class="row align-items-start">
					<div class="col-3">
						<div className="userList">
							<h2>UserList</h2>
							{userList2 !== null && userList2.length > 0 ? (
								userList2
									// Filter out the current user
									.filter(user => user.userId !== currentUserId)
									.map((user, index) => (
										<div key={index}>
											<a
												className="nav-link"
												href="#"
												onClick={() => renderProfile(user.userId)}
											>
												{user.username}
											</a>
											<FollowButton
												followerId={currentUserId}
												user={user}
											/>
										</div>
									))
							) : (
								<p>No Users?!</p>
							)}
						</div>
					</div>
					<div class="col-6">

						{/* Rendering Public Posts */}
						<div className="publicPostsWithComments">
							<h2>Public Posts</h2>
							{publicPostsWithComments !== null &&
								publicPostsWithComments.length > 0 ? (
								publicPostsWithComments.map((publicPostsWithComment, index) => (
									<PostCard
										key={index}
										post={publicPostsWithComment.post}
										comments={publicPostsWithComment.comments}
										showCommentForm={true} />
								))
							) : (
								<p>public posts</p>
							)}
						</div>

						{/* Rendering Almost Private Posts */}
						<div className="almostPrivatePosts">
							<h2>Almost Private Posts</h2>
							{almostPrivatePosts !== null && almostPrivatePosts.length > 0 ? (
								almostPrivatePosts.map((almostPrivatePost) => (
									<PostCard
										key={almostPrivatePost.createdAt}
										post={almostPrivatePost.post}
										comments={almostPrivatePost.comments}
										showCommentForm={true}
									/>
								))
							) : (
								<p>No almost private posts</p>
							)}
						</div>

						{/* Rendering Private Posts */}
						<div className="privatePosts">
							<h2>Private Posts</h2>
							{privatePosts !== null && privatePosts.length > 0 ? (
								privatePosts.map((privatePost) => (
									<PostCard
										key={privatePost.createdAt}
										post={privatePost.post}
										comments={privatePost.comments}
										showCommentForm={true} />
								))
							) : (
								<p>No private posts</p>
							)}
						</div>

						{/* Rendering User Groups */}
						<div className="userGroups">
							<h2>Groups</h2>
							<ul>
								{userGroups !== null &&
									userGroups.map((userGroup) => (
										<li key={userGroup.createdAt}>
											{userGroup.title}

										</li>
									))}
							</ul>
						</div>
					</div>
					<div class="col-3">

						<div className="userEvents">
							<h2>Events that you're attending</h2>
							{userEvents !== null && userEvents.length > 0 ? (
								userEvents.map((event) => (
									<li key={event.dateTime}>
										{event.title} - {event.description} 
										- {formattedDate(event.dateTime)} 
									</li>
								))
							) : (
								<p>No almost private posts</p>
							)}
						</div>
					</div>
				</div>
			</div>
		</main>
	)
}
