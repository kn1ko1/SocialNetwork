const { useState, useEffect } = React

export function FollowButton({ followerId, user }) {
	const [isFollowing, setIsFollowing] = useState(user.isFollowed)
	useEffect(() => {
		setIsFollowing(user.isFollowed)
	}, [user.isFollowed])

	const handleFollowToggle = async () => {
		if (isFollowing) {
			// If already following, unfollow the user
			await handleUnfollow(followerId, user.userId)
			setIsFollowing(!isFollowing)
		} else {
			// If not following, follow the user
			if (user.isPublic) {
				await handleFollowPublic(followerId, user.userId)
				setIsFollowing(!isFollowing)
			} else {
				await handleFollowPrivate(followerId, user.userId)
			}

		}
		// Toggle the local follow state

	}

	const handleFollowPublic = async (followerId, userId) => {
		try {
			const bodyData = { followerId, userId };
			const response = await fetch(
				`http://localhost:8080/api/userUsers`,
				{
					method: "POST",
					credentials: "include",
					body: JSON.stringify(bodyData),
				}
			)

			if (response.ok) {
				console.log("Successfully followed the user.")
				return true // Return true if the follow request is successful
			} else {
				console.error("Failed to follow the user.")
			}
		} catch (error) {
			console.error("Error following the user:", error)
		}

		return false // Return false if the follow request fails
	}

	async function AddGroupUser(userId, groupId, notificationType) {


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


	const handleFollowPrivate = async (followerId, userId) => {
		const notificationtData = {
			notificationType: "followRequest",
			objectId: followerId,
			senderId: followerId,
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
				console.log('Follow notification added successfully!');
			} else {
				// Handle error response
				console.error('Failed to send follow notification:', response.statusText);
			}
		} catch (error) {
			console.error('Error adding group user:', error);
		}
	}

	const handleUnfollow = async (followerId, userId) => {
		try {
			const response = await fetch(
				`http://localhost:8080/api/users/${followerId}/userUsers/${userId}`,
				{
					method: "DELETE",
					credentials: "include",
				}
			)

			if (response.ok) {
				console.log("Successfully unfollowed the user.")
				return true // Return true if the follow request is successful
			} else {
				console.error("Failed to unfollow the user.")
			}
		} catch (error) {
			console.error("Error following the user:", error)
		}

		return false // Return false if the follow request fails
	}

	return (
		<button className="btn btn-primary btn-sm" onClick={handleFollowToggle}>
			{isFollowing ? "Unfollow" : "Follow"}
		</button>
	)
}