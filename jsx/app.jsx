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

function Navbar(props) {

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

	const renderLogin = () => {

		//Some logout logic and function needs implemented here rather than lines below

		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Login />, appContainer);
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
							<a className="nav-link" href="#" onClick={renderLogin}>LOGOUT</a>
						</li>
					</ul>
				</div>
			</div>
		</nav>
	)
}

function Login(props) {
	const [usernameOrEmail, setUsernameOrEmail] = useState("")
	const [password, setPassword] = useState("")
	const [redirectVar, setRedirectVar] = useState(false)
	const [error, setError] = useState(null);
	const [isLoggedIn, setIsLoggedIn] = useState(false);
	const errorMessage = document.querySelector(".error-message")
	const [showForm, setShowForm] = useState(true);

	//this is the sign in button
	const submit = async (e) => {
		e.preventDefault() // prevent reload.

		//this is user input 
		const userToLogin = {
			usernameOrEmail,
			password,
		}

		try {
			//check credentials with backend
			const response = await fetch('http://localhost:8080/auth/login', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify(userToLogin),
			});

			if (!response.ok) {
				errorMessage.innerHTML = 'Invalid credentials'
				throw new Error('Invalid credentials');
			}

			//takes response from backend and processes
			const data = await response.json();
			if (data.success) {
				setIsLoggedIn(true);
			} else {
				errorMessage.innerHTML = 'Invalid credentials'
				throw new Error('Invalid credentials');
			}
		} catch (error) {
			errorMessage.innerHTML = 'Invalid credentials'
			setError('Invalid credentials');
		}
	};

	//if credentials frontend match backend then we render home
	if (isLoggedIn) {
		const appContainer = document.querySelector('.app-container');
		//ReactDOM.render(<Home />, appContainer);
		ReactDOM.render(<Home />, appContainer);
	}

	//this is the register button, when pressed will serve registration form
	const renderRegister = () => {
		const appContainer = document.querySelector('.app-container');
		ReactDOM.render(<Register />, appContainer);
	};


	return (
		<div className="container login-container">
			<h1 className="h3 mb-3 fw-normal login-text">log in</h1>
			<form onSubmit={submit}>
				<div class="mb-3">
					<label for="exampleInputEmail1" class="form-label">Email address</label>
					<input
						type="email"
						className="form-control form-control-lg"
						id="exampleInputEmail1"
						aria-describedby="emailHelp"
						onChange={(e) => setUsernameOrEmail(e.target.value)} />
				</div>
				<div class="mb-3">
					<label for="exampleInputPassword1" class="form-label">Password</label>
					<input
						type="password"
						className="form-control form-control-lg"
						id="exampleInputPassword1"
						onChange={(e) => setPassword(e.target.value)} />
				</div>
				<button type="submit" class="btn btn-primary">Log in</button>
			</form>
			<div className="error-message"></div>
			<br /> {/* Add a line break for spacing */}
			<div className="mb3">
				<span className="login-text">Don't have an account? &nbsp;</span>
				<button type="submit" className="btn btn-primary" onClick={renderRegister}>
					Register
				</button>
			</div>
		</div>
	)
}

