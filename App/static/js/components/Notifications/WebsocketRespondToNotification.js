export const websocketRespondToNotification = (reply, notification) => {
  const notificationResponse = {
    reply: reply,
    notification: notification
  };
  const JSONnotificationResponse = JSON.stringify(notificationResponse);
  let obj = {
    code: 11,
    body: JSON.stringify(JSONnotificationResponse)
  };
  socket.send(JSON.stringify(obj));
};