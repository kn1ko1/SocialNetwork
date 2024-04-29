const { useState, useEffect } = React
import { getCurrentUserId } from "./shared/getCurrentUserId.js"
import { PostForm } from "./components/PostForm.js"
import { PostCard } from "./components/PostCard.js"
import { FollowButton } from "./components/FollowButton.js"
import { renderProfile } from "./Profile.js"
import { Chat } from "./Chat.js"

export const renderHome  = ({ socket }) => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Home socket={socket} />, pageContainer)
}

// Display information relating to homepage
export function Home({socket}) {
	const { currentUserId } = getCurrentUserId()
	const [userList, setUserList] = useState([])
	const [followedUsers, setFollowedUsers] = useState([]);
	const [almostPrivatePosts, setAlmostPrivatePosts] = useState([])
	const [privatePosts, setPrivatePosts] = useState([])
	const [publicPostsWithComments, setPublicPostsWithComments] = useState([])
	const [userGroups, setUserGroups] = useState([])

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

	useEffect(() => {
		// Filter userList to get only the followed users
		const filteredFollowedUsers = userList.filter(user => user.isFollowed === true);

		// Set the filtered list to followedUsers state
		setFollowedUsers(filteredFollowedUsers);
	}, [userList]);


	return (
		<main className="homePage">
			<PostForm groupId={0} followedUsers={followedUsers} />

			<div class="container text-center">
				<div class="row align-items-start">
					<div class="col-3">
						<div className="userList">
							<h2>UserList</h2>
							{userList !== null && userList.length > 0 ? (
								userList
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
												subjectId={user.userId}
												isFollowed={user.isFollowed}
											/>
										</div>
									))
							) : (
								<p>No Users?!</p>
							)}
						</div>
					</div>
					<div class="col-6">
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

						{/* Rendering Public Posts */}
						<div className="publicPostsWithComments">
							<h2>Public Posts With Comments</h2>
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

						{/* Rendering User Groups */}
						<div className="userGroups">
							<h2>Groups</h2>
							<ul>
								{userGroups !== null &&
									userGroups.map((userGroup) => (
										<li key={userGroup.createdAt}>
											{userGroup.Title}{" "}
											{/* Render whatever user properties you need */}
										</li>
									))}
							</ul>
						</div>
					</div>
					<div class="col-3">
						<div className="card">
							<Chat socket={socket} />
						</div>
					</div>
				</div>
			</div>




		</main>
	)
}
