import { jsx as _jsx } from "react/jsx-runtime";
import { jsxs as _jsxs } from "react/jsx-runtime";
import { Fragment as _Fragment } from "react/jsx-runtime";
const {
  useState
} = React;
const App = () => {
  return /*#__PURE__*/_jsxs("div", {
    className: "app-container",
    children: [/*#__PURE__*/_jsx(Login, {}), /*#__PURE__*/_jsx(Register, {}), /*#__PURE__*/_jsx(Home, {}), /*#__PURE__*/_jsx(Profile, {})]
  });
};
function Login(props) {
  const [usernameOrEmail, setUsernameOrEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirectVar, setRedirectVar] = useState(false);
  const submit = async e => {
    e.preventDefault(); // prevent reload.

    const userToLogin = {
      usernameOrEmail,
      password
    };
    console.log(userToLogin);

    // Send user data to golang register function.
    const response = await fetch("http://localhost:8080/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      credentials: "include",
      body: JSON.stringify(userToLogin)
    });
    const validUser = await response.json();
    setRedirectVar(true);
    props.setName(validUser.first);
  };
  return /*#__PURE__*/_jsx("div", {
    children: /*#__PURE__*/_jsxs("main", {
      className: "form-signin w-100 m-auto",
      style: {
        display: "block"
      },
      children: [/*#__PURE__*/_jsx("h1", {
        className: "h3 mb-3 fw-normal",
        children: "Please sign in"
      }), /*#__PURE__*/_jsxs("form", {
        onSubmit: submit,
        children: [/*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            type: "email",
            className: "form-control",
            id: "floatingInput",
            placeholder: "name@example.com",
            onChange: e => setUsernameOrEmail(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "floatingInput",
            children: "Email address"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            type: "password",
            className: "form-control",
            id: "floatingPassword",
            placeholder: "Password",
            onChange: e => setPassword(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "floatingPassword",
            children: "Password"
          })]
        }), /*#__PURE__*/_jsx("button", {
          className: "w-100 btn btn-lg btn-primary",
          type: "submit",
          children: "Sign in"
        })]
      }), /*#__PURE__*/_jsx("span", {
        children: "Already have an account? \xA0"
      })]
    })
  });
}
function Register(props) {
  const [email, setEmail] = useState("");
  const [encryptedPassword, setEncryptedPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [dob, setDob] = useState("");
  const [imageURL, setImageURL] = useState("");
  const [username, setUsername] = useState("");
  const [bio, setBio] = useState("");
  const [isPublic, setIsPublic] = useState("");
  const [redirectVar, setRedirectVar] = useState(false);

  // Redirect
  //const navigate = useNavigate();

  const submit = async e => {
    e.preventDefault(); // prevent reload.

    // Create new user as JS object.
    const newUser = {
      email,
      encryptedPassword,
      firstName,
      lastName,
      dob,
      imageURL,
      username,
      bio,
      isPublic
    };
    // Send user data to golang register function.
    const response = await fetch("http://localhost:8080/auth/registration", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(newUser)
    });
    console.log("dob", newUser.dob);
    await response.json();
    // let result = await response.json()
    // if (result.email === email) {
    setRedirectVar(true);
    // }
  };

  // if (redirectVar) {
  // 	return navigate("/login"); // This is still iffy!!! ????????????
  // }

  return /*#__PURE__*/_jsx("div", {
    children: /*#__PURE__*/_jsxs("main", {
      className: "form-signin w-100 m-auto",
      style: {
        display: "block"
      },
      children: [/*#__PURE__*/_jsx("h1", {
        className: "h3 mb-3 fw-normal",
        children: "Please register"
      }), /*#__PURE__*/_jsxs("form", {
        onSubmit: submit,
        children: [/*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            required: true,
            type: "email",
            className: "form-control",
            id: "floatingInput",
            placeholder: "name@example.com",
            onChange: e => setEmail(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "floatingInput",
            children: "Email address"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            required: true,
            type: "password",
            className: "form-control reginput",
            id: "regpassword",
            placeholder: "Password",
            onChange: e => setEncryptedPassword(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "regpassword",
            children: "Password"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            required: true,
            type: "text",
            className: "form-control reginput",
            id: "firstName",
            placeholder: "John",
            onChange: e => setFirstName(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "firstName",
            children: "First Name"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            required: true,
            type: "text",
            className: "form-control reginput",
            id: "lastName",
            placeholder: "Doe",
            onChange: e => setLastName(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "lastName",
            children: "Last Name"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            required: true,
            type: "date",
            className: "form-control reginput",
            id: "dob",
            placeholder: "16/01/1998",
            onChange: e => setDob(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "dob",
            children: "Date of Birth"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            type: "text",
            className: "form-control reginput",
            id: "imageURL",
            placeholder: "https://...",
            onChange: e => setImageURL(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "imageURL",
            children: "ImageURL"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            type: "text",
            className: "form-control reginput",
            id: "username",
            placeholder: "Johnny",
            onChange: e => setUsername(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "username",
            children: "Username"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsxs("div", {
            className: "form-control reginput status",
            children: [/*#__PURE__*/_jsxs("div", {
              children: [/*#__PURE__*/_jsx("input", {
                required: true,
                type: "radio",
                id: "public-status",
                value: "public",
                name: "status",
                checked: true,
                onClick: e => setIsPublic(e.target.value)
              }), /*#__PURE__*/_jsx("label", {
                htmlFor: "public-status",
                children: "Public"
              })]
            }), /*#__PURE__*/_jsxs("div", {
              children: [/*#__PURE__*/_jsx("input", {
                required: true,
                type: "radio",
                id: "private-status",
                value: "private",
                name: "status",
                onClick: e => setIsPublic(e.target.value)
              }), /*#__PURE__*/_jsx("label", {
                htmlFor: "private-status",
                children: "Private"
              })]
            })]
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "",
            children: "Status"
          })]
        }), /*#__PURE__*/_jsxs("div", {
          className: "form-floating",
          children: [/*#__PURE__*/_jsx("input", {
            className: "form-control reginput",
            name: "bio",
            placeholder: "About Me",
            id: "bio",
            cols: "30",
            rows: "10",
            onChange: e => setBio(e.target.value)
          }), /*#__PURE__*/_jsx("label", {
            htmlFor: "about",
            children: "About me"
          })]
        }), /*#__PURE__*/_jsx("button", {
          className: "w-100 btn btn-lg btn-primary",
          type: "submit",
          children: "Register"
        })]
      }), /*#__PURE__*/_jsx("span", {
        children: "Already have an account? \xA0"
      })]
    })
  });
}
function Home(props) {
  return /*#__PURE__*/_jsx("main", {
    children: /*#__PURE__*/_jsx("div", {
      className: "contentContainer",
      children: props.name ? /*#__PURE__*/_jsxs(_Fragment, {
        children: [/*#__PURE__*/_jsx(ProfileImgContainer, {
          name: props.name,
          user: props.user,
          imageURL: props.imageURL
        }), /*#__PURE__*/_jsx(GroupContainer, {
          groups: props.groups,
          socket: props.socket
        }), /*#__PURE__*/_jsx(PostForm, {
          imageURL: props.imageURL
        }), /*#__PURE__*/_jsx(RightSide, {
          openConnection: props.openConnection,
          fetchRequestData: props.fetchRequestData
        }), /*#__PURE__*/_jsx(GetChat, {})]
      }) : /*#__PURE__*/_jsx(_Fragment, {
        children: /*#__PURE__*/_jsx("p", {
          children: "You are not logged in"
        })
      })
    })
  });
}
function Profile(props) {
  const [status, setStatus] = useState("");
  const [privatePosts, setPrivatePosts] = useState([]);

  // Update status to props.user.status.
  // useEffect(() => {
  //   setStatus(props.user.status);
  // }, [props.user.status]);

  const sendStatusToBackend = async data => {
    console.log(data);
    await fetch("http://localhost:8080/update-user-status", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      credentials: "include",
      body: JSON.stringify(data)
    });
  };
  const updateUserStatus = async ev => {
    let buttonClicked = ev.target.getAttribute("data-type");
    if (buttonClicked === "private") {
      sendStatusToBackend({
        user: props.user.email,
        setStatus: "private"
      });
      setStatus("private");
    } else if (buttonClicked === "public") {
      // update on backend if user is not already public
      sendStatusToBackend({
        user: props.user.email,
        setStatus: "public"
      });
      setStatus("public");
    }
  };
  return /*#__PURE__*/_jsxs("div", {
    className: "profileContainer",
    children: ["name=", props.name, "user=", props.user, "imageURL=", props.imageURL, "socket=", props.socket, "currentUser=", props.currentUser, "fetchUsersData=", props.fetchUsersData, "update=", props.update, "setUpdate=", props.setUpdate, /*#__PURE__*/_jsxs("div", {
      className: "formContainer",
      children: [/*#__PURE__*/_jsx("div", {
        className: "smallAvatar",
        children: /*#__PURE__*/_jsx("img", {
          src: props.imageURL,
          alt: "profile photo"
        })
      }), /*#__PURE__*/_jsxs("div", {
        className: "profile-page-title",
        children: [props.name, "'s Posts"]
      })]
    }), props.currentUser === undefined ? /*#__PURE__*/_jsx("div", {
      id: "set-public-private",
      className: "privacyButtons",
      style: {
        width: "100%",
        backgroundColor: "white",
        justifyContent: "space-evenly",
        alignItems: "center"
      },
      children: /*#__PURE__*/_jsxs(_Fragment, {
        children: [/*#__PURE__*/_jsx("button", {
          className: "postType",
          onClick: updateUserStatus,
          "data-type": "private",
          disabled: status === "private" ? true : false,
          style: {
            backgroundColor: status === "private" ? "rgba(129, 25, 41, 0.55)" : "rgb(148, 28, 47)"
          },
          children: "Set Private"
        }), /*#__PURE__*/_jsx("button", {
          className: "postType",
          onClick: updateUserStatus,
          "data-type": "public",
          disabled: status === "public" ? true : false,
          style: {
            backgroundColor: status === "public" ? "rgba(129, 25, 41, 0.55)" : "rgb(148, 28, 47)"
          },
          children: "Set Public"
        })]
      })
    }) : /*#__PURE__*/_jsx("div", {
      id: "set-public-private",
      className: "privacyButtons",
      style: {
        width: "100%",
        backgroundColor: "rgba(250, 250, 250, 0.5)"
      }
    })]
  });
}

async function PublicPosts() {
	// Fetch data from the server
	const fetchData = async () => {
	  const response = await fetch("http://localhost:8080/api/home", {
		method: "GET",
		headers: {
		  "Content-Type": "application/json"
		},
		credentials: "include"
	  });

	  if (!response.ok) {
		throw new Error(`HTTP error! Status: ${response.status}`);
	  }

	  return response.json(); 
	};

	const data = await fetchData();

	// Render the fetched data here
	return (

	  <div className="public-posts">
		<h2>All Users</h2>
		<ul>
		  {data.users.map(user => (
			<li key={user.id}>{user.name}</li>
		  ))}
		</ul>

	  </div>

	);
  }

const root = document.querySelector("#root");
ReactDOM.render( /*#__PURE__*/_jsx(App, {}), root);