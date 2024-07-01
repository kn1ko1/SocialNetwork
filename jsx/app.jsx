import { Home, renderHome } from "./Home.js";
import { Login } from "./Login.js"
import { renderNavbar } from "./components/shared/Navbar.js";


export function initializeSocket() {
	// if (!socket) {
	let socket = new WebSocket("ws://localhost:8080/ws");
	socket.onopen = function (event) {
		console.log("WebSocket connection established.");
	};
	// }
	return socket;
}

function getCookieValue(name) {
	const value = `; ${document.cookie}`;
	const parts = value.split(`; ${name}=`);
	if (parts.length === 2) return parts.pop().split(';').shift();
	return null;
}

const App = () => {
	let socket = null
	const cookieValue = getCookieValue("SessionID"); // Correctly assign the value without destructuring
	if (cookieValue !== null) { // Correct syntax for if statement
		console.log("Cookie Value is:", cookieValue);
		socket = initializeSocket()
		renderNavbar({socket})
		renderHome({socket})
	}

	return (
		<div className="app-container">
			<div className="nav-container">
			</div>
			<div className="page-container">
				{cookieValue ?
				(<Home socket={socket} />) :
				(<Login />)}


			</div>
		</div>
	)
}

const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
