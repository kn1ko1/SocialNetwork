import { getCurrentUserId } from "./shared/getCurrentUserId.js"
import { FollowButton } from "./components/FollowButton.js"
const { useState, useEffect } = React

export function Profile({ userId, isEditable }) {
	const { currentUserId, isLoading, error } = getCurrentUserId()
	const [profileUserData, setProfileUserData] = useState({})
	const [userPostData, setUserPostData] = useState([])
	const [userFollowerData, setUserFollowerData] = useState([])
	const [userFollowsData, setUserFollowsData] = useState([])
	const [isPublicValue, setIsPublicValue] = useState(null)
	const [isFollowed, setIsFollowed] = useState(false)

	useEffect(() => {
		fetchProfileData()
	}, [userId])

	useEffect(() => {
		if (!isPublicValue && !isEditable && currentUserId) {
			checkIfFollowed(currentUserId)
		}
	}, [isPublicValue, isEditable, currentUserId])

	const fetchProfileData = async () => {
		try {
			const response = await fetch(
				`http://localhost:8080/api/profile/${userId}`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
					},
				}
			)

			if (!response.ok) {
				throw new Error(
					`Failed to fetch profile data: ${response.status} ${response.statusText}`
				)
			}

			const data = await response.json()
			setProfileUserData(data.profileUserData)
			setUserPostData(data.userPostData || [])
			setUserFollowerData(data.userFollowerData || [])
			setUserFollowsData(data.userFollowsData || [])
			setIsPublicValue(data.profileUserData.isPublic)
			console.log("This is my data with followers", data)
		} catch (error) {
			console.error("Error fetching profile data:", error)
		}
	}

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
				setIsFollowed(true)
				console.log("checkIfFollowed.  isFollowed", isFollowed)
				console.log("response", response)
			} else if (response.status === 404) {
				setIsFollowed(false)
				console.log("checkIfFollowed.  isFollowed", isFollowed)
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

	return (
		<div>
			<div id="profileData">
				<h2>{profileUserData.username}'s Profile</h2>
				{!isEditable && (
					<FollowButton
						followerId={currentUserId}
						subjectId={userId}
						isFollowed={isFollowed}
					/>
				)}
				{isPublicValue || isEditable || isFollowed ? (
					<>
						{isEditable ? (
							<div id="isPublicToggle">
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
							</div>
						) : (
							<p>
								<strong>Privacy:</strong> {isPublicValue ? "Public" : "Private"}
							</p>
						)}

						<p>
							<strong>User ID:</strong> {profileUserData.userId}
						</p>
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
						<p>
							<strong>Image URL:</strong> {profileUserData.imageURL}
						</p>

						<h2>{profileUserData.username}'s Posts</h2>
						<div id="myPostsData">
							{userPostData.map((post) => (
								<div key={post.postId}>
									<p>
										<strong>Post ID:</strong> {post.postId}
									</p>
									<p>
										<strong>Created At:</strong> {post.createdAt}
									</p>
									<p>
										<strong>Body:</strong> {post.body}
									</p>
									<p>
										<strong>Image URL:</strong> {post.imageURL}
									</p>
								</div>
							))}
						</div>

						<h2>{profileUserData.username}'s Followers</h2>
						<div id="myFollowersData">
							{userFollowerData &&
								userFollowerData.map((follower) => (
									<p key={follower.username}>{follower.username}</p>
								))}
						</div>

						<h2>{profileUserData.username}'s Followed</h2>
						<div id="usersIFollowData">
							{userFollowsData &&
								userFollowsData.map((user) => (
									<p key={user.username}>{user.username}</p>
								))}
						</div>
					</>
				) : (
					<p>This profile is private.</p>
				)}
			</div>
		</div>
	)
}