import { formattedDate } from '../shared/FormattedDate.js';
export function CommentCard({
  comment
}) {
  return /*#__PURE__*/React.createElement("div", {
    className: "card mt-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "d-flex flex-start align-items-center"
  }, comment.user.imageURL ? /*#__PURE__*/React.createElement("img", {
    src: comment.user.imageURL,
    className: "rounded-circle shadow-1-strong me-3 img-fluid rounded-circle border border-2",
    width: "60",
    height: "60",
    style: {
      padding: '5px'
    }
  }) : /*#__PURE__*/React.createElement("img", {
    src: "https://static-00.iconduck.com/assets.00/avatar-default-symbolic-icon-479x512-n8sg74wg.png",
    className: "rounded-circle shadow-1-strong me-3 img-fluid rounded-circle border border-2",
    width: "60",
    height: "60",
    style: {
      padding: '5px'
    }
  }), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h6", {
    className: "fw-bold text-primary mb-1",
    onClick: () => renderProfile(comment.comment.userId)
  }, comment.user.username), /*#__PURE__*/React.createElement("p", {
    className: "text-muted small mb-0"
  }, formattedDate(comment.comment.createdAt)))), comment.comment.imageURL && /*#__PURE__*/React.createElement("div", {
    className: "mt-3 mb-2 pb-1"
  }, /*#__PURE__*/React.createElement("img", {
    src: comment.comment.imageURL,
    className: "img-fluid",
    alt: "comment"
  })), /*#__PURE__*/React.createElement("div", {
    className: "card-body"
  }, /*#__PURE__*/React.createElement("p", {
    className: "card-text"
  }, comment.comment.body)));
}