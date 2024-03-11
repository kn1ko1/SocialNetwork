import { useState } from "react";
import { jsx as _jsx } from "react/jsx-runtime";
import { jsxs as _jsxs } from "react/jsx-runtime";
import { Fragment as _Fragment } from "react/jsx-runtime";
export const AddGroupComment = newComment => {
  const [urlImage, setUrlImage] = useState("");
  const [selectedImage, setSelectedImage] = useState(null);
  const [localImage, setLocalImage] = useState("");
  const [emoji, setEmoji] = useState("");
  const [thread, setThread] = useState("");
  const [threadArr, setThreadArr] = useState([]);
  const [visible, setVisible] = useState(false);
  const [local, setLocal] = useState(false);
  const [errorMes, setErrorMes] = useState("");
  const postId = newComment.id;
  const openCommentForm = () => {
    setVisible(prev => !prev);
  };
  const closeCommentForm = () => {
    setVisible(prev => !prev);
  };
  const handleLocalChange = location => {
    if (location) {
      setLocal(true);
    } else {
      setLocal(false);
    }
  };
  const addThread = () => {
    if (thread != "") {
      let hashtag = "#" + thread;
      setThreadArr(threadArr => {
        if (threadArr !== null) {
          return [...threadArr, hashtag];
        } else {
          return [hashtag];
        }
      });
      setThread("");
    }
  };
  const removeThread = index => {
    const newThreads = threadArr.filter((_, i) => i !== index);
    setThreadArr(newThreads);
  };
  const handleCommentSubmit = evt => {
    evt.preventDefault();
    const data = new FormData(evt.target);
    let values = Object.fromEntries(data.entries());
    if (local) {
      values["comment-image"] = localImage;
    } else {
      values["comment-image"] = urlImage;
    }
    if (threadArr.length != 0) {
      values["comment-threads"] = threadArr.join(",");
    }
    values['comment-time'] = new Date().getTime();
    fetch("http://localhost:8080/create-group-post-comment", {
      method: "POST",
      headers: {
        'Content-Type': "multipart/form-data"
      },
      body: JSON.stringify(values)
    }).then(response => response.json())
    // return array of posts and send to the top.
    .then(response => {
      console.log(response);
      if (response.hasOwnProperty("error")) {
        setErrorMes(response["error"]);
        setTimeout(() => {
          setErrorMes("");
        }, 5000);
      } else {
        newComment["onSubmit"](response);
        closeCommentForm();
      }
    });
  };
  const handleKeyPress = evt => {
    if (evt.key === "#") {
      evt.preventDefault();
    }
  };
  return /*#__PURE__*/_jsxs(_Fragment, {
    children: [visible && /*#__PURE__*/_jsx("div", {
      className: "create-comment-container",
      children: /*#__PURE__*/_jsxs("form", {
        className: "add-comment-form",
        onSubmit: handleCommentSubmit,
        children: [/*#__PURE__*/_jsx("button", {
          className: "close-button",
          type: "button",
          onClick: closeCommentForm,
          children: /*#__PURE__*/_jsx("span", {
            children: "\xD7"
          })
        }), /*#__PURE__*/_jsx("h1", {
          children: "Create Comment "
        }), /*#__PURE__*/_jsx("input", {
          type: "hidden",
          name: "post-id",
          value: postId
        }), /*#__PURE__*/_jsxs("div", {
          className: "image-location",
          children: [/*#__PURE__*/_jsxs("div", {
            children: [/*#__PURE__*/_jsx("input", {
              type: "radio",
              id: "Url",
              name: "img-location",
              value: "Url",
              onChange: () => handleLocalChange(false),
              defaultChecked: true
            }), /*#__PURE__*/_jsx("label", {
              htmlFor: "Url",
              children: "Add Online Image"
            })]
          }), /*#__PURE__*/_jsxs("div", {
            children: [/*#__PURE__*/_jsx("input", {
              type: "radio",
              id: "local",
              name: "img-location",
              value: "local",
              onChange: () => handleLocalChange(true)
            }), /*#__PURE__*/_jsx("label", {
              htmlFor: "local",
              children: "Add Local Image"
            })]
          })]
        }), local ? /*#__PURE__*/_jsxs(_Fragment, {
          children: [selectedImage && /*#__PURE__*/_jsx("div", {
            className: "create-post-image-container",
            children: /*#__PURE__*/_jsx("img", {
              src: URL.createObjectURL(selectedImage),
              alt: "",
              onClick: () => {
                document.querySelector(".create-post-image").value = "";
                setLocalImage("");
                setSelectedImage(null);
              }
            })
          }), /*#__PURE__*/_jsx("div", {
            className: "add-post-image",
            children: /*#__PURE__*/_jsx("input", {
              type: "file",
              className: "create-post-image",
              onChange: e => {
                if (e.target.files[0].size < 20000000) {
                  setSelectedImage(e.target.files[0]);
                  const fileReader = new FileReader();
                  fileReader.onload = function (e) {
                    setLocalImage(e.target.result);
                  };
                  fileReader.readAsDataURL(e.target.files[0]);
                }
                ;
              }
            })
          })]
        }) : /*#__PURE__*/_jsxs(_Fragment, {
          children: [urlImage && /*#__PURE__*/_jsx("div", {
            className: "create-post-image-container",
            children: /*#__PURE__*/_jsx("img", {
              src: urlImage,
              alt: "",
              onClick: () => {
                document.querySelector(".create-post-image").value = "";
                setUrlImage("");
              }
            })
          }), /*#__PURE__*/_jsx("div", {
            className: "add-post-image",
            children: /*#__PURE__*/_jsx("input", {
              type: "text",
              className: "create-post-image",
              id: "create-post-image",
              placeholder: "https://...",
              onChange: e => setUrlImage(e.target.value)
            })
          })]
        }), /*#__PURE__*/_jsx("p", {
          children: "File Must Not Exceed 20MB"
        }), /*#__PURE__*/_jsx("textarea", {
          name: "comment-text",
          contentEditable: true,
          className: "post-text-content",
          onChange: e => setEmoji(e.target.value),
          placeholder: "For Emojis Press: 'Windows + ;' or 'Ctrl + Cmd + Space'"
        }), /*#__PURE__*/_jsxs("div", {
          className: "create-post-threads",
          children: [/*#__PURE__*/_jsx("input", {
            type: "text",
            className: "add-thread-input",
            placeholder: "Add Thread",
            value: thread,
            onChange: e => setThread(e.target.value),
            onKeyPress: handleKeyPress
          }), /*#__PURE__*/_jsx("button", {
            className: "add-thread-button",
            type: "button",
            onClick: addThread,
            children: "+"
          })]
        }), threadArr && /*#__PURE__*/_jsxs(_Fragment, {
          children: [/*#__PURE__*/_jsx("p", {
            className: "remove-thread",
            children: "Click the # to remove"
          }), /*#__PURE__*/_jsx("div", {
            className: "thread-container",
            children: threadArr.map((t, index) => /*#__PURE__*/_jsx("p", {
              className: "added-thread",
              onClick: () => removeThread(index),
              children: t
            }, index))
          })]
        }), errorMes && /*#__PURE__*/_jsx("p", {
          className: "error-message",
          children: errorMes
        }), /*#__PURE__*/_jsx("input", {
          type: "submit",
          className: "create-post-submit-button",
          value: "Create Comment"
        })]
      })
    }), /*#__PURE__*/_jsx("button", {
      type: "button",
      className: "add-comment-button",
      onClick: openCommentForm,
      children: " Add Comment"
    })]
  });
};