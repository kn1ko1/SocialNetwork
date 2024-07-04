export const fetchCommentsForPosts = async (postsData) => {
  try {
    // Iterate over each post in the postsData array
    for (const post of postsData) {
      const postId = post.postId; // Assuming postId is the identifier for each post

      // Fetch comments for the current post
      const response = await fetch(`http://localhost:8080/api/posts/${postId}/comments`);
      if (!response.ok) {
        throw new Error(`Failed to fetch comments for post with ID ${postId}`);
      }
      const commentsData = await response.json();

      // Update the post object with its comments
      post.comments = commentsData;
    }

    // After fetching comments for all posts, update the state or perform further actions
    // For example, you can set the state with the updated postsData
    return postsData
  } catch (error) {
    console.error('Error fetching comments for posts:', error);
  }
};
