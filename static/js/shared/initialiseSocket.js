let socket = null;
export function initializeSocket() {
  if (!socket) {
    socket = new WebSocket("ws://localhost:8080/ws");
    socket.onopen = function (event) {
      console.log("WebSocket connection established.");
    };
  }
  return socket;
}