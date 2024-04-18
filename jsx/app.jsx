const { useState, useEffect } = React

let socket

const App = () => {
	return (
		<div className="app-container">
			<div className="nav-container"></div>
			<div className="page-container">
				<Login />
			</div>
		</div>
	)
}

const getCurrentUserId = () => {
	const [currentUserId, setCurrentUserId] = useState(null)
	const [isLoading, setIsLoading] = useState(true)
	const [error, setError] = useState(null)

	useEffect(() => {
		const fetchUserId = async () => {
			try {
				const response = await fetch("http://localhost:8080/api/userId", {
					credentials: "include",
				})

				if (response.ok) {
					const userId = await response.json()
					setCurrentUserId(userId)
				} else {
					setError("Failed to fetch userId")
				}
			} catch (error) {
				setError("Error fetching userId")
			} finally {
				setIsLoading(false)
			}
		}

		fetchUserId()
	}, [])

	return { currentUserId, isLoading, error }
}

const renderNavbar = () => {
	const navContainer = document.querySelector(".nav-container")
	ReactDOM.render(<Navbar />, navContainer)
}

function Navbar() {
	const { currentUserId, isLoading, error } = getCurrentUserId()

	const logout = async () => {
		try {
			const response = await fetch("http://localhost:8080/auth/logout", {
				method: "POST",
				credentials: "include",
			})

			console.log(response)

			if (response.ok) {
				socket.close()
				socket.addEventListener("close", (event) => {
					console.log("The connection has been closed successfully.")
				})
				renderLogin()
				const navContainer = document.querySelector(".nav-container")
				ReactDOM.render(null, navContainer)
				console.log("Logout successful!")
			} else {
				console.log("Failed to logout. Server response not OK.")
			}
		} catch (error) {
			console.error("An error occurred during logout:", error)
		}
	}

	return (
		<nav className="navbar navbar-expand-md bg-body-tertiary">
			<div className="container-fluid">
				<button
					className="navbar-toggler"
					type="button"
					data-bs-toggle="collapse"
					data-bs-target="#navbarSupportedContent"
					aria-controls="navbarSupportedContent"
					aria-expanded="false"
					aria-label="Toggle navigation"
				>
					<span className="navbar-toggler-icon"></span>
				</button>
				<div className="collapse navbar-collapse" id="navbarSupportedContent">
					<ul className="navbar-nav me-auto mx-auto mb-2 mb-lg-0">
						<li className="nav-item">
							<a
								className="nav-link"
								href="#"
								onClick={() => renderProfile(currentUserId, true)}
							>
								PROFILE
							</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderHome}>
								HOME
							</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderNotifications}>
								NOTIFICATIONS
							</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderChat}>
								CHAT
							</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={renderGroup}>
								GROUP
							</a>
						</li>
						<li className="nav-item">
							<a className="nav-link" href="#" onClick={logout}>
								LOGOUT
							</a>
						</li>
					</ul>
				</div>
			</div>
		</nav>
	)
}

const renderLogin = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Login />, pageContainer)
}

function Login() {
	const [usernameOrEmail, setUsernameOrEmail] = useState("")
	const [password, setPassword] = useState("")
	const [isLoggedIn, setIsLoggedIn] = useState(false)
	const [errorMessage, setErrorMessage] = useState("")

	const handleUsernameOrEmailChange = (e) => {
		setUsernameOrEmail(e.target.value)
	}

	const handlePasswordChange = (e) => {
		setPassword(e.target.value)
	}

	const handleSubmit = async (e) => {
		e.preventDefault()

		const userToLogin = {
			usernameOrEmail,
			password,
		}

		try {
			const response = await fetch("http://localhost:8080/auth/login", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				credentials: "include",
				body: JSON.stringify(userToLogin),
			})

			if (!response.ok) {
				setErrorMessage("Invalid credentials")
				throw new Error("Invalid credentials")
			}

			const data = await response.json()
			if (data.success) {
				setIsLoggedIn(true)
				setErrorMessage("")
			} else {
				setErrorMessage("Invalid credentials")
				throw new Error("Invalid credentials")
			}
		} catch (error) {
			setErrorMessage("Invalid credentials")
		}
	}

	useEffect(() => {
		if (isLoggedIn) {
			renderNavbar()
			renderHome()

			socket = new WebSocket("ws://localhost:8080/ws")
			socket.onopen = function (event) {
				console.log("WebSocket connection established.")
			}
		}
	}, [isLoggedIn])

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
				<button
					type="button"
					className="btn btn-primary"
					onClick={renderRegister}
				>
					Register
				</button>
			</div>
		</div>
	)
}

