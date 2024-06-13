import { GroupMessage } from "./GroupMessage.js";
import { PrivateMessage } from "./PrivateMessage.js";
import { FollowRequest } from "./FollowRequest.js";
import { GroupRequest } from "./GroupRequest.js";
import { GroupInvite } from "./GroupInvite.js";
import { EventInvite } from "./EventInvite.js";

const codeToHeaderText = {
    1: "Group Chat Message",
    2: "Private Message",
    3: "Follow Request",
    4: "Group Request",
    5: "Group Invite",
    6: "Event Invite"
};



export const NotificationPopUp = ({ data, onClose }) => {
    try {
        const notification = JSON.parse(data.body);
        const code = parseInt(data.code, 10);
        console.log("socket message notification:", notification)

        // Get the header text based on the code
        const headerText = codeToHeaderText[code] || "Notification";
        return (
            <div id="notificationPopup">

                <div className="toast show position-fixed bottom-0 end-0 p-3 m-3" style={{ zIndex: 1000 }}>
                    <div className="toast-header">
                        <strong className="me-auto">{headerText}</strong>
                        <button type="button" className="btn-close" aria-label="Close" onClick={onClose}></button>
                    </div>
                    <div className="toast-body">
                    {(() => {
                    switch (code) {
                        case 1:
                            return <GroupMessage notification={notification} />;
                        case 2:
                            return <PrivateMessage notification={notification} />;
                        case 3:
                            return <FollowRequest notification={notification} onNotificationResponse={onClose} />;
                        case 4:
                            return <GroupRequest notification={notification} onNotificationResponse={onClose} />;
                        case 5:
                            return <GroupInvite notification={notification} onNotificationResponse={onClose} />;
                        case 6:
                            return <EventInvite notification={notification} onNotificationResponse={onClose} />;
                        default:
                            return null;
                    }
                })()}
                    </div>
                </div>
            </div>

        );
    } catch (error) {
        console.error("Error processing notification data:", error);
        return null;
    }
};
