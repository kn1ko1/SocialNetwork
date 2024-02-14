window.addEventListener("load", () => {
    const main = document.getElementById("main");
    const socket = new WebSocket("ws://localhost:8080/ws");
    socket.onopen = () => {
        console.log("Web Socket connection established.");
    }
    socket.onclose = () => {
        console.log("Web Socket connection closed.");
    }
    socket.onerror = (e) => {
        console.log(e.data);
    }
    socket.onmessage = (e) => {
        console.log(e.data);
    }
    const button = document.createElement("button");
    button.type = "button";
    button.innerText = "Click";
    button.addEventListener("click", () => {
        const msg = {
            code: 2,
            body: JSON.stringify({
                name: "Matt",
                age: 27
            })
        }
        socket.send(JSON.stringify(msg))
    })
    main.appendChild(button);
})