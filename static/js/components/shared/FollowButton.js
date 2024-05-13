const {
  useState,
  useEffect
} = React;
export function FollowButton({
  followerId,
  user
}) {
  const [isFollowing, setIsFollowing] = useState(user.isFollowed);
  useEffect(() => {
    setIsFollowing(user.isFollowed);
  }, [user.isFollowed]);
  const handleFollowToggle = async () => {
    if (isFollowing) {
      // If already following, unfollow the user
      await handleUnfollow(followerId, user.userId);
      setIsFollowing(!isFollowing);
    } else {
      // If not following, follow the user
      if (user.isPublic) {
        await handleFollowPublic(followerId, user.userId);
        setIsFollowing(!isFollowing);
      } else {
        await handleFollowPrivate(followerId, user.userId);
      }
    }
    // Toggle the local follow state
  };
  const handleFollowPublic = async (followerId, userId) => {
    try {
      const bodyData = {
        followerId,
        subjectId: userId
      };
      const response = await fetch(`http://localhost:8080/api/userUsers`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(bodyData)
      });
      if (response.ok) {
        console.log("Successfully followed the user.");
        return true; // Return true if the follow request is successful
      } else {
        console.error("Failed to follow the user.");
      }
    } catch (error) {
      console.error("Error following the user:", error);
    }
    return false; // Return false if the follow request fails
  };
  const handleFollowPrivate = async (followerId, userId) => {
    try {
      const bodyData = {
        notificationType: "followRequest",
        objectId: userId,
        senderId: followerId,
        status: "pending",
        targetId: userId
      };
      const response = await fetch(`http://localhost:8080/api/notifications`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(bodyData)
      });
      if (response.ok) {
        console.log("Successfully sent follow notification to user", userId);
        return true; // Return true if the follow request is successful
      } else {
        console.error("Failed to send follow notification to user", userId);
      }
    } catch (error) {
      console.error("Error sending follow notification to user:", error);
    }
    return false; // Return false if the follow request fails
  };
  const handleUnfollow = async (followerId, userId) => {
    try {
      const response = await fetch(`http://localhost:8080/api/users/${followerId}/userUsers/${userId}`, {
        method: "DELETE",
        credentials: "include"
      });
      if (response.ok) {
        console.log("Successfully unfollowed the user.");
        return true; // Return true if the follow request is successful
      } else {
        console.error("Failed to unfollow the user.");
      }
    } catch (error) {
      console.error("Error following the user:", error);
    }
    return false; // Return false if the follow request fails
  };
  return /*#__PURE__*/React.createElement("button", {
    className: "btn btn-primary btn-sm",
    onClick: handleFollowToggle
  }, isFollowing ? "Unfollow" : "Follow");
}