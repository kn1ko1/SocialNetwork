const {
  createContext,
  useContext,
  useState
} = React;
const SocketContext = createContext();
export const useSocket = () => {
  const context = useContext(SocketContext);
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
  const [socket, setSocket] = useState(null);
  const [currentUserId, setCurrentUserId] = useState(null);
  const updateContext = (newSocket, userId) => {
    setSocket(newSocket);
    setCurrentUserId(userId);
    console.log("userId in UserProvider", userId);
  };
  return /*#__PURE__*/React.createElement(SocketContext.Provider, {
    value: {
      socket,
      currentUserId,
      updateContext
    }
  }, children);
};