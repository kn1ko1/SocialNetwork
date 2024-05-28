

export const NotificationPopUp = ({ message, onClose }) => {
    return (
        <div className="toast show position-fixed bottom-0 end-0 p-3 m-3" style={{ zIndex: 1000 }}>
        <div className="toast-header">
            <strong className="me-auto">Notification</strong>
            <button type="button" className="btn-close" aria-label="Close" onClick={onClose}></button>
        </div>
        <div className="toast-body">
            {message}
        </div>
    </div>
    );
};
