const template = document.createElement("template");
template.innerHTML = `
<div>
    <h1>Hi</h1>
    <p>What's up?</p>
</div>
`;

export default class Example extends HTMLElement {
    shadowRoot;
    constructor() {
        super();
        this.shadowRoot = this.attachShadow({mode: "open"})
        const content = template.content.cloneNode(true);
        this.shadowRoot.appendChild(content);
    }

    connectedCallback() {
        console.log("Component added to DOM.");
    }

    disconnectedCallback() {
        console.log("Component removed from DOM.");
    }
}