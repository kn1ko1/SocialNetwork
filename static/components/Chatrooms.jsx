import React, { useState } from "react";
import { CreateChat } from "./ChatroomForm";
import { ChatBox } from "./Chatbox";

export const GetChat = () => {
  const [isPrivate, setIsPrivate] = useState(false);
  const [chats, setChats] = useState([]);
  const [visible, setVisible] = useState(false);

  const displayPrivateChatRooms = (privateChat) => {
    if (privateChat) {
      fetch("http://localhost:8080/get-chat")
        .then((response) => response.json())
        .then((data) => {
          //sort here
          data["private-chatrooms"].map((chat) => {
            console.log({ chat }, chat["last-message-date"]);
          });
          setChats(data);
        });
      setIsPrivate(true);
      console.log("private rooms", isPrivate, chats);
    } else {
      fetch("http://localhost:8080/get-chat")
        .then((response) => response.json())
        .then((data) => {
          //sort here
          setChats(data);
        });
      setIsPrivate(false);
      console.log("group rooms", isPrivate, chats);
    }
  };

  const closeChatRooms = () => {
    setIsPrivate(false);
    setVisible((prev) => !prev);
  };

  const openChatRooms = () => {
    fetch("http://localhost:8080/get-chat")
      .then((response) => response.json())
      .then((data) => {
        setChats(data);
      });
    setVisible((prev) => !prev);
  };

  const newChatCreated = (chatInfo) => {
    console.log(chatInfo);
    if (chatInfo["chat-type"] === "private") {
      setChats((chatrooms) => {
        let privateChats = chatrooms["private-chatrooms"] || [];
        return {
          ...chatrooms,
          "private-chatrooms": [...privateChats, chatInfo],
        };
      });
    } else {
      setChats((chatrooms) => {
        let groupChats = chatrooms["group-chatrooms"] || [];
        return { ...chatrooms, "group-chatrooms": [...groupChats, chatInfo] };
      });
    }
  };

  const checkGroupDisplay = () => {
    fetch("http://localhost:8080/get-chat")
      .then((response) => response.json())
      .then((data) => {
        setChats(data);
      });
  };

  return (
    <>
      {visible && (
        <>
          <div
            className="open-chat-container"
            style={{
              zIndex: 3,
              boxShadow:
                "rgb(0 0 0 / 10%) 0px 4px 6px -1px, rgb(0 0 0 / 6%) 0px 2px 4px -1px",
            }}
          >
            <div className="open-chat-close-container">
              <button
                className="open-chat-close-button"
                type="button"
                onClick={closeChatRooms}
              >
                <span>&times;</span>
              </button>
              <h1>Chat Rooms</h1>
              <CreateChat onSubmit={newChatCreated} />
            </div>
            <div className="chatroom-type">
              <div>
                <input
                  type="radio"
                  name="chat-type"
                  id="group"
                  value="group"
                  onChange={() => displayPrivateChatRooms(false)}
                  defaultChecked
                />
                <label htmlFor="group">Group</label>
              </div>
              <div>
                <input
                  type="radio"
                  name="chat-type"
                  id="private"
                  value="private"
                  onChange={() => displayPrivateChatRooms(true)}
                />
                <label htmlFor="private">Private</label>
              </div>
            </div>
            <div className="chatrooms">
              {isPrivate ? (
                <>
                  {chats["private-chatrooms"] ? (
                    <>
                      {chats["private-chatrooms"].map((chat) => (
                        <ChatBox
                          r={chat["chatroom-id"]}
                          n={""}
                          u={chat["users"]}
                          t={isPrivate}
                          i={chat["chat-avatar"]}
                          onClose={checkGroupDisplay}
                        />
                      ))}
                    </>
                  ) : (
                    <h1>No Private Chats Yet?</h1>
                  )}
                </>
              ) : (
                <>
                  {chats["group-chatrooms"] ? (
                    <>
                      {chats["group-chatrooms"].map((chat) => (
                        <ChatBox
                          r={chat["chatroom-id"]}
                          n={chat["chat-name"]}
                          u={chat["users"]}
                          t={isPrivate}
                          i={chat["chat-avatar"]}
                          onClose={checkGroupDisplay}
                        />
                      ))}
                    </>
                  ) : (
                    <>
                      <h1>No Group Chats Yet!</h1>
                    </>
                  )}
                </>
              )}
            </div>
          </div>
          <div
            style={{
              position: "absolute",
              width: "100%",
              height: "100vh",
              backgroundColor: "rgba(0, 0, 0, 0.2)",
              zIndex: 2,
            }}
          ></div>
        </>
      )}

      <button id="open-chat-button">
        <img
          src="https://cdn-icons-png.flaticon.com/512/5780/5780993.png"
          onClick={openChatRooms}
          alt=""
        />
      </button>
    </>
  );
};