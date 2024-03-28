const { useState, useEffect } = React

const App = () => {
	return (
		<div className="app-container">
			<Login />
		</div>
	)
}

// function Dummy(props) {
// 	return (
// 		<div className="container-fluid text-center">
// 		<div className="row mb-2">
// 			<div className="col-lg-3">
// 				<h1>Welcome</h1>
// 			</div>
// 			<div className="col-lg-6">
// 				<h1>Welcome</h1>
// 			</div>
// 			<div className="col-lg-3">
// 				<h1>Welcome</h1>
// 			</div>
// 		</div>
// 		<div className="row">
// 			<div className="col-6">
// 				<h1>Welcome</h1>
// 			</div>
// 			<div className="col-6">
// 				<h1>Welcome</h1>
// 			</div>
// 		</div>
// 	</div>
// 	)
// }

function Navbar() {

	const renderHome = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Home />, appContainer);
	};

	const renderProfile = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Profile />, appContainer);
	};

	const renderNotifications = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Notifications />, appContainer);
	};

	const renderChat = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Chat />, appContainer);
	};

	const renderGroup = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Group />, appContainer);
	};

	const logout = async () => {

		try {
			const response = await fetch("http://localhost:8080/auth/logout", {
				method: "POST",
				credentials: "include",
			});

			if (response.ok) {
				const appContainer = document.querySelector('.app-container');
				ReactDOM.render(<Login />, appContainer);
				console.log("Logout successful!");
			} else {
				console.log("Failed to logout. Server response not OK.");
			}
		} catch (error) {
			console.error("An error occurred during logout:", error);
		}
	};

	return (
		<nav className="navbar navbar-expand-md bg-body-tertiary">
			<div className="container-fluid">
				<button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
					<span className="navbar-toggler-icon"></span>
				</button>
				<div className="collapse navbar-collapse" id="navbarSupportedContent">
					<ul className="navbar-nav me-auto mx-auto mb-2 mb-lg-0">
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderProfile}>PROFILE</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderHome}>HOME</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderNotifications}>NOTIFICATIONS</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderChat}>CHAT</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderGroup}>GROUP</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={logout}>LOGOUT</a>
						</li>
					</ul>
				</div>
			</div>
		</nav>
	)
}

function Login() {
	const [usernameOrEmail, setUsernameOrEmail] = useState('');
	const [password, setPassword] = useState('');
	const [isLoggedIn, setIsLoggedIn] = useState(false);
	const [errorMessage, setErrorMessage] = useState('');

	const handleUsernameOrEmailChange = (e) => {
		setUsernameOrEmail(e.target.value);
	};

	const handlePasswordChange = (e) => {
		setPassword(e.target.value);
	};

	const handleSubmit = async (e) => {
		e.preventDefault();

		const userToLogin = {
			usernameOrEmail,
			password,
		};

		try {
			const response = await fetch('http://localhost:8080/auth/login', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify(userToLogin),
			});

			if (!response.ok) {
				setErrorMessage('Invalid credentials');
				throw new Error('Invalid credentials');
			}

			const data = await response.json();
			if (data.success) {
				setIsLoggedIn(true);
				setErrorMessage('');
			} else {
				setErrorMessage('Invalid credentials');
				throw new Error('Invalid credentials');
			}
		} catch (error) {
			setErrorMessage('Invalid credentials');
		}
	};

	const renderRegister = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Register />, appContainer);
	};

	if (isLoggedIn) {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Home />, appContainer);
	}

	return (
		<div className="container login-container">
			<h1 className="h3 mb-3 fw-normal login-text">Log in</h1>
			<form onSubmit={handleSubmit}>
				<div className="mb-3">
					<label htmlFor="exampleInputEmail1" className="form-label">
						Email address
					</label>
					<input
						type="email"
						className="form-control form-control-lg"
						id="exampleInputEmail1"
						aria-describedby="emailHelp"
						onChange={handleUsernameOrEmailChange}
					/>
				</div>
				<div className="mb-3">
					<label htmlFor="exampleInputPassword1" className="form-label">
						Password
					</label>
					<input
						type="password"
						className="form-control form-control-lg"
						id="exampleInputPassword1"
						onChange={handlePasswordChange}
					/>
				</div>
				<button type="submit" className="btn btn-primary">
					Log in
				</button>
			</form>
			{errorMessage && <div className="error-message">{errorMessage}</div>}
			<br />
			<div className="mb3">
				<span className="login-text">Don't have an account? &nbsp;</span>
				<button type="button" className="btn btn-primary" onClick={renderRegister}>
					Register
				</button>
			</div>
		</div>
	);
}


