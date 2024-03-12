import Example from "./components/Example.js";

customElements.define("example-component", Example);

window.addEventListener("load", () => {
    const example = new Example();
    const main = document.getElementById("main");
    main.appendChild(example);
})