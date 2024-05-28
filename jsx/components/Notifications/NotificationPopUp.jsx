const GROUP_CHAT_MESSAGE = 1;
const PRIVATE_MESSAGE = 2;
const CREATE_EVENT = 3;

export const NotificationPopUp = ({ data, onClose }) => {
    let message = JSON.parse(data.body).body;
    let code = JSON.parse(data.code);
    console.log("socket message data:", data)
    return (
        <div id="notificationPopup">
            
            <div className="toast show position-fixed bottom-0 end-0 p-3 m-3" style={{ zIndex: 1000 }}>
                <div className="toast-header">
                    <strong className="me-auto">Notification</strong>
                    <button type="button" className="btn-close" aria-label="Close" onClick={onClose}></button>
                </div>
                <div className="toast-body">
                    {message}
                </div>
            </div>
        </div>

    );
};