const renderRegister = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Register />, pageContainer)
}

function Register() {
	const [email, setEmail] = useState("")
	const [encryptedPassword, setEncryptedPassword] = useState("")
	const [firstName, setFirstName] = useState("")
	const [lastName, setLastName] = useState("")
	const [dob, setDob] = useState("")
	const [imageURL, setImageURL] = useState("")
	const [username, setUsername] = useState("")
	const [bio, setBio] = useState("")
	const [isPublic, setIsPublic] = useState(true)
	const [isRegistered, setIsRegistered] = useState(false)

	const handleChange = (e) => {
		setIsPublic(e.target.value === "true")
	}
	//this is register button
	const submit = async (e) => {
		e.preventDefault() // prevent reload.

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
		}

		try {
			// Send user data to backend
			const response = await fetch("http://localhost:8080/auth/registration", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(newUser),
			})

			if (!response.ok) {
				throw new Error("Invalid credentials")
			}

			//takes response from backend and processes
			const data = await response.json()
			if (data.success) {
				setIsRegistered(true)
			} else {
				throw new Error("Invalid credentials")
			}
		} catch (error) {
			throw new Error("Invalid credentials")
		}
	}

	//if credentials frontend succesfully create a new user then we render home
	if (isRegistered) {
		socket = new WebSocket("ws://localhost:8080/ws")
		socket.onopen = function (event) {
			console.log("WebSocket connection established.")
		}
		renderNavbar()
		renderHome()
	}

	//this is the login button, when pressed will serve login form

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
						onChange={(e) => setEmail(e.target.value)}
					/>
				</div>

				<div className="mb-3">
					<label htmlFor="regpassword">Password</label>
					<input
						required
						type="password"
						className="form-control reginput"
						id="regpassword"
						placeholder="Password"
						onChange={(e) => setEncryptedPassword(e.target.value)}
					/>
				</div>

				<div className="mb-3">
					<label htmlFor="firstName">First Name</label>
					<input
						required
						type="text"
						className="form-control reginput"
						id="firstName"
						placeholder="John"
						onChange={(e) => setFirstName(e.target.value)}
					/>
				</div>

				<div className="mb-3">
					<label htmlFor="lastName">Last Name</label>
					<input
						required
						type="text"
						className="form-control reginput"
						id="lastName"
						placeholder="Doe"
						onChange={(e) => setLastName(e.target.value)}
					/>
				</div>

				<div className="mb-3">
					<label htmlFor="dob">Date of Birth</label>
					<input
						required
						type="date"
						className="form-control reginput"
						id="dob"
						placeholder="16/01/1998"
						onChange={(e) => setDob(e.target.value)}
					/>
				</div>

				<div className="mb-3">
					<label htmlFor="imageURL">ImageURL</label>
					<input
						type="text"
						className="form-control reginput"
						id="imageURL"
						placeholder="https://..."
						onChange={(e) => setImageURL(e.target.value)}
					/>
				</div>

				<div className="mb-3">
					<label htmlFor="username">Username</label>
					<input
						type="text"
						className="form-control reginput"
						id="username"
						placeholder="Johnny"
						onChange={(e) => setUsername(e.target.value)}
					/>
				</div>

				<div className="form-check">
					<input
						className="form-check-input"
						type="radio"
						id="public-status"
						value={true}
						name="status"
						checked={isPublic === true}
						onChange={handleChange}
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
						value={false}
						name="status"
						checked={isPublic === false}
						onChange={handleChange}
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
						onChange={(e) => setBio(e.target.value)}
					/>
				</div>

				<button className="btn btn-primary" type="submit">
					Register
				</button>
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
	)
}

const renderProfile = (userId, isEditable) => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(
		<Profile userId={userId} isEditable={isEditable} />,
		pageContainer
	)
}

