const { useState, useEffect } = React

const App = () => {
	return (
		<div className="app-container">
			<Login />
			<Register />
		</div>
	)
}

function Login(props) {
	const [usernameOrEmail, setUsernameOrEmail] = useState("")
	const [password, setPassword] = useState("")
	// const [redirectVar, setRedirectVar] = useState(false)

	const submit = async (e) => {
		e.preventDefault() // prevent reload.

		const userToLogin = {
			usernameOrEmail,
			password,
		}
		console.log(userToLogin)

		// Send user data to golang register function.
		await fetch("http://localhost:8080/auth/login", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			credentials: "include",
			body: JSON.stringify(userToLogin),
		})

		// setRedirectVar(true)
		// props.setName(validUser.first)
		const appContainer = document.querySelector('.app-container');

		ReactDOM.render(completedHomePage(), appContainer);
		// const validUser = await response.json()

	}


	return (
		<div>
			<main className="form-signin w-100 m-auto" style={{ display: "block" }}>
				<h1 className="h3 mb-3 fw-normal">Please sign in</h1>
				<form onSubmit={submit}>
					<div className="form-floating">
						<input
							type="email"
							className="form-control"
							id="floatingInput"
							placeholder="name@example.com"
							onChange={(e) => setUsernameOrEmail(e.target.value)}
						/>
						<label htmlFor="floatingInput">Email address</label>
					</div>
					<div className="form-floating">
						<input
							type="password"
							className="form-control"
							id="floatingPassword"
							placeholder="Password"
							onChange={(e) => setPassword(e.target.value)}
						/>
						<label htmlFor="floatingPassword">Password</label>
					</div>
					<button className="w-100 btn btn-lg btn-primary" type="submit">
						Sign in
					</button>
				</form>
				<span>Already have an account? &nbsp;</span>
				{/* <Link to="/register" style={{ color: "white" }}>
					Register
				</Link> */}
			</main>
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
	const [isPublic, setIsPublic] = useState("");
	const [redirectVar, setRedirectVar] = useState(false);

	// Redirect
	//const navigate = useNavigate();

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
		// Send user data to golang register function.
		const response = await fetch("http://localhost:8080/auth/registration", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify(newUser),
		});
		console.log("dob", newUser.dob)
		await response.json()
		// let result = await response.json()
		// if (result.email === email) {
		setRedirectVar(true);
		// }
	};

	// if (redirectVar) {
	// 	return navigate("/login"); // This is still iffy!!! ????????????
	// }

	return (
		<div>
			<main className="form-signin w-100 m-auto" style={{ display: "block" }}>
				<h1 className="h3 mb-3 fw-normal">Please register</h1>
				<form onSubmit={submit}>

					<div className="form-floating">
						<input
							required
							type="email"
							className="form-control"
							id="floatingInput"
							placeholder="name@example.com"
							onChange={(e) => setEmail(e.target.value)}
						/>
						<label htmlFor="floatingInput">Email address</label>
					</div>
					<div className="form-floating">
						<input
							required
							type="password"
							className="form-control reginput"
							id="regpassword"
							placeholder="Password"
							onChange={(e) => setEncryptedPassword(e.target.value)}
						/>
						<label htmlFor="regpassword">Password</label>
					</div>
					<div className="form-floating">
						<input
							required
							type="text"
							className="form-control reginput"
							id="firstName"
							placeholder="John"
							onChange={(e) => setFirstName(e.target.value)}
						/>
						<label htmlFor="firstName">First Name</label>
					</div>
					<div className="form-floating">
						<input
							required
							type="text"
							className="form-control reginput"
							id="lastName"
							placeholder="Doe"
							onChange={(e) => setLastName(e.target.value)}
						/>
						<label htmlFor="lastName">Last Name</label>
					</div>
					<div className="form-floating">
						<input
							required
							type="date"
							className="form-control reginput"
							id="dob"
							placeholder="16/01/1998"
							onChange={(e) => setDob(e.target.value)}
						/>
						<label htmlFor="dob">Date of Birth</label>
					</div>

					<div className="form-floating">
						<input
							type="text"
							className="form-control reginput"
							id="imageURL"
							placeholder="https://..."
							onChange={(e) => setImageURL(e.target.value)}
						/>
						<label htmlFor="imageURL">ImageURL</label>
					</div>
					<div className="form-floating">
						<input
							type="text"
							className="form-control reginput"
							id="username"
							placeholder="Johnny"
							onChange={(e) => setUsername(e.target.value)}
						/>
						<label htmlFor="username">Username</label>
					</div>
					<div className="form-floating">
						<div className="form-control reginput status">
							<div>
								<input
									required
									type="radio"
									id="public-status"
									value="public"
									name="status"
									checked
									onClick={(e) => setIsPublic(e.target.value)}
								/>
								<label htmlFor="public-status">Public</label>
							</div>
							<div>
								<input
									required
									type="radio"
									id="private-status"
									value="private"
									name="status"
									onClick={(e) => setIsPublic(e.target.value)}
								/>
								<label htmlFor="private-status">Private</label>
							</div>
						</div>
						<label htmlFor="">Status</label>
					</div>
					<div className="form-floating">
						<input
							className="form-control reginput"
							name="bio"
							placeholder="About Me"
							id="bio"
							cols="30"
							rows="10"
							onChange={(e) => setBio(e.target.value)}
						></input>
						<label htmlFor="about">About me</label>
					</div>
					<button className="w-100 btn btn-lg btn-primary" type="submit">
						Register
					</button>
				</form>
				<span>Already have an account? &nbsp;</span>
				{/* <Link to="/login" style={{ color: "white" }}>
          Login
        </Link> */}
			</main>
		</div>
	);
}

