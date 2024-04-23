export const renderNotifications = () => {
	const pageContainer = document.querySelector(".page-container")
	ReactDOM.render(<Notifications />, pageContainer)
}

export function Notifications() {
	return (
		<div>
			<h1>Notifications</h1>
		</div>
	)
}