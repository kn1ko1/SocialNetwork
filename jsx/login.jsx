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