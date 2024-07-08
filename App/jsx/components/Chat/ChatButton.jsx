const { useState, useEffect } = React

// CURRENTLY NOT IMPLIMENTED
export function ChatButton( targetId ) {
	
	useEffect(() => {
		setIsFollowing(isFollowed)
	}, [isFollowed])

	const handleFollowToggle = async () => {
		if (isFollowing) {
			// If already following, unfollow the user
			await handleUnfollow(followerId, subjectId)
		} else {
			// If not following, follow the user
			await handleFollow(followerId, subjectId)
		}
		// Toggle the local follow state
		setIsFollowing(!isFollowing)
	}


	return (
		<button className="btn btn-primary btn-sm" onClick={handleFollowToggle}>
			{isFollowing ? "Unfollow" : "Follow"}
		</button>
	)
}