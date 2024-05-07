export function CommentCard({ comment }) {
	const formattedDate = new Date(comment.createdAt).toLocaleString()

	return (
		<div className="card mt-3">
			<div className="d-flex flex-start align-items-center">
				{comment.userAvatar ? (
					<img
						src={comment.userAvatar}
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
				<div>
					<h6
						className="fw-bold text-primary mb-1"
						onClick={() => renderProfile(comment.userId)}
					>
						{comment.userId}
					</h6>
					<p className="text-muted small mb-0">{formattedDate}</p>
				</div>
			</div>
			{comment.imageURL && (
				<div className="mt-3 mb-2 pb-1">
					<img src={comment.imageURL} className="img-fluid" alt="comment" />
				</div>
			)}
			<div className="card-body">
				<p className="card-text">{comment.body}</p>
			</div>
		</div>
	);
}