function Register(props) {
	const [email, setEmail] = useState("");
	const [encryptedPassword, setEncryptedPassword] = useState("");
	const [firstName, setFirstName] = useState("");
	const [lastName, setLastName] = useState("");
	const [dob, setDob] = useState("");
	const [imageURL, setImageURL] = useState("");
	const [username, setUsername] = useState("");
	const [bio, setBio] = useState("");
	const [isPublic, setIsPublic] = useState("public");
	const [redirectVar, setRedirectVar] = useState(false);
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
			setError('Invalid credentials');
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
			<form>
  <div className="mb-3">
    <label htmlFor="exampleInputEmail1" className="form-label">Title</label>
    <input type="email" className="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" />
    <div id="emailHelp" className="form-text">We'll never share your email with anyone else.</div>
  </div>
  <div className="mb-3">
    <label htmlFor="exampleInputPassword1" className="form-label">Description</label>
    <input type="password" className="form-control" id="exampleInputPassword1" />
  </div>
  <div className="mb-3 form-check">
    <input type="checkbox" className="form-check-input" id="exampleCheck1" />
    <label className="form-check-label" for="exampleCheck1">Check me out</label>
  </div>
  <button type="submit" className="btn btn-primary">Submit</button>
</form>
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

	// Needs to be changed to get info from... cookie?
	const userId = Number(36);

	// Upon submitting:
	const submit = async (e) => {
		e.preventDefault(); // prevent reload.
	
		const formData = new FormData();
	
		// Append form data
		formData.append('body', body);
		formData.append('privacy', privacy);
		formData.append('groupId', groupId);
		formData.append('userId', userId);
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


			<div className="almostPrivatePosts">
				<h2>Almost Private Posts</h2>
				<ul>
					{almostPrivatePosts !== null && almostPrivatePosts.map(almostPrivatePost => (
						<li key={almostPrivatePost.createdAt}>
							{almostPrivatePost.body} - {almostPrivatePost.UserId}
							{/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="privatePosts">
				<h2>Private Posts</h2>
				<ul>
					{privatePosts !== null && privatePosts.map(privatePost => (
						<li key={privatePost.createdAt}>
							{privatePost.body} - {privatePost.UserId} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

			<div className="publicPostsWithComments">
				<h2>Public Posts</h2>
				<ul>
					{publicPostsWithComments !== null && publicPostsWithComments.map(publicPostsWithComment => (
						<li key={publicPostsWithComment.post.CreatedAt}>
							{publicPostsWithComment.post.Body} - {publicPostsWithComment.post.UserId} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

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

// function Profile() {

// 	const [usernameOrEmail, setUsernameOrEmail] = useState("")
// 	const [password, setPassword] = useState("")

// 	//this is the sign in button
// 	const submit = async (e) => {
// 		e.preventDefault() // prevent reload.

// 		//this is user input 
// 		const userToLogin = {
// 			usernameOrEmail,
// 			password,
// 		}

// 		try {
// 			//check credentials with backend
// 			const response = await fetch('http://localhost:8080/auth/login', {
// 				method: 'POST',
// 				headers: { 'Content-Type': 'application/json' },
// 				credentials: 'include',
// 				body: JSON.stringify(userToLogin),
// 			});

// 			if (!response.ok) {
// 				errorMessage.innerHTML = 'Invalid credentials'
// 				throw new Error('Invalid credentials');
// 			}

// 			//takes response from backend and processes
// 			const data = await response.json();
// 			if (data.success) {
// 				setIsLoggedIn(true);
// 			} else {
// 				errorMessage.innerHTML = 'Invalid credentials'
// 				throw new Error('Invalid credentials');
// 			}
// 		} catch (error) {
// 			errorMessage.innerHTML = 'Invalid credentials'
// 			setError('Invalid credentials');
// 		}
// 	};

// 	//if credentials frontend match backend then we render home
// 	if (isLoggedIn) {
// 		const appContainer = document.querySelector('.app-container');
// 		ReactDOM.render(<Home />, appContainer);
// 	}

// 	//this is the register button, when pressed will serve registration form
// 	const renderRegister = () => {
// 		const appContainer = document.querySelector('.app-container');
// 		ReactDOM.render(<Register />, appContainer);
// 	};


// 	return (
// 		<div className="login-container">
// 			<main className="form-signin w-100 m-auto" style={{ display: "block" }}>
// 				<h1 className="h3 mb-3 fw-normal login-text">log in</h1>
// 				<form onSubmit={submit}>
// 					<div className="form-floating">
// 						<label htmlFor="floatingInput" className="login-text">Email address</label>
// 						<input
// 							type="email"
// 							className="form-control login-text"
// 							id="floatingInput"
// 							placeholder="name@example.com"
// 							onChange={(e) => setUsernameOrEmail(e.target.value)}
// 						/>
// 					</div>

// 					<div className="form-floating">
// 						<label htmlFor="floatingPassword" className="login-text">Password</label>
// 						<input
// 							type="password"
// 							className="form-control login-text"
// 							id="floatingPassword"
// 							placeholder="Password"
// 							onChange={(e) => setPassword(e.target.value)}
// 						/>
// 					</div>
// 					<button className="w-100 btn btn-lg btn-primary login-button" type="submit">
// 						Sign in
// 					</button>
// 				</form>
// 				<div className="error-message"></div>
// 				<br /> {/* Add a line break for spacing */}
// 				<span className="login-text">Don't have an account? &nbsp;</span>
// 				<button className="w-100 btn btn-lg btn-primary login-button" onClick={renderRegister}>
// 					Register
// 				</button>
// 			</main>
// 		</div>
// 	)
// }

const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
