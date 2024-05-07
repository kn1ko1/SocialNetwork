import { CommentCard } from "./CommentCard.js"
import { formattedDate } from "../shared/FormattedDate.js";
const { useState } = React

const postCardStyle = {
	maxWidth: '600px',
	background: 'linear-gradient(to bottom, #c7ddef, #ffffff)', // Light blue/grey to white gradient
	borderRadius: '10px',
	boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', // Optional: Add shadow for depth
	padding: '20px',
	margin: 'auto',
	marginBottom: '20px', // Adjust spacing between post cards
};

export function PostCard({ post, comments, showCommentForm }) {
	const [body, setBody] = useState("")
	const [selectedFile, setSelectedFile] = useState(null)

	const postDate = formattedDate(post.createdAt)

	const submit = async (e) => {
		e.preventDefault() // prevent reload.

		const formData = new FormData()

		// Append form data
		formData.append("body", body)
		formData.append("postId", post.postId)
		if (selectedFile) {
			formData.append("image", selectedFile)
		}

		console.log("Form data being sent to backend: ", formData)

		// Send user data to golang api/PostHandler.go.
		await fetch("http://localhost:8080/api/comments", {
			method: "POST",
			credentials: "include",
			body: formData,
		})

		// Reset the form fields to their default state
		setBody("")
		setSelectedFile(null)

		document.getElementById("commentTextArea").value = ""
	}

	// Function to handle file selection
	const handleFileChange = (e) => {
		setSelectedFile(e.target.files[0])
		// const file = e.target.files[0];
	}

	const handleSelectFile = () => {
		const commentFileInput = document.getElementById(
			`commentFileInput${post.postId}`
		)
		commentFileInput.click()
	}

	return (
		<div className="card" style={postCardStyle}>
			<div className="card-body">
				<div className="d-flex flex-start align-items-center">
					{post.userAvatar ? (
						<img
							src={post.userAvatar}
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
						<div className="d-flex align-items-center mb-1">
							<a
								className="fw-bold text-primary mb-0 me-2"
								href="#"
								onClick={() => renderProfile(post.userId)}
							>
								{post.userId}
							</a>
						</div>
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
			{showCommentForm && (
				<div
					className="card-footer py-3 border-0"
					style={{ backgroundColor: "#f8f9fa" }}
				>
					<div className="d-flex flex-start w-100">
						<div className="form-outline w-100">
							<textarea
								className="form-control"
								id="commentTextArea"
								rows="4"
								style={{ background: "#fff" }}
								onChange={(e) => setBody(e.target.value)}
							>Reply here...</textarea>


						</div>
					</div>
					<div className="float-end mt-2 pt-1">
						<button
							type="button"
							className="btn btn-primary"
							onClick={handleSelectFile}
						>
							Select File
						</button>
						<span>{selectedFile ? selectedFile.name : "No file selected"}</span>
						<input
							type="file"
							id={`commentFileInput${post.postId}`}
							accept="image/*"
							style={{ display: "none" }}
							onChange={handleFileChange}
						/>
						<button
							type="submit"
							className="btn btn-primary btn-sm"
							onClick={submit}
						>
							Post comment
						</button>
					</div>
					{/* If there are comments then render them, otherwise... don't */}
					{comments && comments.length > 0 && (
						<div className="comments">
							<h2>Comments</h2>
							{comments.map((comment) => (
								<CommentCard key={comment.createdAt} comment={comment} />
							))}
						</div>
					)}
				</div>
			)}
		</div >
	)
}