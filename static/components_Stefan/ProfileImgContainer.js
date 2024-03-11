import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import Swal from "sweetalert2";
import FollowerWindow from "./FollowerWIndow";

// Component that contains image, followers etc.
import { jsx as _jsx } from "react/jsx-runtime";
import { jsxs as _jsxs } from "react/jsx-runtime";
import { Fragment as _Fragment } from "react/jsx-runtime";
export default function ProfileImgContainer(props) {
  // Check if we are at other user's profile. If so, show follow button instead of my profile button.
  const otherUser = window.location.href.split("/").at(-1);
  // Variable to check following status. Set to true if user presses follow or on refreshing the page.
  const [isFollowing, setIsFollowing] = useState(false);
  useEffect(() => {
    if (props.update) if (props.user.status === "private") {
      (async () => {
        const response = await fetch("http://localhost:8080/api/followers", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          credentials: "include",
          body: JSON.stringify({
            follower: props.currentUser ? props.currentUser.email : null,
            followee: props.user.email
          })
        });

        /* is this redundant ???? */
        let result = await response.json();

        // if (result === null) {
        // } else {
        //   setIsFollowing(true);
        // }
        // The above is the same as:
        setIsFollowing(result !== null);
        props["setUpdate"](false);
        return;
      })();
    }
  }, [props.update]);

  // Get all the followers of the user rendered on component.
  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:8080/api/followers", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        credentials: "include",
        body: JSON.stringify({
          follower: props.currentUser ? props.currentUser.email : null,
          followee: props.user.email
        })
      });
      let result = await response.json();
      // if (result === null) {
      // } else {
      //   setIsFollowing(true);
      // }
      // The above is the same as:
      setIsFollowing(result !== null);
    })();
  }, []);

  // Follow button handler
  const followHandler = () => {
    let newIsFollowing = !isFollowing;
    setIsFollowing(newIsFollowing);

    //if the user is private then wait for request to be accepted, and set IsFollowing back to false.
    if (props.user.status === "private" && newIsFollowing === true) {
      setIsFollowing(!newIsFollowing);
    }
    let follow = JSON.stringify({
      followRequest: props.currentUser.email,
      toFollow: props.user.email,
      isFollowing: newIsFollowing,
      followers: props.user.followers
    });
    if (newIsFollowing === false) {
      follow = JSON.stringify({
        followRequest: props.currentUser.email,
        toFollow: props.user.email,
        isFollowing: newIsFollowing,
        followers: props.user.followers,
        //send followRequest-accepted:true so it goes to the else condition in client's followMessage switch case.
        "followRequest-accepted": true
      });
    }
    props.socket.send(follow);
    props.fetchUsersData();
  };

  // Show and hide window.
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const showFollowers = async event => {
    // Get all followers for onclick event.
    const response = await fetch("http://localhost:8080/api/allFollowers", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      credentials: "include",
      // Send user whose followers to fetch.
      body: JSON.stringify(props.user)
    });
    let result = await response.json();
    if (event.target.id === "followerText") {
      setFollowers(result.followers);
      setFollowing([]);
    } else {
      setFollowing(result.following);
      setFollowers([]);
    }
  };
  const closeFollowersWindow = () => {
    setFollowers([]);
    setFollowing([]);
  };
  function showUserInfo(props) {
    // Show user details if following and private or if user is public. Else, display error.
    if (props.user.status === "private" && isFollowing || props.user.status === "public" || props.user.status === "") {
      // Sweet Alert notification
      // e.g. [1998, 03, 13]
      const birthday = props.user.dob.split("-");
      Swal.fire({
        title: props.name,
        html: "<div id='name-email-dob'>" + `<p> <span class="bold">Email:&nbsp</span> ${props.user.email} </p>` + `<p> <span class="bold">Nickname:&nbsp</span> ${props.user.nickname} </p>` + `<p> <span class="bold">Birthday:&nbsp</span> ${birthday[1] + " / " + birthday[2]} </p>` + "</div>",
        icon: "info",
        confirmButtonText: "Close"
      });
    } else if (props.user.status === "private" && !isFollowing) {
      Swal.fire({
        title: props.name,
        icon: "error",
        text: "Private information. Send a follow request to view!",
        confirmButtonText: "Ok"
      });
    }
  }
  return /*#__PURE__*/_jsxs("div", {
    className: "profileImgContainer",
    children: [props.name ? /*#__PURE__*/_jsxs("div", {
      className: "profileImgParent",
      children: [props.avatar ? /*#__PURE__*/_jsx("img", {
        className: "profileImg",
        src: props.avatar,
        alt: props.user.name + "'s profile image"
      }) : /*#__PURE__*/_jsx("img", {
        className: "profileImg",
        src: "https://www.transparentpng.com/thumb/user/gray-user-profile-icon-png-fP8Q1P.png",
        alt: "No Image"
      }), /*#__PURE__*/_jsxs("span", {
        className: "firstLast",
        children: [props.name, " ", props.user.last]
      }), /*#__PURE__*/_jsx("p", {
        className: "aboutme",
        children: props.user.aboutme
      }), /*#__PURE__*/_jsx("hr", {
        className: "break"
      }), /*#__PURE__*/_jsxs("div", {
        className: "followerDiv",
        children: [/*#__PURE__*/_jsx("div", {
          children: /*#__PURE__*/_jsx("span", {
            className: "followerFollowing",
            id: "followingText",
            onClick: showFollowers,
            children: "Following"
          })
        }), /*#__PURE__*/_jsx("div", {
          children: /*#__PURE__*/_jsx("span", {
            className: "count",
            id: "following",
            children: props.user.following
          })
        })]
      }), /*#__PURE__*/_jsx("hr", {
        className: "break"
      }), /*#__PURE__*/_jsxs("div", {
        className: "followerDiv",
        children: [/*#__PURE__*/_jsx("div", {
          children: /*#__PURE__*/_jsx("span", {
            className: "followerFollowing",
            id: "followerText",
            onClick: showFollowers,
            children: "Followers"
          })
        }), /*#__PURE__*/_jsx("div", {
          children: /*#__PURE__*/_jsx("span", {
            className: "count",
            id: `${props.user.email}-followers`,
            children: props.user.followers
          })
        })]
      }), /*#__PURE__*/_jsx("hr", {
        className: "break"
      }), /*#__PURE__*/_jsx("div", {
        className: "followerDiv",
        children: otherUser && otherUser !== "profile" ? /*#__PURE__*/_jsxs(_Fragment, {
          children: [/*#__PURE__*/_jsx("span", {
            children: /*#__PURE__*/_jsx("button", {
              className: "redText",
              style: {
                backgroundColor: "transparent",
                border: "none",
                marginBottom: "0px"
              },
              onClick: followHandler,
              children: isFollowing ? "Unfollow" : "Follow"
            })
          }), /*#__PURE__*/_jsx("hr", {
            className: "break"
          }), /*#__PURE__*/_jsx("span", {
            children: /*#__PURE__*/_jsx("button", {
              className: "moreInfo",
              style: {
                backgroundColor: "transparent",
                border: "none",
                marginBottom: "15px",
                fontSize: "large",
                fontWeight: 500
              },
              onClick: () => showUserInfo(props),
              children: "More Info"
            })
          })]
        }) : otherUser === "profile" ? /*#__PURE__*/_jsx(Link, {
          to: "/",
          style: {
            textDecoration: "none",
            marginBottom: "15px"
          },
          children: /*#__PURE__*/_jsx("span", {
            className: "redText",
            children: "Back"
          })
        }) : /*#__PURE__*/_jsx(Link, {
          to: "/profile",
          style: {
            textDecoration: "none",
            marginBottom: "15px"
          },
          children: /*#__PURE__*/_jsx("span", {
            className: "redText",
            children: "My Profile"
          })
        })
      })]
    }) : /*#__PURE__*/_jsx("div", {
      children: " loading... "
    }), /*#__PURE__*/_jsx(FollowerWindow, {
      followers: followers,
      following: following,
      closeFollowersWindow: closeFollowersWindow
    })]
  });
}