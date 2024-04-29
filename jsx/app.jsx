import { UserProvider } from './shared/UserProvider.js';
import { Login } from "./Login.js"

const App = () => {
	return (
		<div className="app-container">
			<div className="nav-container">
			</div>
			<div className="page-container">
				<Login />
			</div>
		</div>
	);
};

const root = document.querySelector("#root");
ReactDOM.render(
	<UserProvider>
		<App />
	</UserProvider>,
	root
);

