import { formattedDate } from '../shared/FormattedDate.js';
const { useState, useEffect } = React;
export function CommentCard({ comment }) {
  return (
    <div className="card mt-3">
      <div className="d-flex flex-start align-items-center">
        {comment.user.imageURL ? (
          <img
            src={comment.user.imageURL}
            className="rounded-circle shadow-1-strong me-3 img-fluid rounded-circle border border-2"
            width="60"
            height="60"
            style={{ padding: '5px' }}
          />
        ) : (
          <img
            src="https://static-00.iconduck.com/assets.00/avatar-default-symbolic-icon-479x512-n8sg74wg.png"
            className="rounded-circle shadow-1-strong me-3 img-fluid rounded-circle border border-2"
            width="60"
            height="60"
            style={{ padding: '5px' }}
          />
        )}
        <div>
          <h6
            className="fw-bold text-primary mb-1"
            onClick={() => renderProfile(comment.comment.userId)}
          >
            {comment.user.username}
          </h6>
          <p className="text-muted small mb-0">
            {formattedDate(comment.comment.createdAt)}
          </p>
        </div>
      </div>
      {comment.comment.imageURL && (
        <div className="mt-3 mb-2 pb-1">
          <img
            src={comment.comment.imageURL}
            className="img-fluid"
            alt="comment"
          />
        </div>
      )}
      <div className="card-body">
        <p className="card-text">{comment.comment.body}</p>
      </div>
    </div>
  );
}