function Profile({ userId, isEditable }) {
	const { currentUserId, isLoading, error } = getCurrentUserId()
	const [profileUserData, setProfileUserData] = useState({})
	const [userPostData, setUserPostData] = useState([])
	const [userFollowerData, setUserFollowerData] = useState([])
	const [userFollowsData, setUserFollowsData] = useState([])
	const [isPublicValue, setIsPublicValue] = useState(null)
	const [isFollowed, setIsFollowed] = useState(false)

	useEffect(() => {
		fetchProfileData()
	}, [userId])

	useEffect(() => {
		if (!isPublicValue && !isEditable && currentUserId) {
			checkIfFollowed(currentUserId)
		}
	}, [isPublicValue, isEditable, currentUserId])

	const fetchProfileData = async () => {
		try {
			const response = await fetch(
				`http://localhost:8080/api/profile/${userId}`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
					},
				}
			)

			if (!response.ok) {
				throw new Error(
					`Failed to fetch profile data: ${response.status} ${response.statusText}`
				)
			}

			const data = await response.json()
			setProfileUserData(data.profileUserData)
			setUserPostData(data.userPostData || [])
			setUserFollowerData(data.userFollowerData || [])
			setUserFollowsData(data.userFollowsData || [])
			setIsPublicValue(data.profileUserData.isPublic)
			console.log("This is my data with followers", data)
		} catch (error) {
			console.error("Error fetching profile data:", error)
		}
	}

	const checkIfFollowed = async (currentUserId) => {
		try {
			const response = await fetch(
				`http://localhost:8080/api/users/${currentUserId}/userUsers/${userId}`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
					},
				}
			)

			if (response.ok) {
				setIsFollowed(true)
				console.log("checkIfFollowed.  isFollowed", isFollowed)
				console.log("response", response)
			} else if (response.status === 404) {
				setIsFollowed(false)
				console.log("checkIfFollowed.  isFollowed", isFollowed)
			} else {
				console.error("Error fetching user user data:", response.statusText)
			}
		} catch (error) {
			console.error("Error fetching user user data:", error)
		}
	}

	const handlePrivacyChange = (event) => {
		const newPrivacySetting = JSON.parse(event.target.value)

		setIsPublicValue(newPrivacySetting)

		fetch("http://localhost:8080/api/profile/privacy", {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				userId: profileUserData.userId,
				isPublic: newPrivacySetting,
			}),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error("Failed to update privacy status")
				}
			})
			.catch((error) => {
				console.error("Error updating privacy status:", error)
				setIsPublicValue(!newPrivacySetting)
			})
	}

	return (
		<div>
			<div id="profileData">
				<h2>{profileUserData.username}'s Profile</h2>
				{!isEditable && (
					<FollowButton
						followerId={currentUserId}
						subjectId={userId}
						isFollowed={isFollowed}
					/>
				)}
				{isPublicValue || isEditable || isFollowed ? (
					<>
						{isEditable ? (
							<div id="isPublicToggle">
								<label>
									<input
										type="radio"
										value={true}
										checked={isPublicValue === true}
										onChange={handlePrivacyChange}
									/>
									Public
								</label>
								<label>
									<input
										type="radio"
										value={false}
										checked={isPublicValue === false}
										onChange={handlePrivacyChange}
									/>
									Private
								</label>
							</div>
						) : (
							<p>
								<strong>Privacy:</strong> {isPublicValue ? "Public" : "Private"}
							</p>
						)}

						<p>
							<strong>User ID:</strong> {profileUserData.userId}
						</p>
						<p>
							<strong>Username:</strong> {profileUserData.username}
						</p>
						<p>
							<strong>Email:</strong> {profileUserData.email}
						</p>
						<p>
							<strong>First Name:</strong> {profileUserData.firstName}
						</p>
						<p>
							<strong>Last Name:</strong> {profileUserData.lastName}
						</p>
						<p>
							<strong>Date of Birth:</strong>{" "}
							{new Date(profileUserData.dob).toLocaleDateString()}
						</p>
						<p>
							<strong>Bio:</strong> {profileUserData.bio}
						</p>
						<p>
							<strong>Image URL:</strong> {profileUserData.imageURL}
						</p>

						<h2>{profileUserData.username}'s Posts</h2>
						<div id="myPostsData">
							{userPostData.map((post) => (
								<div key={post.postId}>
									<p>
										<strong>Post ID:</strong> {post.postId}
									</p>
									<p>
										<strong>Created At:</strong> {post.createdAt}
									</p>
									<p>
										<strong>Body:</strong> {post.body}
									</p>
									<p>
										<strong>Image URL:</strong> {post.imageURL}
									</p>
								</div>
							))}
						</div>

						<h2>{profileUserData.username}'s Followers</h2>
						<div id="myFollowersData">
							{userFollowerData &&
								userFollowerData.map((follower) => (
									<p key={follower.username}>{follower.username}</p>
								))}
						</div>

						<h2>{profileUserData.username}'s Followed</h2>
						<div id="usersIFollowData">
							{userFollowsData &&
								userFollowsData.map((user) => (
									<p key={user.username}>{user.username}</p>
								))}
						</div>
					</>
				) : (
					<p>This profile is private.</p>
				)}
			</div>
		</div>
	)
}

