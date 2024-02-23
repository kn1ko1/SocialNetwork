window.addEventListener("load", () => {
    renderHomePage();
});

function renderHomePage() {
    const main = document.getElementById("main");
    const h1 = document.createElement("h1");
    h1.innerText = "Home";
    main.appendChild(h1);
    const button = document.createElement("button");
    button.innerText = "Other Page";
    button.type = "click";
    button.addEventListener("click", () => {
        renderOtherPage();
    })
}

function renderOtherPage() {
    const main = document.getElementById("main");
    const h1 = document.createElement("h1");
    h1.innerText = "Other";
    main.appendChild(h1);
    const button = document.createElement("button");
    button.innerText = "Home Page";
    button.type = "click";
    button.addEventListener("click", () => {
        renderHomePage();
    })
}