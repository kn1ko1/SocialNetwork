import { Login } from "./Login.js";

const { createContext, useContext, useState } = React;

const SocketContext = createContext();

export const useSocket = () => {
    const context = useContext(SocketContext);
    console.log("context", context); // Check if context is null or contains the SocketContext
    if (!context) {
        throw new Error('useSocket must be used within a UserProvider');
    }
    const { socket, currentUserId, updateContext } = context;
    return { socket, currentUserId, updateContext };
};


export const UserProvider = ({ children }) => {
    console.log("UserProvider rendered");

    const [socket, setSocket] = useState(null);
    const [currentUserId, setCurrentUserId] = useState(null);

    const updateContext = (newSocket, userId) => {
        console.log("Updating context with socket:", newSocket, "and userId:", userId);
        setSocket(newSocket);
        setCurrentUserId(userId);
    };

    return (
        <SocketContext.Provider value={{ socket, currentUserId, updateContext }}>
            {children}
        </SocketContext.Provider>
    );
};




const App = () => {
    console.log("App rendered");
    
    return (
        <UserProvider>
            <div className="app-container">
                <div className="nav-container"></div>
                <div className="page-container">
                    <Login />
                </div>
            </div>
        </UserProvider>
    );
};


const root = document.querySelector("#root");
ReactDOM.render(<App />, root);