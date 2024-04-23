const {
  useState
} = React;
export const renderChat = ({
  socket
}) => {
  const pageContainer = document.querySelector(".page-container");
  ReactDOM.render( /*#__PURE__*/React.createElement(Chat, {
    socket: socket
  }), pageContainer);
};
export function Chat({
  socket
}) {
  const [sendMessage, setSendMessage] = useState("");
  const [receiveMessage, setReceiveMessage] = useState("");
  let messages = document.getElementById("messages");
  const handleMessages = e => {
    setSendMessage(e.target.value);
  };
  const handleSubmit = e => {
    e.preventDefault();
    let bodymessage = {
      message: sendMessage
    };
    let obj = {
      code: 1,
      body: JSON.stringify(bodymessage)
    };
    socket.send(JSON.stringify(obj));
    setSendMessage("");
  };
  socket.onmessage = function (e) {
    let data = JSON.parse(e.data);
    let msg = JSON.parse(data.body).message;
    // setReceiveMessage(msg)
    // console.log("receiveMessage:", receiveMessage)
    let entry = document.createElement("li");
    entry.appendChild(document.createTextNode(msg));
    messages.appendChild(entry);
  };
  const messageStyle = {
    color: "orange"
  };
  return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Chat"), /*#__PURE__*/React.createElement("ul", {
    id: "messages",
    style: messageStyle
  }), /*#__PURE__*/React.createElement("form", {
    id: "chatbox",
    onSubmit: handleSubmit
  }, /*#__PURE__*/React.createElement("textarea", {
    onChange: handleMessages
  }), /*#__PURE__*/React.createElement("button", {
    type: "submit",
    className: "btn btn-primary"
  }, "send")));
}