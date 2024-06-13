const { useState, useEffect } = React
// PostForm component
// This component renders a form for creating a new post.
// It accepts a `groupId` prop to determine the group for the post
// and followedUsers so that you can assign users to Almost Private Posts
export function PostForm({ groupId, followedUsers, fetchFunc }) {
	const [body, setBody] = useState("");
	const [privacy, setPrivacy] = useState("");
	const [selectedFile, setSelectedFile] = useState(null);
	const [selectedUserIds, setSelectedUserIds] = useState([]);
	const [showFollowedUsersList, setShowFollowedUsersList] = useState(false);
	const [followedUsersForAP, setFollowedUsersForAP] = useState(followedUsers || []);

	useEffect(() => {
		setFollowedUsersForAP(followedUsers);
	}, [followedUsers]);

	const handleCheckboxChange = (e) => {
		const userId = e.target.value;
		const isChecked = e.target.checked;

		if (isChecked) {
			setSelectedUserIds((prevSelectedUserIds) => [...prevSelectedUserIds, userId]);
		} else {
			setSelectedUserIds((prevSelectedUserIds) =>
				prevSelectedUserIds.filter((id) => id !== userId)
			);
		}
	};

	// Handler for form submission
	const submit = async (e) => {
		e.preventDefault() // Prevent page reload

		const formData = new FormData()

		// Append form data
		formData.append("body", body)
		formData.append("privacy", privacy)
		if (privacy === "private") {
			groupId = -1 // Set groupId to -1 for private posts
		}
		if (privacy === "almost private") {
			groupId = -2; // Set groupId to -2 for almost private posts
			formData.append("almostPrivatePostUsers", JSON.stringify(selectedUserIds));
		}
		formData.append("groupId", groupId);
		if (selectedFile) {
			formData.append("image", selectedFile)
		}

		console.log("Form data being sent to backend: ", formData)

		try {
			// Send user data to the server
			await fetch("http://localhost:8080/api/posts", {
				method: "POST",
				credentials: "include",
				body: formData,
			})

			// Reset form fields after successful submission
			setBody("");
			setPrivacy("public");
			setSelectedFile(null);
			setSelectedUserIds([]);
			document.getElementById("postFormBody").value = "";
			setShowFollowedUsersList(false);
		} catch (error) {
			console.error("Error submitting post:", error)
		}
		// Really we should do something clever with websockets and updating useStates, but this is much easier
		fetchFunc()
	};

	const handlePrivacyChange = (e) => {
		const newValue = e.target.value;
		setPrivacy(newValue);
		if (newValue === 'almost private') {
			setShowFollowedUsersList(true);
		} else {
			setShowFollowedUsersList(false);
		}
	};


	// Handler for file selection
	const handleFileChange = (e) => {
		setSelectedFile(e.target.files[0])
	}

	const handleSelectFile = () => {
		const fileInput = document.getElementById("fileInput");
		fileInput.click();
	};

	const followedUsersList = showFollowedUsersList ? (
		followedUsersForAP !== null && followedUsersForAP.length > 0 ? (
			<ul>
				{followedUsersForAP.map((followedUser) => (
					<li key={followedUser.username}>
						<label>
							<input
								type="checkbox"
								value={followedUser.userId}
								onChange={handleCheckboxChange}
							/>
							{followedUser.username}
						</label>
					</li>
				))}
			</ul>
		) : (
			<p className="text-muted">No followed users</p>
		)
	) : null;

	return (
		<div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '40vh', padding: '10px' }}>
		<main className="postForm container" style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', padding: '0'}}>
		  <div className="border" style={{ borderRadius: "10px", boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', border: "3px solid #333", padding: "10px", width: '100%', maxWidth: '600px', background: 'linear-gradient(to bottom, #c7ddef, #ffffff)'  }}>
			<div className="col-12">
			  <h1 className="h3 mb-3 fw-normal" style={{ textDecoration: 'underline', textAlign: "center" }}>New Post</h1>
			  <form onSubmit={submit}>
				<div className="form-floating mb-3" style={{ textAlign: "center" }}>
				  <input
							type="text"
							className="form-control"
							id="postFormBody"
							placeholder="Type your post here..."
							onChange={(e) => setBody(e.target.value)}
						/>
					</div>
					<div>
						<button
							type="button"
							className="btn btn-primary"
							onClick={handleSelectFile}
							style={{ marginRight: "10px" }}
						>
							Select File
						</button>
						<span>{selectedFile ? selectedFile.name : "No file selected"}</span>
						<input
							type="file"
							id="fileInput"
							accept="image/*"
							style={{ display: "none" }}
							onChange={handleFileChange}
						/>
					</div>
					<br /> {/* Line break */}
					<div className="form-floating mb-3">
						<div className="form-check">
							<input
								required
								type="radio"
								id="post-public-status"
								value="public"
								name="status"
								checked={privacy === "public"}
								onClick={handlePrivacyChange}
								className="form-check-input"
							/>
							<label htmlFor="post-public-status" className="form-check-label">
								Public
							</label>
						</div>
						<div className="form-check">
							<input
								required
								type="radio"
								id="post-private-status"
								value="private"
								name="status"
								checked={privacy === "private"}
								onClick={handlePrivacyChange}
								className="form-check-input"
							/>
							<label htmlFor="private-status" className="form-check-label">
								Private
							</label>
						</div>
						<div className="form-check">
							<input
								required
								type="radio"
								id="post-almostPrivate-status"
								value="almost private"
								name="status"
								checked={privacy === "almost private"}
								onClick={handlePrivacyChange}
								className="form-check-input"
							/>
							<label htmlFor="private-status" className="form-check-label">
								Almost Private
							</label>
						</div>
					</div>
					
					{followedUsersList}
					<button className="w-100 btn btn-lg btn-primary" type="submit">
						Submit
					</button>
				</form>
				</div>
				</div>
			</main>
		</div>
	)
}