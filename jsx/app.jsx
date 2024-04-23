import { Login } from "./Login.js"


export function initializeSocket() {
	// if (!socket) {
		let socket = new WebSocket("ws://localhost:8080/ws");
		socket.onopen = function (event) {
			console.log("WebSocket connection established.");
		};
	// }
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
