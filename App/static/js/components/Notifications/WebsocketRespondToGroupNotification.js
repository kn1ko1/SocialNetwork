export function websocketRespondToGroupNotification(notificationObject, socket) {
  socket.send(JSON.stringify(notificationObject));
}
;