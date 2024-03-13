const loginTemplate = document.createElement("template");
loginTemplate.innerHTML = `
<div>
    <h1>Hi</h1>
    <p>What's up?</p>
</div>
`;

export default class Login extends HTMLElement {
    shadowRoot;
    constructor() {
        super();
        this.shadowRoot = this.attachShadow({mode: "open"})
        const content = loginTemplate.content.cloneNode(true);
        this.shadowRoot.appendChild(content);
    }

    connectedCallback() {
        console.log("Component added to DOM.");
    }

    disconnectedCallback() {
        console.log("Component removed from DOM.");
    }
}


export function Login(props) {
    const [usernameOrEmail, setUsernameOrEmail] = useState("");
    const [password, setPassword] = useState("");
    // const [redirectVar, setRedirectVar] = useState(false);
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
      // setRedirectVar(true);
      props.setName(validUser.first);
    };
    return /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("main", {
      className: "form-signin w-100 m-auto",
      style: {
        display: "block"
      }
    }, /*#__PURE__*/React.createElement("h1", {
      className: "h3 mb-3 fw-normal"
    }, "Please sign in"), /*#__PURE__*/React.createElement("form", {
      onSubmit: submit
    }, /*#__PURE__*/React.createElement("div", {
      className: "form-floating"
    }, /*#__PURE__*/React.createElement("input", {
      type: "email",
      className: "form-control",
      id: "floatingInput",
      placeholder: "name@example.com",
      onChange: e => setUsernameOrEmail(e.target.value)
    }), /*#__PURE__*/React.createElement("label", {
      htmlFor: "floatingInput"
    }, "Email address")), /*#__PURE__*/React.createElement("div", {
      className: "form-floating"
    }, /*#__PURE__*/React.createElement("input", {
      type: "password",
      className: "form-control",
      id: "floatingPassword",
      placeholder: "Password",
      onChange: e => setPassword(e.target.value)
    }), /*#__PURE__*/React.createElement("label", {
      htmlFor: "floatingPassword"
    }, "Password")), /*#__PURE__*/React.createElement("button", {
      className: "w-100 btn btn-lg btn-primary",
      type: "submit"
    }, "Sign in")), /*#__PURE__*/React.createElement("span", null, "Already have an account? \xA0")));
  }