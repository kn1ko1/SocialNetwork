import { formattedDate } from "../../../static/js/shared/FormattedDate.js";
export function CommentCard({
  comment
}) {
  return /*#__PURE__*/React.createElement("div", {
    className: "card mt-3"
  }, /*#__PURE__*/React.createElement("div", {
    className: "d-flex flex-start align-items-center"
  }, comment.userAvatar ? /*#__PURE__*/React.createElement("img", {
    src: comment.userAvatar,
    className: "rounded-circle shadow-1-strong me-3 img-fluid rounded-circle",
    width: "60",
    height: "60"
  }) : /*#__PURE__*/React.createElement("img", {
    src: "https://static-00.iconduck.com/assets.00/avatar-default-symbolic-icon-479x512-n8sg74wg.png",
    className: "rounded-circle shadow-1-strong me-3 img-fluid rounded-circle",
    width: "60",
    height: "60"
  }), /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h6", {
    className: "fw-bold text-primary mb-1",
    onClick: () => renderProfile(comment.userId)
  }, comment.userId), /*#__PURE__*/React.createElement("p", {
    className: "text-muted small mb-0"
  }, formattedDate(comment.createdAt)))), comment.imageURL && /*#__PURE__*/React.createElement("div", {
    className: "mt-3 mb-2 pb-1"
  }, /*#__PURE__*/React.createElement("img", {
    src: comment.imageURL,
    className: "img-fluid",
    alt: "comment"
  })), /*#__PURE__*/React.createElement("div", {
    className: "card-body"
  }, /*#__PURE__*/React.createElement("p", {
    className: "card-text"
  }, comment.body)));
}