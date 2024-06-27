import { CommentCard } from '../Home/CommentCard.js';
import { formattedDate } from './FormattedDate.js';
import { renderProfile } from '../../Profile.js';
const { useState } = React;

const postCardStyle = {
  maxWidth: '600px',
  background: 'linear-gradient(to bottom, #c7ddef, #ffffff)', // Light blue/grey to white gradient
  borderRadius: '10px',
  boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', // Optional: Add shadow for depth
  padding: '20px',
  margin: 'auto',
  marginBottom: '20px', // Adjust spacing between post cards
};

export function PostCard({ post, comments, showCommentForm, fetchFunc }) {
  const [body, setBody] = useState('');
  const [selectedFile, setSelectedFile] = useState(null);

  const postDate = formattedDate(post.post.createdAt);

  const submit = async (e) => {
    e.preventDefault(); // prevent reload.

    const formData = new FormData();

    // Append form data
    formData.append('body', body);
    formData.append('postId', post.post.postId);
    if (selectedFile) {
      formData.append('image', selectedFile);
    }

    console.log('Form data being sent to backend: ', formData);

    // Send user data to golang api/PostHandler.go.
    await fetch('http://localhost:8080/api/comments', {
      method: 'POST',
      credentials: 'include',
      body: formData,
    });

    // Reset the form fields to their default state
    setBody('');
    setSelectedFile(null);

    document.getElementById('commentTextArea').value = '';
    fetchFunc();
  };

  // Function to handle file selection
  const handleFileChange = (e) => {
    setSelectedFile(e.target.files[0]);
    // const file = e.target.files[0];
  };

  const handleSelectFile = () => {
    const commentFileInput = document.getElementById(
      `commentFileInput${post.post.postId}`
    );
    commentFileInput.click();
  };


  return (
    <div className="card" style={postCardStyle}>
      <div className="card-body">
        <div className="d-flex flex-start align-items-center">
          {showCommentForm && (
            <>
              {post.user.imageURL ? (
                <img
                  src={post.user.imageURL}
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
            </>
          )}

          <div>
            <div className="d-flex align-items-center mb-1">
              <a
                className="fw-bold text-primary mb-0 me-2"
                href="#"
                onClick={() => renderProfile(post.post.userId)}
              >
                {post.user.username}
              </a>
            </div>
            <p className="text-muted small mb-0">{postDate}</p>
          </div>
        </div>
        {/* Image, if there is one */}
        {!post.post.imageURL ? null : (
          <p className="mt-3 mb-2 pb-1">
            <img src={post.post.imageURL} className="img-fluid" />
          </p>
        )}
        {/* Post Body */}
        <p className="mt-3 mb-2 pb-1">{post.post.body}</p>
      </div>
      {showCommentForm && (
        <div
          className="card-footer py-3 border-0"
          style={{
            backgroundColor: '#f8f9fa',
            borderRadius: '10px',
            border: '1px solid #ccc',
          }}
        >
          <div className="d-flex flex-start w-100">
            <div className="form-outline w-100">
              <textarea
                className="form-control"
                id="commentTextArea"
                rows="2"
                style={{ background: '#fff' }}
                placeholder="Reply here..."
                onChange={(e) => setBody(e.target.value)}
              ></textarea>
            </div>
          </div>
          <div style={{ marginTop: '20px', paddingTop: '10px' }}>
            <button
              type="button"
              className="btn btn-primary"
              onClick={handleSelectFile}
              style={{ marginRight: '10px' }}
            >
              Select File
            </button>
            <span style={{ marginRight: '10px' }}>
              {selectedFile ? selectedFile.name : 'No file selected'}
            </span>
            <input
              type="file"
              id={`commentFileInput${post.post.postId}`}
              accept="image/*"
              style={{ display: 'none' }}
              onChange={handleFileChange}
            />
            <button
              type="submit"
              className="btn btn-primary btn-sm"
              onClick={submit}
              style={{ marginTop: '10px', marginBottom: '10px' }}
            >
              Post comment
            </button>
          </div>

          {/* If there are comments then render them, otherwise... don't */}
          {comments && comments.length > 0 && (
            <div className="comments" style={{ marginTop: '20px' }}>
              <h4>Comments</h4>
              <div
                style={{
                  display: 'flex',
                  flexDirection: 'column',
                  gap: '10px',
                }}
              >
                {comments.map((comment) => (
                  <CommentCard key={comment.createdAt} comment={comment} />
                ))}
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
