const { useState, useEffect } = React
import { getCurrentUserId } from "./shared/getCurrentUserId.js"
import { GroupInvite } from "./components/GroupInvite.js"
import { GroupRequest } from "./components/GroupRequest.js"

export const renderNotifications = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Notifications />, pageContainer)
}

export function Notifications() {
	const { currentUserId } = getCurrentUserId()
	const [notifications, setNotifications] = useState(null);


	useEffect(() => {
		if (currentUserId !== null) {
			fetchNotifications();
		}
	}, [currentUserId]);

	const fetchNotifications = () => {
		fetch(`http://localhost:8080/api/users/${currentUserId}/notifications`)
			.then((response) => response.json())
			.then((data) => {
				setNotifications(data);
			})
			.catch((error) => {
				console.error("Error fetching notifications data:", error);
			});
	};


	return (
		<div>
			<h1>Notifications</h1>
			{notifications !== null && Object.keys(notifications).length > 0 ? (
				<ul>
					{Object.values(notifications).map((notification, index) => (
						<li key={index}><GroupInvite notification={notification} /></li>
					))}
				</ul>
			) : (
				<div>No notifications</div>
			)}

		</div>
	);
}