const renderChat = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Chat />, pageContainer)
}

function Chat() {
	const [sendMessage, setSendMessage] = useState("")
	const [receiveMessage, setReceiveMessage] = useState("")

	let messages = document.getElementById("messages")

	const handleMessages = (e) => {
		setSendMessage(e.target.value)
	}

	const handleSubmit = (e) => {
		e.preventDefault()
		let bodymessage = { message: sendMessage }
		let obj = { code: 1, body: JSON.stringify(bodymessage) }
		socket.send(JSON.stringify(obj))
		setSendMessage("")
	}

	socket.onmessage = function (e) {
		let data = JSON.parse(e.data)
		let msg = JSON.parse(data.body).message
		// setReceiveMessage(msg)
		// console.log("receiveMessage:", receiveMessage)
		let entry = document.createElement("li")
		entry.appendChild(document.createTextNode(msg))
		messages.appendChild(entry)
	}

	const messageStyle = {
		color: "orange",
	}

	return (
		<div>
			<h1>Chat</h1>
			<ul id="messages" style={messageStyle}></ul>
			<form id="chatbox" onSubmit={handleSubmit}>
				<textarea onChange={handleMessages}></textarea>
				<button type="submit" className="btn btn-primary">
					send
				</button>
			</form>
		</div>
	)
}

