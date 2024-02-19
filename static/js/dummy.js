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


    const loginButton = document.getElementById("loginSubmit");
    loginButton.addEventListener("click", () => {
        const msg = {
            code: 2,
            body: JSON.stringify({
                name: "Rupert",
                age: 36
            })
        }
        socket.send(JSON.stringify(msg))
    })


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

    const eventButton = document.createElement("button");
    eventButton.type = "button";
    eventButton.innerText = "Event Me!";
    eventButton.addEventListener("click", () => {
        const event = {
            code: 3,
            body: JSON.stringify({
                eventId: 1337,
                createdAt: 36,
                dateTime: 37,
                description: "This is a dummy event using websockets",
                groupId: 2,
                title: "DUMMY",
                updatedAt: 36,
                userId: 1
            })
        }
        socket.send(JSON.stringify(event))
    })
    main.appendChild(eventButton);
})