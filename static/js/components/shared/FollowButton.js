const {
  useState,
  useEffect
} = React;
export function FollowButton({
  followerId,
  subjectId,
  isFollowed
}) {
  const [isFollowing, setIsFollowing] = useState(isFollowed);
  useEffect(() => {
    setIsFollowing(isFollowed);
  }, [isFollowed]);
  const handleFollowToggle = async () => {
    if (isFollowing) {
      // If already following, unfollow the user
      await handleUnfollow(followerId, subjectId);
    } else {
      // If not following, follow the user
      await handleFollowPublic(followerId, subjectId);
    }
    // Toggle the local follow state
    setIsFollowing(!isFollowing);
  };
  const handleFollowPublic = async (followerId, subjectId) => {
    try {
      const bodyData = {
        followerId,
        subjectId
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
  const handleFollowPrivate = async (followerId, subjectId) => {
    try {
      const response = await fetch(`http://localhost:8080/api/users/${followerId}/userUsers/`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
          subjectId
        })
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
  const handleUnfollow = async (followerId, subjectId) => {
    try {
      const response = await fetch(`http://localhost:8080/api/users/${followerId}/userUsers/${subjectId}`, {
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