function GroupDetails({ group }) {

	const [groupPosts, setGroupPosts] = useState([]);

    useEffect(() => {
        const fetchGroupPosts = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/groups/${group.groupId}/posts`);
                if (!response.ok) {
                    throw new Error('Failed to fetch group posts');
                }
                const posts = await response.json();
                setGroupPosts(posts);
				console.log("posts in groupDetails:", posts)
            } catch (error) {
                console.error('Error fetching group posts:', error);
            }
        };

        fetchGroupPosts();
    }, [group.groupId]);

	return (
		<div className="group-details">
			<h2>{group.title}</h2>
			<p>{group.description}</p>
			{/* <p>Members: {group.members}</p> */}
			{/* Add more details you want to display */}
			{/* Display the PostForm component for creating new posts */}

			   {/* Render group posts here */}
			   <ul>
                {groupPosts.map((post) => (
                    <li key={post.id}>{post.body}</li>
                ))}
            </ul>

			<PostFormGroup groupId={group.groupId} />
		</div>
	)
}

const renderGroup = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Group />, pageContainer)
}

function Group() {
	const [title, setTitle] = useState("")
	const [description, setDescription] = useState("")
	const [groupData, setGroupData] = useState([])
	const [selectedGroup, setSelectedGroup] = useState(null)
	//const [showGroupDetails, setShowGroupDetails] = useState(false);

	const fetchGroupData = async () => {
		try {
			const response = await fetch("http://localhost:8080/api/groups", {
				method: "GET",
				credentials: "include",
				headers: {
					"Content-Type": "application/json",
				},
			})

			if (!response.ok) {
				throw new Error("Failed to fetch group data")
			}

			const data = await response.json()
			setGroupData(data) // Set the fetched group data to state
			console.log("Fetched group data:", data)
		} catch (error) {
			console.error("Error fetching group data:", error)
		}
	}

	useEffect(() => {
		fetchGroupData()
	}, [])

	const create = async (e) => {
		e.preventDefault() // prevent reload.

		const groupData = new FormData()

		// Append form data
		groupData.append("group-title", title)
		groupData.append("group-description", description)

		console.log("Group data being sent to backend:", title)
		console.log("Group data being sent to backend:", description)

		// Send user data to golang api/PostHandler.go.
		await fetch("http://localhost:8080/api/groups", {
			method: "POST",
			credentials: "include",
			body: groupData,
		})

		setTitle("")
		setDescription("")
		document.getElementById("exampleTitle").value = ""
		document.getElementById("exampleDescription").value = ""

		fetchGroupData()
	}

	const handleGroupClick = (group) => {
		setSelectedGroup(group)
		//setShowGroupDetails(true);
	}

	const handleGoBack = () => {
		setSelectedGroup(null)
		setShowGroupDetails(false) // Update showGroupDetails to false when going back
	}

	return (
		<div>
			{selectedGroup ? (
				<div>
					<button onClick={() => setSelectedGroup(null)}>Go Back</button>
					<GroupDetails group={selectedGroup} />
				</div>
			) : (
				<div>
					<form onSubmit={create}>
						<div className="mb-3">
							<label htmlFor="exampleTitle" className="form-label">
								Title
							</label>
							<input
								type="text"
								className="form-control"
								id="exampleTitle"
								aria-describedby="emailHelp"
								value={title}
								onChange={(e) => setTitle(e.target.value)}
							/>
						</div>
						<div className="mb-3">
							<label htmlFor="exampleInputPassword1" className="form-label">
								Description
							</label>
							<input
								type="text"
								className="form-control"
								id="exampleDescription"
								value={description}
								onChange={(e) => setDescription(e.target.value)}
							/>
						</div>
						<button type="submit" className="btn btn-primary">
							Create
						</button>
					</form>

					<div id="groupData">
						{groupData !== null ? (
							groupData.map((group) => (
								<div key={group.title} onClick={() => handleGroupClick(group)}>
									<h3>{group.title}</h3>
									<p>{group.description}</p>
								</div>
							))
						) : (
							<div id="noGroupsError">There are no created groups yet</div>
						)}
					</div>
				</div>
			)}
		</div>
	)
}

const renderNotifications = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Notifications />, pageContainer)
}

function Notifications() {
	return (
		<div>
			<h1>Notifications</h1>
		</div>
	)
}

function FollowButton({ followerId, subjectId, isFollowed }) {
	const [isFollowing, setIsFollowing] = useState(isFollowed)
	useEffect(() => {
		setIsFollowing(isFollowed)
	}, [isFollowed])

	const handleFollowToggle = async () => {
		if (isFollowing) {
			// If already following, unfollow the user
			await handleUnfollow(followerId, subjectId)
		} else {
			// If not following, follow the user
			await handleFollow(followerId, subjectId)
		}
		// Toggle the local follow state
		setIsFollowing(!isFollowing)
	}

	const handleFollow = async (followerId, subjectId) => {
		try {
			const response = await fetch(
				`http://localhost:8080/api/users/${followerId}/userUsers/`,
				{
					method: "POST",
					credentials: "include",
					body: JSON.stringify({ subjectId }),
				}
			)

			if (response.ok) {
				console.log("Successfully followed the user.")
				return true // Return true if the follow request is successful
			} else {
				console.error("Failed to follow the user.")
			}
		} catch (error) {
			console.error("Error following the user:", error)
		}

		return false // Return false if the follow request fails
	}

	const handleUnfollow = async (followerId, subjectId) => {
		try {
			const response = await fetch(
				`http://localhost:8080/api/users/${followerId}/userUsers/${subjectId}`,
				{
					method: "DELETE",
					credentials: "include",
				}
			)

			if (response.ok) {
				console.log("Successfully unfollowed the user.")
				return true // Return true if the follow request is successful
			} else {
				console.error("Failed to unfollow the user.")
			}
		} catch (error) {
			console.error("Error following the user:", error)
		}

		return false // Return false if the follow request fails
	}

	return (
		<button className="btn btn-primary btn-sm" onClick={handleFollowToggle}>
			{isFollowing ? "Unfollow" : "Follow"}
		</button>
	)
}

