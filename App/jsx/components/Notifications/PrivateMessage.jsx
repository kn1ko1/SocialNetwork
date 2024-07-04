import { notificationCardStyle } from "./NotificationCardStyle.jsx";


export function PrivateMessage({ notification }) {

    return (
        <div id={"privateMessage"} style={notificationCardStyle} className="card">
            <div className="row">
                <div className="col">
                    {notification.senderUsername}:
                </div>
                <div className="col-auto d-flex align-items-center">
                    {notification.body}
                </div>
            </div>
        </div>
    );
}