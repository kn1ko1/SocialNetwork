const { useState, useEffect } = React
import { fetchUserById } from "../shared/FetchUserById.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";


export function PrivateMessage({ notification }) {
    const [username, setUsername] = useState("");
    const { senderId, body } = notification;
    console.log("senderId, body:", senderId, body)
    useEffect(() => {
        fetchUserById(senderId)
            .then(user => setUsername(user.username));
    }, [senderId]);


    return (
        <div id={"privateMessage"} style={notificationCardStyle} className="card">
            <div className="row">
                <div className="col">
                    {username}:
                </div>
                <div className="col-auto d-flex align-items-center">
                    {body}
                </div>
            </div>
        </div>
    );
}