function Register() {
	const [email, setEmail] = useState("");
	const [encryptedPassword, setEncryptedPassword] = useState("");
	const [firstName, setFirstName] = useState("");
	const [lastName, setLastName] = useState("");
	const [dob, setDob] = useState("");
	const [imageURL, setImageURL] = useState("");
	const [username, setUsername] = useState("");
	const [bio, setBio] = useState("");
	const [isPublic, setIsPublic] = useState("public");
	const [isRegistered, setIsRegistered] = useState(false);

	//this is register button
	const submit = async (e) => {
		e.preventDefault(); // prevent reload.

		// Create new user as JS object.
		const newUser = {
			email,
			encryptedPassword,
			firstName,
			lastName,
			dob,
			imageURL,
			username,
			bio,
			isPublic,
		};

		try {
			// Send user data to backend
			const response = await fetch("http://localhost:8080/auth/registration", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(newUser),
			});

			if (!response.ok) {
				throw new Error('Invalid credentials');
			}

			//takes response from backend and processes
			const data = await response.json();
			if (data.success) {
				setIsRegistered(true);
			} else {
				throw new Error('Invalid credentials');
			}
		} catch (error) {
			throw new Error('Invalid credentials')
		}
	};

	//if credentials frontend succesfully create a new user then we render home
	if (isRegistered) {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Home />, appContainer);
	}

	//this is the login button, when pressed will serve login form
	const renderLogin = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Login />, appContainer);
	};


	return (
		<div className="container login-container">
			<h1 className="h3 mb-3 fw-normal login-text">register</h1>
			<form onSubmit={submit}>



				<div className="mb-3">
					<label htmlFor="floatingInput">Email address</label>
					<input
						required
						type="email"
						className="form-control"
						id="floatingInput"
						placeholder="name@example.com"
						onChange={(e) => setEmail(e.target.value)} />
				</div>

				<div className="mb-3">
					<label htmlFor="regpassword">Password</label>
					<input
						required
						type="password"
						className="form-control reginput"
						id="regpassword"
						placeholder="Password"
						onChange={(e) => setEncryptedPassword(e.target.value)} />
				</div>

				<div className="mb-3">
					<label htmlFor="firstName">First Name</label>
					<input
						required
						type="text"
						className="form-control reginput"
						id="firstName"
						placeholder="John"
						onChange={(e) => setFirstName(e.target.value)} />
				</div>

				<div className="mb-3">
					<label htmlFor="lastName">Last Name</label>
					<input
						required
						type="text"
						className="form-control reginput"
						id="lastName"
						placeholder="Doe"
						onChange={(e) => setLastName(e.target.value)} />
				</div>

				<div className="mb-3">
					<label htmlFor="dob">Date of Birth</label>
					<input
						required
						type="date"
						className="form-control reginput"
						id="dob"
						placeholder="16/01/1998"
						onChange={(e) => setDob(e.target.value)} />
				</div>

				<div className="mb-3">
					<label htmlFor="imageURL">ImageURL</label>
					<input
						type="text"
						className="form-control reginput"
						id="imageURL"
						placeholder="https://..."
						onChange={(e) => setImageURL(e.target.value)} />
				</div>

				<div className="mb-3">
					<label htmlFor="username">Username</label>
					<input
						type="text"
						className="form-control reginput"
						id="username"
						placeholder="Johnny"
						onChange={(e) => setUsername(e.target.value)} />
				</div>

				<div className="form-check">
					<input
						className="form-check-input"
						type="radio"
						id="public-status"
						value="public"
						name="status"
						checked={isPublic === "public"}
						onChange={(e) => setIsPublic(e.target.value)}
					/>
					<label className="form-check-label" htmlFor="public-status">
						Public
					</label>
				</div>

				<div className="form-check">
					<input
						className="form-check-input"
						type="radio"
						id="private-status"
						value="private"
						name="status"
						checked={isPublic === "private"}
						onChange={(e) => setIsPublic(e.target.value)}
					/>
					<label className="form-check-label" htmlFor="private-status">
						Private
					</label>
				</div>


				<div className="mb-3">
					<label htmlFor="about">About me</label>
					<input
						type="text"
						className="form-control reginput"
						id="bio"
						placeholder="About Me"
						cols="30"
						rows="10"
						onChange={(e) => setBio(e.target.value)} />
				</div>



				<button className="btn btn-primary" type="submit">Register</button>

			</form>




			<div className="error-message"></div>
			<br /> {/* Add a line break for spacing */}
			<div className="mb3">
				<span className="login-text">Already have an account? &nbsp;</span>
				<button type="submit" className="btn btn-primary" onClick={renderLogin}>
					Log in
				</button>
			</div>
		</div>
	);
}

function Profile() {

	return (
		<div>
			<Navbar />
			<h1>Profile</h1>
		</div>
	)
}


