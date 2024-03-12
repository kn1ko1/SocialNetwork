import React from 'react';
import { ToastContainer } from 'react-toastify';
import "react-toastify/dist/ReactToastify.css";
import 'react-toastify/dist/ReactToastify.min.css'


export function Notification(props) {
    return (
        <>
            <ToastContainer
                style={{
                    zIndex: "var(--toastify-z-index)",
                    padding: "4px",
                    position: "inherit",
                    width: "var(--toastify-toast-width)",
                    boxSizing: "border-box",
                    color: "#fff",
                    display: props.page === "profiles" && 'flex',
                    flexDirection: props.page === "profiles" && 'row',
                    width: props.page === "profiles" && "100%",
                }}
                autoClose={false}
                hideProgressBar={false}
                newestOnTop={true}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss
                draggable
                pauseOnHover
                theme="dark"
            />
            {/* Same as */}
        </>
    );
}   