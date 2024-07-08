import { formattedDate } from "../shared/FormattedDate.js";
const postCardStyle = {
  maxWidth: '600px',
  background: 'linear-gradient(to bottom, #c7ddef, #ffffff)',
  // Light blue/grey to white gradient
  borderRadius: '10px',
  boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
  // Optional: Add shadow for depth
  padding: '20px',
  margin: 'auto',
  marginBottom: '20px' // Adjust spacing between post cards
};

// a paired down version of postCard, for profile
export function PostCardProfile({
  post
}) {
  const postDate = formattedDate(post.createdAt);
  return /*#__PURE__*/React.createElement("div", {
    className: "card",
    style: postCardStyle
  }, /*#__PURE__*/React.createElement("div", {
    className: "card-body"
  }, /*#__PURE__*/React.createElement("div", {
    className: "d-flex flex-start align-items-center"
  }, /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("p", {
    className: "text-muted small mb-0"
  }, postDate))), !post.imageURL ? null : /*#__PURE__*/React.createElement("p", {
    className: "mt-3 mb-2 pb-1"
  }, /*#__PURE__*/React.createElement("img", {
    src: post.imageURL,
    className: "img-fluid"
  })), /*#__PURE__*/React.createElement("p", {
    className: "mt-3 mb-2 pb-1"
  }, post.body)));
}