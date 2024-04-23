const { useState } = React
import { initializeSocket } from "../shared/socket.js";

export const renderRegister = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Register />, pageContainer)
}


export function Register() {
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
		socket = initializeSocket()
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