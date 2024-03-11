const { useState } = React

const App = () => {
	return (
		<div className="app-container">
			<Login />
			<Register />
			<Home />
			<Profile />
		</div>
	)
}

function Login(props) {
	const [usernameOrEmail, setUsernameOrEmail] = useState("")
	const [password, setPassword] = useState("")
	const [redirectVar, setRedirectVar] = useState(false)

	const submit = async (e) => {
		e.preventDefault() // prevent reload.

		const userToLogin = {
			usernameOrEmail,
			password,
		}
		console.log(userToLogin)

		// Send user data to golang register function.
		const response = await fetch("http://localhost:8080/auth/login", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			credentials: "include",
			body: JSON.stringify(userToLogin),
		})

		const validUser = await response.json()
		setRedirectVar(true)
		props.setName(validUser.first)
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

function Home(props) {
	return (
		<main>
			<div className="contentContainer">
				{props.name ? (
					<>
						<ProfileImgContainer
							name={props.name}
							user={props.user}
							imageURL={props.imageURL}
						/>
						<GroupContainer groups={props.groups} socket={props.socket} />
						<PostForm imageURL={props.imageURL} />
						<RightSide openConnection={props.openConnection} fetchRequestData={props.fetchRequestData} />
						<GetChat />
					</>
				) : (
					<>
						<p>You are not logged in</p>
						{/* <Link to="/login">Login</Link> */}
					</>
				)}
			</div>
		</main>
	);
}

function Profile(props) {
	const [status, setStatus] = useState("");
	const [privatePosts, setPrivatePosts] = useState([]);


	// Update status to props.user.status.
	// useEffect(() => {
	//   setStatus(props.user.status);
	// }, [props.user.status]);

	const sendStatusToBackend = async (data) => {
		console.log(data);
		await fetch("http://localhost:8080/update-user-status", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			credentials: "include",
			body: JSON.stringify(data),
		});
	};

	const updateUserStatus = async (ev) => {
		let buttonClicked = ev.target.getAttribute("data-type");
		if (buttonClicked === "private") {
			sendStatusToBackend({
				user: props.user.email,
				setStatus: "private",
			});
			setStatus("private");
		} else if (buttonClicked === "public") {
			// update on backend if user is not already public
			sendStatusToBackend({
				user: props.user.email,
				setStatus: "public",
			});
			setStatus("public");
		}
	};

	return (
		<div className="profileContainer">

			name={props.name}
			user={props.user}
			imageURL={props.imageURL}
			socket={props.socket}
			currentUser={props.currentUser}
			fetchUsersData={props.fetchUsersData}
			update={props.update}
			setUpdate={props.setUpdate}

			<div className="formContainer">
				<div className="smallAvatar">
					<img src={props.imageURL} alt="profile photo" />
				</div>
				<div className="profile-page-title">{props.name}'s Posts</div>
			</div>

			{/* If my profile */}

			{props.currentUser === undefined ? (
				<div
					id="set-public-private"
					className="privacyButtons"
					style={{
						width: "100%",
						backgroundColor: "white",
						justifyContent: "space-evenly",
						alignItems: "center",
					}}
				>
					{/* currentUser is not passed to profile when redirecting to myProfile */}
					<>
						<button
							className="postType"
							onClick={updateUserStatus}
							data-type="private"
							disabled={status === "private" ? true : false}
							style={{
								backgroundColor:
									status === "private"
										? "rgba(129, 25, 41, 0.55)"
										: "rgb(148, 28, 47)",
							}}
						>
							Set Private
						</button>
						<button
							className="postType"
							onClick={updateUserStatus}
							data-type="public"
							disabled={status === "public" ? true : false}
							style={{
								backgroundColor:
									status === "public"
										? "rgba(129, 25, 41, 0.55)"
										: "rgb(148, 28, 47)",
							}}
						>
							Set Public
						</button>
					</>
				</div>
			) : (
				<div
					id="set-public-private"
					className="privacyButtons"
					style={{ width: "100%", backgroundColor: "rgba(250, 250, 250, 0.5)" }}
				></div>
			)}

			{/* <AllPosts user={props.user} privatePosts={privatePosts} />
		<RightSide openConnection={props.openConnection} fetchRequestData={props.fetchRequestData}  /> */}
		</div>
	);
}


const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