function Chat() {
	return (
		<div>
			<Navbar />
			<h1>Chat</h1>
		</div>
	)
}

function Group() {
	return (
		<div>
			<Navbar />
			<h1>Group</h1>
		</div>

	)
}

function Notifications() {
	return (
		<div>
			<Navbar />
			<h1>Notifications</h1>
		</div>
	)
}



// Main post form, defaults to sending posts to public group (0)
function PostForm({ groupId }) {
	const [body, setBody] = useState("");
	const [privacy, setPrivacy] = useState("");
	const [selectedFile, setSelectedFile] = useState(null);

	// Upon submitting:
	const submit = async (e) => {
		e.preventDefault(); // prevent reload.

		const formData = new FormData();

		// Append form data
		formData.append('body', body);
		formData.append('privacy', privacy);
		if (privacy === "private") {
			groupId = -1
		}
		formData.append('groupId', groupId);
		if (selectedFile) {
			formData.append('image', selectedFile);
		}

		console.log("Form data being sent to backend: ", formData);

		// Send user data to golang api/PostHandler.go.
		await fetch("http://localhost:8080/api/posts", {
			method: "POST",
			credentials: "include",
			body: formData,
		});

		// Reset the form fields to their default state
		setBody("");
		setPrivacy("");
		setSelectedFile(null);

		document.getElementById('postFormBody').value = "";
	};


	// Function to handle file selection
	const handleFileChange = (e) => {
		setSelectedFile(e.target.files[0]);
		// const file = e.target.files[0];
	};

	const handleSelectFile = () => {
		const fileInput = document.getElementById('fileInput');
		fileInput.click();
	};

	return (
		<div>
			<main className="postForm container" style={{ maxWidth: "400px" }}>
				<h1 className="h3 mb-3 fw-normal">Post Message Here</h1>
				<form onSubmit={submit}>
					<div className="form-floating mb-3">
						<input
							type="text"
							className="form-control"
							id="postFormBody"
							placeholder="Type your post here..."
							onChange={(e) => setBody(e.target.value)}
						/>
					</div>

					<div>
						<button type="button" className="btn btn-primary" onClick={handleSelectFile}>
							Select File
						</button>
						<span>{selectedFile ? selectedFile.name : 'No file selected'}</span>
						<input
							type="file"
							id="fileInput"
							accept="image/*"
							style={{ display: 'none' }}
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
								onClick={(e) => setPrivacy(e.target.value)}
								className="form-check-input"
							/>
							<label htmlFor="post-public-status" className="form-check-label">Public</label>
						</div>
						<div className="form-check">
							<input
								required
								type="radio"
								id="private-status"
								value="private"
								name="status"
								checked={privacy === "private"}
								onClick={(e) => setPrivacy(e.target.value)}
								className="form-check-input"
							/>
							<label htmlFor="private-status" className="form-check-label">Private</label>
						</div>
					</div>

					<button className="w-100 btn btn-lg btn-primary" type="submit">
						Submit
					</button>
				</form>
			</main>
		</div>
	);
}

const postCardStyle = {
	maxWidth: '600px',
	background: 'linear-gradient(to bottom, #c7ddef, #ffffff)', // Light blue/grey to white gradient
	borderRadius: '10px',
	boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)', // Optional: Add shadow for depth
	padding: '20px',
	margin: 'auto',
	marginBottom: '20px', // Adjust spacing between post cards
};

