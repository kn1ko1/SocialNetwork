// window.addEventListener("load", () => {
//     renderHomePage();
// });

// function renderHomePage() {
//     const main = document.getElementById("main");
//     main.innerHTML = "";
//     const h1 = document.createElement("h1");
//     h1.innerText = "Home";
//     main.appendChild(h1);
//     const button = document.createElement("button");
//     button.innerText = "Other Page";
//     button.type = "click";
//     button.addEventListener("click", () => {
//         renderOtherPage();
//     })
//     main.appendChild(button);
// }

// // Some react function
// function renderOtherPage(props) {
//     const main = document.getElementById("main");
//     main.innerHTML = "";
//     const h1 = document.createElement("h1");
//     h1.innerText = "Other";
//     main.appendChild(h1);
//     const button = document.createElement("button");
//     button.innerText = "Home Page";
//     button.type = "click";
//     button.addEventListener("click", () => {
//         renderHomePage();
//     })
//     main.appendChild(button);
// }

// import Login from "./js/Login.js"

const { useState } = React

const App = () => {
	return (
		<div className="app-container">
			<Login />
		</div>
	)
}

const Login = (props) => {
	const [email, setEmail] = useState("")
	const [password, setPassword] = useState("")
	const [redirectVar, setRedirectVar] = useState(false)

	// Redirect
	// const navigate = useNavigate()

	const submit = async (e) => {
		e.preventDefault() // prevent reload.

		// Create new user as JS object.
		const userToLogin = {
			email,
			password,
		}

		// Send user data to golang register function.
		const response = await fetch("http://localhost:8080/login", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			credentials: "include",
			body: JSON.stringify(userToLogin),
		})

		const validUser = await response.json()
		setRedirectVar(true)
		props.setName(validUser.first)
	}

	// if (redirectVar) {
	// 	return navigate("/")
	// }

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
							onChange={(e) => setEmail(e.target.value)}
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

const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
