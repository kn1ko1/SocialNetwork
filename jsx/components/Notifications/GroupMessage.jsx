const { useState, useEffect } = React
import { fetchUserById } from "../shared/FetchUserById.js";
import { fetchGroupById } from "../shared/FetchGroupById.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";


export function GroupMessage({ notification }) {
    const [username, setUsername] = useState("");
    const [groupName, setGroupName] = useState("")
    const { senderId, targetId, body } = notification;
    console.log("senderId, body:", senderId, body)
    useEffect(() => {
        fetchUserById(senderId)
            .then(user => setUsername(user.username));
        fetchGroupById(targetId)
            .then(group => setGroupName(group.title));
    }, [senderId]);


    return (
        <div id={"GroupMessage"} style={notificationCardStyle} className="card">
            <div className="row">
                <div className="col">
                   {groupName}, {username}:
                </div>
                <div className="col-auto d-flex align-items-center">
                    {body}
                </div>
            </div>
        </div>
    );
}