// PostForm component
// This component renders a form for creating a new post.
// It accepts a `groupId` prop to determine the group for the post.
function PostForm({ groupId, followedUsers }) {
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
			</main>
		</div>
	)
}

function PostFormGroup({ groupId }) {
	const [body, setBody] = useState("");
	const [selectedFile, setSelectedFile] = useState(null);

	// Handler for form submission
	const submit = async (e) => {
		e.preventDefault(); // Prevent page reload

		const formData = new FormData();

		// Append form data
		formData.append("body", body);
		formData.append("groupId", groupId);
		if (selectedFile) {
			formData.append("image", selectedFile);
		}

		console.log("Form data being sent to backend: ", formData);

		try {
			// Send user data to the server
			await fetch("http://localhost:8080/api/posts", {
				method: "POST",
				credentials: "include",
				body: formData,
			});

			// Reset form fields after successful submission
			setBody("");
			setSelectedFile(null);
			document.getElementById("postFormBody").value = "";
		} catch (error) {
			console.error("Error submitting post:", error);
		}
	};

	// Handler for file selection
	const handleFileChange = (e) => {
		setSelectedFile(e.target.files[0]);
	};

	const handleSelectFile = () => {
		const fileInput = document.getElementById("fileInput");
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
							id="fileInput"
							accept="image/*"
							style={{ display: "none" }}
							onChange={handleFileChange}
						/>
					</div>
					<br /> {/* Line break */}
					<button className="w-100 btn btn-lg btn-primary" type="submit">
						Submit
					</button>
				</form>
			</main>
		</div>
	)
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
	const [isFollowing, setIsFollowing] = useState(false)
	const [body, setBody] = useState("")
	const [selectedFile, setSelectedFile] = useState(null)

	const milliseconds = post.post.createdAt
	const date = new Date(milliseconds)
	const formattedDate = date.toLocaleString()

	const handleFollowClick = async () => {
		const followSuccess = await handleFollow(post.post.userId)
		setIsFollowing(followSuccess)
	}

	const submit = async (e) => {
		e.preventDefault() // prevent reload.

		const formData = new FormData()

		// Append form data
		formData.append("body", body)
		formData.append("postId", post.post.postId)
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
			`commentFileInput${post.post.postId}`
		)
		commentFileInput.click()
	}

	return (
		<div className="card" style={postCardStyle}>
			<div className="card-body">
				<div className="d-flex flex-start align-items-center">
					<img
						className="rounded-circle shadow-1-strong me-3"
						src={post.post.imageURL}
						alt="avatar"
						width="60"
						height="60"
					/>
					<div>
						<div className="d-flex align-items-center mb-1">
							<a
								className="fw-bold text-primary mb-0 me-2"
								href="#"
								onClick={() => renderProfile(post.post.userId)}
							>
								{post.post.userId}
							</a>

							<button
								className="btn btn-primary btn-sm"
								onClick={handleFollowClick}
								disabled={isFollowing}
							>
								{isFollowing ? "Following" : "Follow"}
							</button>
						</div>
						<p className="text-muted small mb-0">{formattedDate}</p>
					</div>
				</div>
				{/* Image, if there is one */}
				{!post.post.imageURL ? null : (
					<p className="mt-3 mb-2 pb-1">
						<img src={post.post.imageURL} className="img-fluid" />
					</p>
				)}
				{/* Post Body */}
				<p className="mt-3 mb-2 pb-1">{post.post.body}</p>
			</div>
			<div
				className="card-footer py-3 border-0"
				style={{ backgroundColor: "#f8f9fa" }}
			>
				<div className="d-flex flex-start w-100">
					<img
						className="rounded-circle shadow-1-strong me-3"
						src={post.avatar}
						alt="avatar"
						width="40"
						height="40"
					/>
					<div className="form-outline w-100">
						<textarea
							className="form-control"
							id="commentTextArea"
							rows="4"
							style={{ background: "#fff" }}
							onChange={(e) => setBody(e.target.value)}
						></textarea>

						<label className="form-label" htmlFor="textAreaExample">
							Message
						</label>
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
						id={`commentFileInput${post.post.postId}`}
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
				<div className="comments">
					<h2>Comments</h2>
					{post.comments !== null && post.comments.length > 0 ? (
						post.comments.map((comment) => (
							<CommentCard key={comment.createdAt} comment={comment} />
						))
					) : (
						<p className="text-muted">No comments</p>
					)}
				</div>
			</div>
		</div>
	)
}

