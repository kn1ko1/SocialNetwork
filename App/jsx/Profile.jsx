import { getCurrentUserId } from "./components/shared/GetCurrentUserId.js"
import { FollowButton } from "./components/shared/FollowButton.js"
import { PostCardProfile } from "./components/Profile/PostCardProfile.js"
const { useState, useEffect } = React

export const renderProfile = (socket, userId, isEditable) => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(
		<Profile socket={socket} userId={userId} isEditable={isEditable} />,
		pageContainer
	)
}

export function Profile({ socket, userId, isEditable }) {
	const { currentUserId } = getCurrentUserId()
	const [profileUserData, setProfileUserData] = useState({})
	const [userPostData, setUserPostData] = useState([])
	const [userFollowerData, setUserFollowerData] = useState([])
	const [userFollowsData, setUserFollowsData] = useState([])
	const [isPublicValue, setIsPublicValue] = useState(null)
	const [isFollowed, setIsFollowed] = useState(false)

	useEffect(() => {
		const fetchProfileData = async () => {
			try {

				const initialFetch = await fetch(`http://localhost:8080/api/users/${userId}`);
				const userData = await initialFetch.json();
				let userPostData = null
				let usersIFollowData = null
				let usersFollowMeData = null

				if (userData.isPublic || isEditable) {
					const promises = [];

					promises.push(fetch(`http://localhost:8080/api/users/${userId}/posts`));
					promises.push(fetch(`http://localhost:8080/api/users/${userId}/followedUsers`));
					promises.push(fetch(`http://localhost:8080/api/users/${userId}/followerUsers`));

					const results = await Promise.all(promises);

					const userPostResponse = results[0]
					const usersIFollowResponse = results[1];
					const usersFollowMeResponse = results[2];

					if (!userPostResponse.ok) {
						throw new Error('Failed to fetch user posts');
					}
					if (!usersIFollowResponse.ok) {
						throw new Error('Failed to fetch followed users list');
					}
					if (!usersFollowMeResponse.ok) {
						throw new Error('Failed to fetch follower users list');
					}


					// const userData = await userDataResponse.json();
					userPostData = await userPostResponse.json();
					usersIFollowData = await usersIFollowResponse.json();
					usersFollowMeData = await usersFollowMeResponse.json();
				}

				setProfileUserData(userData);
				setUserPostData(userPostData || []);
				setUserFollowerData(usersIFollowData || []);
				setUserFollowsData(usersFollowMeData || []);
				setIsPublicValue(userData.isPublic);
			} catch (error) {
				console.error('Error fetching user data:', error);
			}
		};

		fetchProfileData()


	}, [userId])

	useEffect(() => {
		if (!isEditable && currentUserId) {
			checkIfFollowed(currentUserId)
		}
	}, [currentUserId, isEditable, userId])


	const checkIfFollowed = async (currentUserId) => {
		try {
			const response = await fetch(
				`http://localhost:8080/api/users/${currentUserId}/userUsers/${userId}`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
					},
				}
			)

			if (response.ok) {
				setIsFollowed(true);
				setProfileUserData((prevData) => ({
					...prevData,
					isFollowed: true,
				}));
			} else if (response.status === 404) {
				setIsFollowed(false);
				setProfileUserData((prevData) => ({
					...prevData,
					isFollowed: false,
				}));

			} else {
				console.error("Error fetching user user data:", response.statusText)
			}
		} catch (error) {
			console.error("Error fetching user user data:", error)
		}
	}

	const handlePrivacyChange = (event) => {
		const newPrivacySetting = JSON.parse(event.target.value)

		setIsPublicValue(newPrivacySetting)

		fetch("http://localhost:8080/api/profile/privacy", {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				userId: profileUserData.userId,
				isPublic: newPrivacySetting,
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error("Failed to update privacy status")
				}
			})
			.catch((error) => {
				console.error("Error updating privacy status:", error)
				setIsPublicValue(!newPrivacySetting)
			})
	}

	const profileStyle = {
		maxWidth: '1300px',
		background: 'linear-gradient(to bottom, #c7ddef, #ffffff)', // Light blue/grey to white gradient
		borderRadius: '10px',
		boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', // Optional: Add shadow for depth
		padding: '40px',
		margin: 'auto',
		marginBottom: '20px', // Adjust spacing between post cards
		border: '1px solid #ccc', // Add a thin border
	};


	const opaqueStyle = {
		backgroundColor: 'rgba(255, 255, 255, 0.25)', // Adjust the opacity here 
		maxWidth: '1300px',
		borderRadius: '10px',
		boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', // Optional: Add shadow for depth
		padding: '40px',
		margin: 'auto',
		marginBottom: '20px', // Adjust spacing between post cards
	  };

	return (
		<div className="container-fluid" style={opaqueStyle}>
			<div className="row" >
				<div className="col-md-3" style={{...profileStyle, margin: "0 0 20px 0" }} >
					{/* User data */}
					<h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>{profileUserData.username}'s Profile</h2>
					<br />

					<p className="d-flex justify-content-center">
								
								{profileUserData.imageURL ? (
									<img
										src={profileUserData.imageURL}
										className="rounded-circle shadow-1-strong me-3 img-fluid rounded-circle"
										width="60"
										height="60"
									/>
								) : (
									<img
										src="https://static-00.iconduck.com/assets.00/avatar-default-symbolic-icon-479x512-n8sg74wg.png"
										className="rounded-circle shadow-1-strong me-3 img-fluid rounded-circle"
										width="60"
										height="60"
									/>
								)}

							</p>

					{!isEditable && (
						<div className="d-flex justify-content-center align-items-center">
							<FollowButton
								socket={socket}
								followerId={currentUserId}
								user={profileUserData}
							/>
						</div>
					)}
					{isPublicValue || isEditable || isFollowed ? (
						<>
							{isEditable ? (
								<div id="isPublicToggle">
									<p>
										<h4 style={{ fontSize: '14px' }}>Toggle to change profile privacy setting</h4>
										<strong>Privacy:</strong>
										<label>
											<input
												type="radio"
												value={true}
												checked={isPublicValue === true}
												onChange={handlePrivacyChange}
											/>
											Public
										</label>
										<label>
											<input
												type="radio"
												value={false}
												checked={isPublicValue === false}
												onChange={handlePrivacyChange}
											/>
											Private
										</label>
									</p>
								</div>
							) : (
								<p>
									<strong>Privacy:</strong> {isPublicValue ? "Public" : "Private"}
								</p>
							)}
							<p>
								<strong>Username:</strong> {profileUserData.username}
							</p>
							<p>
								<strong>Email:</strong> {profileUserData.email}
							</p>
							<p>
								<strong>First Name:</strong> {profileUserData.firstName}
							</p>
							<p>
								<strong>Last Name:</strong> {profileUserData.lastName}
							</p>
							<p>
								<strong>Date of Birth:</strong>{" "}
								{new Date(profileUserData.dob).toLocaleDateString()}
							</p>
							<p>
								<strong>Bio:</strong> {profileUserData.bio}
							</p>
							
						</>
					) : (
						<p>This profile is private.</p>
					)}
				</div>

				<div className="col-md-6">
					<div style={opaqueStyle}>
					{/* Posts data */}
					<h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>{profileUserData.username}'s Posts</h2>
					<div id="myPostsData">
						{userPostData.map((post) => (
							<div key={post.postId}>
								<PostCardProfile post={post} />
							</div>
						))}
					</div>
				</div>
				</div>

				<div className="col-md-3">
					{/* Followers data */}
					<div style={profileStyle}>
					<h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>{profileUserData.username}'s Followers</h2>
					<div id="myFollowersData">
						{userFollowerData &&
							userFollowerData.map((follower) => (
								<p key={follower.username}>{follower.username}</p>
							))}
					</div>
					</div>

					{/* Followed data */}
					<div style={profileStyle}>
					<h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>Users {profileUserData.username} Follows</h2>
					<div id="usersIFollowData">
						{userFollowsData &&
							userFollowsData.map((user) => (
								<p key={user.username}>{user.username}</p>
							))}
					</div>
					</div>
				</div>
			</div>
		</div>
	)
}
