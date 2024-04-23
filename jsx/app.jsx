import { Login } from "./Login.js"
let socket = null;

export function initializeSocket() {
	if (!socket) {
		socket = new WebSocket("ws://localhost:8080/ws");
		socket.onopen = function (event) {
			console.log("WebSocket connection established.");
		};
	}
	return socket;
}

export function getSocket() {
	return socket;
}

const App = () => {
	return (
		<div className="app-container">
			<div className="nav-container">
			</div>
			<div className="page-container">
				<Login />
			</div>
		</div>
	)
}

const root = document.querySelector("#root")
ReactDOM.render(<App />, root)
