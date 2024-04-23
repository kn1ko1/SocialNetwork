const { useState, useEffect } = React
import { initializeSocket } from "./app.js"
import { renderNavbar } from "./components/Navbar.js"
import { renderRegister } from "./Register.js"
import { renderHome } from "./Home.js"

export const renderLogin = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Login />, pageContainer)
}

export function Login() {
	let socket = null
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
			socket = initializeSocket()
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