function CommentCard({ comment }) {
	const formattedDate = new Date(comment.createdAt).toLocaleString()

	return (
		<div className="card mt-3">
			<div className="d-flex flex-start align-items-center">
				<img
					className="rounded-circle shadow-1-strong me-3"
					src={comment.imageURL}
					alt="avatar"
					width="60"
					height="60"
				/>
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
	)
}

const renderHome = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Home />, pageContainer)
}

// Display information relating to homepage
function Home() {
	const { currentUserId, isLoading, error } = getCurrentUserId()
	const [userList, setUserList] = useState([])
	const [followedUsers, setFollowedUsers] = useState([]);
	const [almostPrivatePosts, setAlmostPrivatePosts] = useState([])
	const [privatePosts, setPrivatePosts] = useState([])
	const [publicPostsWithComments, setPublicPostsWithComments] = useState([])
	const [userGroups, setUserGroups] = useState([])

	useEffect(() => {
		fetch("http://localhost:8080/api/home")
			.then((response) => response.json())
			.then((data) => {
				setUserList(data.userList)
				setAlmostPrivatePosts(data.almostPrivatePosts)
				setPrivatePosts(data.privatePosts)
				setPublicPostsWithComments(data.publicPostsWithComments)
				setUserGroups(data.userGroups)
			})
			.catch((error) => {
				console.error("Error fetching data:", error)
			})
	}, [])

	useEffect(() => {
		// Filter userList to get only the followed users
		const filteredFollowedUsers = userList.filter(user => user.isFollowed === true);

		// Set the filtered list to followedUsers state
		setFollowedUsers(filteredFollowedUsers);
	}, [userList]);


	return (
		<main className="homePage">
			<PostForm groupId={0} followedUsers={followedUsers} />
			<div className="userList">
				<h2>UserList</h2>
				{userList !== null && userList.length > 0 ? (
					userList.map((user, index) => (
						<div key={index}>
							<a
								className="nav-link"
								href="#"
								onClick={() => renderProfile(user.userId)}
							>
								{user.username}
							</a>
							<FollowButton
								followerId={currentUserId}
								subjectId={user.userId}
								isFollowed={user.isFollowed}
							/>
						</div>
					))
				) : (
					<p>No Users?!</p>
				)}
			</div>

			{/* Rendering Almost Private Posts */}
			<div className="almostPrivatePosts">
				<h2>Almost Private Posts</h2>
				{almostPrivatePosts !== null && almostPrivatePosts.length > 0 ? (
					almostPrivatePosts.map((almostPrivatePost) => (
						<PostCard
							key={almostPrivatePost.createdAt}
							post={almostPrivatePost}
						/>
					))
				) : (
					<p>No almost private posts</p>
				)}
			</div>

			{/* Rendering Private Posts */}
			<div className="privatePosts">
				<h2>Private Posts</h2>
				{privatePosts !== null && privatePosts.length > 0 ? (
					privatePosts.map((privatePost) => (
						<PostCard key={privatePost.createdAt} post={privatePost} />
					))
				) : (
					<p>No private posts</p>
				)}
			</div>

			{/* Rendering Public Posts */}
			<div className="publicPostsWithComments">
				<h2>Public Posts With Comments</h2>
				{publicPostsWithComments !== null &&
				publicPostsWithComments.length > 0 ? (
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
					{userGroups !== null &&
						userGroups.map((userGroup) => (
							<li key={userGroup.createdAt}>
								{userGroup.Title}{" "}
								{/* Render whatever user properties you need */}
							</li>
						))}
				</ul>
			</div>
		</main>
	)
}

const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