// Main post form, defaults to sending posts to public group
function PostForm(props) {
	const [body, setBody] = useState("")
	const [imageURL, setImageURL] = useState("")

	// const [redirectVar, setRedirectVar] = useState(false)
	const groupId = Number(0)
	const userId = Number(36)
	const privacy = "public"

	const submit = async (e) => {
		e.preventDefault() // prevent reload.

		const postToSend = {
			body,
			groupId,
			imageURL,
			privacy,
			userId
		}
		console.log("Post being sent to backend: ", postToSend)

		// Send user data to golang register function.
		await fetch("http://localhost:8080/api/posts", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			credentials: "include",
			body: JSON.stringify(postToSend),
		})


	}

	return (
		<div>
			<main className="postForm" style={{ display: "block" }}>
				<h1 className="h3 mb-3 fw-normal">Post Message Here</h1>
				<form onSubmit={submit}>
					<div className="form-floating">
						<input
							type="text"
							className="form-control"
							id="postFormBody"
							placeholder="name@example.com"
							onChange={(e) => setBody(e.target.value)}
						/>
					</div>
					<div className="form-floating">
						<input
							type="text"
							className="form-control"
							id="postFormImgURL"
							placeholder="This/Will/Need/A/Button.gif"
							onChange={(e) => setImageURL(e.target.value)}
						/>
						<label htmlFor="postFormImgURL">Image URL</label>
					</div>
					<button className="w-100 btn btn-lg btn-primary" type="submit">
						Submit
					</button>
				</form>
			</main>
		</div>
	)
}

function Home() {
	const [users, setUsers] = useState([]);
	const [almostPrivatePosts, setAlmostPrivatePosts] = useState([]);
	const [privatePosts, setPrivatePosts] = useState([]);
	const [publicPostsWithComments, setPublicPostsWithComments] = useState([]);
	const [userEvents, setUserEvents] = useState([]);
	const [userGroups, setUserGroups] = useState([]);
	const [userNotifications, setUserNotifications] = useState([]);

	useEffect(() => {
		fetch('http://localhost:8080/api/home')
			.then(response => response.json())
			.then(data => {
				setUsers(data.allUsers);
				setAlmostPrivatePosts(data.almostPrivatePosts)
				setPrivatePosts(data.privatePosts)
				setPublicPostsWithComments(data.publicPostsWithComments)
				setUserEvents(data.userEvents)
				setUserGroups(data.userGroups)
				setUserNotifications(data.userNotifications)
			})
			.catch(error => {
				console.error('Error fetching data:', error);
			});
	}, []);

	return (
		<div className="homePage">

			<div className="allUsersList">
				<h2>All Users</h2>
				<ul>
					{users.map(user => (
						<li key={user.userId}>
							{user.username} - {user.email} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>

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

			<div className="userEvents">
				<h2>Events</h2>
				<ul>
					{userEvents !== null && userEvents.map(userEvent => (
						<li key={userEvent.createdAt}>
							{userEvent.Title} {/* Render whatever user properties you need */}
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

			<div className="userNotifications">
				<h2>Notifications</h2>
				<ul>
					{userNotifications !== null && userNotifications.map(userNotification => (
						<li key={userNotification.createdAt}>
							{userNotification.NotificationType} {/* Render whatever user properties you need */}
						</li>
					))}
				</ul>
			</div>
		</div>
	);
}

const completedHomePage = () => {
	return (
		<div className="completedHomePage">
			<PostForm />
			<Home />
		</div>
	)
}


const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
