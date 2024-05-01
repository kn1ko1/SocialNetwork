import { Login } from "./Login.js";
const {
  createContext,
  useContext,
  useState
} = React;
const SocketContext = createContext();
export const useSocket = () => {
  const context = useContext(SocketContext);
  console.log("context", context); // Check if context is null or contains the SocketContext
  if (!context) {
    throw new Error('useSocket must be used within a UserProvider');
  }
  const {
    socket,
    currentUserId,
    updateContext
  } = context;
  return {
    socket,
    currentUserId,
    updateContext
  };
};
export const UserProvider = ({
  children
}) => {
  console.log("UserProvider rendered");
  const [socket, setSocket] = useState(null);
  const [currentUserId, setCurrentUserId] = useState(null);
  const updateContext = (newSocket, userId) => {
    console.log("Updating context with socket:", newSocket, "and userId:", userId);
    setSocket(newSocket);
    setCurrentUserId(userId);
  };
  return /*#__PURE__*/React.createElement(SocketContext.Provider, {
    value: {
      socket,
      currentUserId,
      updateContext
    }
  }, children);
};
const App = () => {
  console.log("App rendered");
  return /*#__PURE__*/React.createElement(UserProvider, null, /*#__PURE__*/React.createElement("div", {
    className: "app-container"
  }, /*#__PURE__*/React.createElement("div", {
    className: "nav-container"
  }), /*#__PURE__*/React.createElement("div", {
    className: "page-container"
  }, /*#__PURE__*/React.createElement(Login, null))));
};
const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/React.createElement(App, null), root);