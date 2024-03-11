import React, { useState, useEffect } from "react";
import { Post } from "./Post";

export const AllPosts = (props) => {
  const [posts, setPosts] = useState([]);
  const [usr, setUsr] = useState({});
  const [loaded, setLoaded] = useState(false);

  // Fetch public posts
  useEffect(() => {
    if (!loaded) {
      fetch("http://localhost:8080/view-public-posts")
        .then((response) => response.json())
        .then((data) => {
          if (
            (props.user && Object.keys(props.user).length !== 0) ||
            Object.keys(usr).length !== 0
          ) {
            const filteredPosts = data.filter(
              (post) =>
                post.author === props.user.nickname ||
                post.author === usr.nickname
            );
            setPosts(filteredPosts);
          } else if (props.user && Object.keys(props.user).length === 0) {
            //getting current user.
            (async () => {
              const response = await fetch("http://localhost:8080/api/user", {
                method: "GET",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
              });
              const user = await response.json();
              setUsr(user);
            })();
          } else {
            setPosts(data);
          }
          setLoaded(true);
        });
    }
  }, [loaded, usr]);

  // Fetch private posts and append to posts list if visiting a profile page.
  useEffect(() => {
    let filtered;
    (async () => {
      if (props.user && Object.keys(props.user).length !== 0) {
        fetch("http://localhost:8080/view-public-posts")
          .then((response) => response.json())
          .then((data) => {
            if (
              (props.user && Object.keys(props.user).length !== 0) ||
              Object.keys(usr).length !== 0
            ) {
              filtered = data.filter(
                (post) =>
                  post.author === props.user.nickname ||
                  post.author === usr.nickname
              );
            }
          });

        let privatePostsPromise = await fetch(
          "http://localhost:8080/view-private-posts"
        );
        let result = await privatePostsPromise.json();
        let filteredResult = result.filter(
          (post) => post.author === props.user.nickname
        );

        setPosts([...(posts || filtered), ...filteredResult]);
      }
    })();
  }, [loaded, usr]);

  //   Handle private posts button on homepage.
  useEffect(() => {
    setPosts(props["posts"]);
  }, [props["posts"]]);

  var ranges = [
    { divider: 1e18, suffix: "E" },
    { divider: 1e15, suffix: "P" },
    { divider: 1e12, suffix: "T" },
    { divider: 1e9, suffix: "G" },
    { divider: 1e6, suffix: "M" },
    { divider: 1e3, suffix: "k" },
  ];

  function formatNumber(n) {
    for (var i = 0; i < ranges.length; i++) {
      if (n >= ranges[i].divider) {
        return (
          (Math.round((n / ranges[i].divider) * 10) / 10).toString() +
          ranges[i].suffix
        );
      }
    }

    return n.toString();
  }

  const handleEditPost = (edited) => {
    console.log("edited post", { edited });
    setPosts((prevPosts) => {
      const index = prevPosts.findIndex(
        (post) => post["post-id"] === edited["post-id"]
      );
      console.log({ index });
      if (index === -1) {
        return prevPosts;
      }
      const newPost = [...prevPosts];
      edited["post-likes"] = formatNumber(edited["post-likes"]);
      edited["post-dislikes"] = formatNumber(edited["post-dislikes"]);
      newPost[index] = edited;
      return newPost;
    });
  };

  const handleDeletePost = (deletePost) => {
    const updatedPosts = posts.filter((post) => post["post-id"] !== deletePost);
    setPosts(updatedPosts);
  };

  return (
    <div className="post-container">
      {loaded &&
        posts &&
        posts
          .slice()
          .reverse()
          .map((post, index) => (
            <div key={index} className="post">
              <Post
                post={post}
                onEdit={handleEditPost}
                onDelete={handleDeletePost}
              />
            </div>
          ))}
      {!loaded && (
        <div className="post-loader-container">
          <img
            src="http://superstorefinder.net/support/wp-content/uploads/2018/01/orange_circles.gif"
            className="post-loader"
          />
        </div>
      )}
    </div>
  );
};
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