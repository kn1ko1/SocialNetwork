export const respondToNotification = (reply, notification) => {
    const data = {
        reply: reply,
        notification: notification
    };

    fetch(`http://localhost:8080/api/notifications/${notification.notificationId}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        // Handle success response
        console.log("Response sent successfully:", data);
    })
    .catch(error => {
        console.error("Error sending response:", error);
    });
};