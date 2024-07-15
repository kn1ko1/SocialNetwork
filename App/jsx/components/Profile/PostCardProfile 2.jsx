import { formattedDate } from "../shared/FormattedDate.js";

const postCardStyle = {
  maxWidth: '600px',
  background: 'linear-gradient(to bottom, #c7ddef, #ffffff)', // Light blue/grey to white gradient
  borderRadius: '10px',
  boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', // Optional: Add shadow for depth
  padding: '20px',
  margin: 'auto',
  marginBottom: '20px', // Adjust spacing between post cards
};

// a paired down version of postCard, for profile
export function PostCardProfile({ post }) {

  const postDate = formattedDate(post.createdAt);


  return (
    <div className="card" style={postCardStyle}>
      <div className="card-body">
        <div className="d-flex flex-start align-items-center">

          <div>
            <p className="text-muted small mb-0">{postDate}</p>
          </div>
        </div>
        {/* Image, if there is one */}
        {!post.imageURL ? null : (
          <p className="mt-3 mb-2 pb-1">
            <img src={post.imageURL} className="img-fluid" />
          </p>
        )}
        {/* Post Body */}
        <p className="mt-3 mb-2 pb-1">{post.body}</p>
      </div>
    </div>
  );
}