function PostCard({ post }) {
	const [body, setBody] = useState("");
	const [selectedFile, setSelectedFile] = useState(null);

	const milliseconds = post.post.createdAt;
	const date = new Date(milliseconds);
	const formattedDate = date.toLocaleString();

	const submit = async (e) => {
		e.preventDefault(); // prevent reload.

		const formData = new FormData();

		// Append form data
		formData.append('body', body);
		formData.append('postId', post.post.postId);
		if (selectedFile) {
			formData.append('image', selectedFile);
		}

		console.log("Form data being sent to backend: ", formData);

		// Send user data to golang api/PostHandler.go.
		await fetch("http://localhost:8080/api/comments", {
			method: "POST",
			credentials: "include",
			body: formData,
		});

		// Reset the form fields to their default state
		setBody("");
		setSelectedFile(null);


		document.getElementById('commentTextArea').value = "";
	};

	// Function to handle file selection
	const handleFileChange = (e) => {
		setSelectedFile(e.target.files[0]);
		// const file = e.target.files[0];
	};

	const handleSelectFile = () => {
		const commentFileInput = document.getElementById(`commentFileInput${post.post.postId}`);
		commentFileInput.click();
	};


	return (
		<div className="card" style={postCardStyle}>
			<div className="card-body">
				<div className="d-flex flex-start align-items-center">
					<img className="rounded-circle shadow-1-strong me-3"
						src={post.post.imageURL} alt="avatar" width="60" height="60" />
					<div>
						<h6 className="fw-bold text-primary mb-1">{post.post.userId}</h6>
						{/* Date, formatted */}
						<p className="text-muted small mb-0">
							{formattedDate}
						</p>
					</div>
				</div>
				{/* Image, if there is one */}
				{!post.post.imageURL ? null : (
					<p className="mt-3 mb-2 pb-1">
						<img src={post.post.imageURL} className="img-fluid" />
					</p>
				)}
				{/* Post Body */}
				<p className="mt-3 mb-2 pb-1">
					{post.post.body}
				</p>
			</div>
			<div className="card-footer py-3 border-0" style={{ backgroundColor: '#f8f9fa' }}>
				<div className="d-flex flex-start w-100">
					<img className="rounded-circle shadow-1-strong me-3"
						src={post.avatar} alt="avatar" width="40" height="40" />
					<div className="form-outline w-100">
						<textarea className="form-control" id="commentTextArea" rows="4"
							style={{ background: '#fff' }} onChange={(e) => setBody(e.target.value)}></textarea>

						<label className="form-label" htmlFor="textAreaExample">Message</label>
					</div>
				</div>
				<div className="float-end mt-2 pt-1">
					<button type="button" className="btn btn-primary" onClick={handleSelectFile}>Select File</button>
					<span>{selectedFile ? selectedFile.name : 'No file selected'}</span>
					<input
						type="file"
						id={`commentFileInput${post.post.postId}`}
						accept="image/*"
						style={{ display: 'none' }}
						onChange={handleFileChange}
					/>
					<button type="submit" className="btn btn-primary btn-sm" onClick={submit}>Post comment</button>
				</div>
				<div className="comments">
					<h2>Comments</h2>
					{post.comments !== null && post.comments.length > 0 ? (
						post.comments.map(comment => (
							<CommentCard key={comment.createdAt} comment={comment} />
						))
					) : (
						<p className="text-muted">No comments</p>
					)}
				</div>
			</div>
		</div>
	);
}

function CommentCard({ comment }) {
	const formattedDate = new Date(comment.createdAt).toLocaleString();

	return (
		<div className="card mt-3">
			<div className="d-flex flex-start align-items-center">
				<img className="rounded-circle shadow-1-strong me-3"
					src={comment.imageURL} alt="avatar" width="60" height="60" />
				<div>
					<h6 className="fw-bold text-primary mb-1">{comment.userId}</h6>
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


// Display information relating to homepage
function Home() {
	const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
	const [privatePosts, setPrivatePosts] = useState([]);
	const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
	const [userGroups, setUserGroups] = useState([]);

	useEffect(() => {
		fetch('http://localhost:8080/api/home')
			.then(response => response.json())
			.then(data => {
				setAlmostPrivatePosts(data.almostPrivatePosts)
				setPrivatePosts(data.privatePosts)
				setPublicPostsWithComments(data.publicPostsWithComments)
				setUserGroups(data.userGroups)
			})
			.catch(error => {
				console.error('Error fetching data:', error);
			});
	}, []);

	return (
		<main className="homePage">
			<Navbar />
			<PostForm groupId={0} />
			{/* Rendering Almost Private Posts */}
			<div className="almostPrivatePosts">
				<h2>Almost Private Posts</h2>
				{almostPrivatePosts !== null && almostPrivatePosts.length > 0 ? (
					almostPrivatePosts.map(almostPrivatePost => (
						<PostCard key={almostPrivatePost.createdAt} post={almostPrivatePost} />
					))
				) : (
					<p>No almost private posts</p>
				)}
			</div>


			{/* Rendering Private Posts */}
			<div className="privatePosts">
				<h2>Almost Private Posts</h2>
				{privatePosts !== null && privatePosts.length > 0 ? (
					privatePosts.map(privatePost => (
						<PostCard key={privatePost.createdAt} post={privatePost} />
					))
				) : (
					<p>No private posts</p>
				)}
			</div>


			{/* Rendering Public Posts */}
			<div className="publicPostsWithComments">
				<h2>Public Posts With Comments</h2>
				{publicPostsWithComments !== null && publicPostsWithComments.length > 0 ? (
					publicPostsWithComments.map((publicPostsWithComment, index) => (
						<PostCard key={index} post={publicPostsWithComment} />
					))
				) : (
					<p>public posts</p>
				)}
			</div>



			{/* Rendering User Groups */}
			<div className="userGroups">
				<h2>Groups</h2>
				<ul>
					{userGroups !== null && userGroups.map(userGroup => (
						<li key={userGroup.createdAt}>
							{userGroup.Title} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>
		</main>
	);
}

const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
