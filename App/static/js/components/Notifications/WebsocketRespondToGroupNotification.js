export function websocketRespondToGroupNotification(notification, socket) {
  console.log(notification);
  let obj = {
    code: 11,
    body: JSON.stringify(notification)
  };
  socket.send(JSON.stringify(obj));
}
;