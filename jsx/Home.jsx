const { useState, useEffect } = React;
import { getCurrentUserId } from './components/shared/GetCurrentUserId.js';
import { formattedDate } from './components/shared/FormattedDate.js';
import { PostForm } from './components/Home/PostForm.js';
import { PostCard } from './components/shared/PostCard.js';
import { FollowButton } from './components/shared/FollowButton.js';
import { renderProfile } from './Profile.js';

export const renderHome = ({ socket }) => {
  const pageContainer = document.querySelector('.page-container');
  ReactDOM.render(<Home socket={socket} />, pageContainer);
};

// Display information relating to homepage
export function Home({ socket }) {
  const { currentUserId } = getCurrentUserId();
  const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
  const [privatePosts, setPrivatePosts] = useState([]);
  const [publicPosts, setPublicPosts] = useState([]);
  const [userList, setUserList] = useState([]);
  const [followedUsersList, setFollowedUsersList] = useState([]);
  const [userEvents, setUserEvents] = useState([]);

  const fetchUserData = async (currentUserId) => {
    try {
      const promises = [];
      promises.push(fetch(`http://localhost:8080/api/users`));
      promises.push(
        fetch(`http://localhost:8080/api/users/${currentUserId}/userUsers`)
      );
      promises.push(
        fetch(`http://localhost:8080/api/users/${currentUserId}/events`)
      );
      promises.push(fetch(`/api/posts/public/withComments`));
      promises.push(fetch(`/api/posts/private/withComments`));
      promises.push(fetch(`/api/posts/almostPrivate/withComments`));

      const results = await Promise.all(promises);

      const userListResponse = results[0];
      const followedUserListResponse = results[1];
      const userEventsResponse = results[2];
      const publicPostsResponse = results[3];
      const privatePostsResponse = results[4];
      const almostPrivatePostsResponse = results[5];

      if (!userListResponse.ok) {
        throw new Error('Failed to fetch user list');
      }
      if (!followedUserListResponse.ok) {
        throw new Error('Failed to fetch followed users list');
      }
      if (!userEventsResponse.ok) {
        throw new Error('Failed to fetch users events');
      }
      if (!publicPostsResponse.ok) {
        throw new Error('Failed to fetch public posts');
      }
      if (!privatePostsResponse.ok) {
        throw new Error('Failed to fetch private posts');
      }
      if (!almostPrivatePostsResponse.ok) {
        throw new Error('Failed to fetch almost private posts');
      }

      const userListData = await userListResponse.json();
      const followedUsersListData = await followedUserListResponse.json();
      const userEventsData = await userEventsResponse.json();
      const publicPostsData = await publicPostsResponse.json();
      const privatePostsData = await privatePostsResponse.json();
      const almostPrivatePostsData = await almostPrivatePostsResponse.json();

      let filteredFollowedUsers = null;
      let updatedUserListData = userListData;

      if (followedUsersListData != null) {
        filteredFollowedUsers = userListData.filter((user) =>
          followedUsersListData.some(
            (followedUser) => followedUser.subjectId === user.userId
          )
        );

        // Add isFollowed property to each user where userId matches subjectId
        updatedUserListData = userListData.map((user) => ({
          ...user,
          isFollowed: followedUsersListData.some(
            (followedUser) => followedUser.subjectId === user.userId
          ),
        }));
      } else {
        // If followedUsersList is null, set isFollowed to false for every user
        updatedUserListData = userListData.map((user) => ({
          ...user,
          isFollowed: false,
        }));
      }

      setUserList(updatedUserListData);
      setFollowedUsersList(filteredFollowedUsers);
      setUserEvents(userEventsData);
      setPublicPosts(publicPostsData);
      setPrivatePosts(privatePostsData);
      setAlmostPrivatePosts(almostPrivatePostsData);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };
  

  useEffect(() => {
    if (currentUserId != null) {
      fetchUserData(currentUserId);
    }
  }, [currentUserId]);

  const homeStyle = {
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
    <main className="homePage" style={opaqueStyle}>
      <div style={{ padding: '20px' }}>
        <PostForm
          groupId={0}
          followedUsers={followedUsersList}
          fetchFunc={() => fetchUserData(currentUserId)}
        />
      </div>
      <div className="container text-center">
        <div className="row align-items-start">
          <div className="col-3">
            <div className="userList" style={homeStyle}>
              <h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>
                User List
              </h2>
              {userList !== null && userList.length > 0 ? (
                userList
                  // Filter out the current user
                  .filter((user) => user.userId !== currentUserId)
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
                        socket={socket}
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

          <div className="col-6" style={opaqueStyle}>
            {/* Rendering Public Posts */}
            <div className="publicPosts">
              <h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>
                Public Posts
              </h2>
              {publicPosts !== null &&
              publicPosts.length > 0 ? (
                [...publicPosts]
                  .sort((a, b) => new Date(b.post.createdAt) - new Date(a.post.createdAt))
                  .map((publicPost) => (
                    <PostCard
                      key={`public-${publicPost.post.postId}`}
                      post={publicPost.post}
                      comments={publicPost.comments}
                      showCommentForm={true}
                      fetchFunc={() => fetchUserData(currentUserId)}
                    />
                  ))
              ) : (
                <p>No public posts</p>
              )}
            </div>

            {/* Rendering Private Posts */}
            <div className="privatePosts">
              <h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>
                Private Posts
              </h2>
              {privatePosts !== null && privatePosts.length > 0 ? (
                [...privatePosts]
                  .sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
                  .map((privatePost) => (
                    <PostCard
                      key={`private-${privatePost.post.id}`}
                      post={privatePost.post}
                      comments={privatePost.comments}
                      showCommentForm={true}
                    />
                  ))
              ) : (
                <p>No private posts</p>
              )}
            </div>

            {/* Rendering Almost Private Posts */}
            <div className="almostPrivatePosts">
              <h2 style={{ textDecoration: 'underline', textAlign: 'center' }}>
                Almost Private Posts
              </h2>
              {almostPrivatePosts !== null && almostPrivatePosts.length > 0 ? (
                [...almostPrivatePosts]
                  .sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
                  .map((almostPrivatePost) => (
                    <PostCard
                      key={`almost-private-${almostPrivatePost.post.id}`}
                      post={almostPrivatePost.post}
                      comments={almostPrivatePost.comments}
                      showCommentForm={true}
                    />
                  ))
              ) : (
                <p>No almost private posts</p>
              )}
            </div>
          </div>

          <div className="col-3">
            <div className="userEvents" style={homeStyle}>
              <h2 style={{ textDecoration: 'underline' }}>Events</h2>
              {userEvents !== null && userEvents.length > 0 ? (
                <div className="accordion" id="eventAccordion">
                  {userEvents.map((event, index) => (
                    <div key={index} className="accordion-item">
                      <h2 className="accordion-header" id={`heading${index}`}>
                        <button
                          className="accordion-button collapsed"
                          type="button"
                          data-bs-toggle="collapse"
                          data-bs-target={`#collapse${index}`}
                          aria-expanded="false"
                          aria-controls={`collapse${index}`}
                        >
                          {event.title}
                        </button>
                      </h2>
                      <div
                        id={`collapse${index}`}
                        className="accordion-collapse collapse"
                        aria-labelledby={`heading${index}`}
                        data-bs-parent="#eventAccordion"
                      >
                        <div className="accordion-body">
                          <p>{event.description}</p>
                          <small>{formattedDate(event.dateTime)}</small>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              ) : (
                <p>No Events</p>
              )}
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}
