const { useState, useEffect } = React
import { fetchGroupById } from "../shared/FetchGroupById.js";
import { notificationCardStyle } from "./NotificationCardStyle.js";


export function GroupMessage({ notification }) {
    const [groupName, setGroupName] = useState("")
    const { targetId } = notification;
    useEffect(() => {
        fetchGroupById(targetId)
            .then(group => setGroupName(group.title));
    }, [notification.messageId]);


    return (
        <div id={"GroupMessage"} style={notificationCardStyle} className="card">
            <div className="row">
                <div className="col">
                   {groupName}, {notification.senderUsername}:
                </div>
                <div className="col-auto d-flex align-items-center">
                    {notification.body}
                </div>
            </div>
        </div>
    );
}