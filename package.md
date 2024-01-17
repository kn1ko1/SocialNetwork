# package info

- [ ] `Models`

  - Core structure definitions (e.g. User, Post, Comment, Groups, Notification, Event etc)
  - Receiver pattern allows for self-contained model validation

- [ ] `SQLite`

  - Implementation of DB Create, Read, Update and Delete operations (CRUD)
    `ALL functions should accept as arguments a *sql.DB instance`

- [ ] `Repo`
  - Interface definitions for each table e.g. Post Repo interface with method declarations for CreatePost, DeletePost etc
  - Implementations of repo interfaces specifically for SQLite in this project
- [ ] `Auth`

  - Definitions for functions for generating a session ID; tracking a logged in User associated with a session ID; removing sessionIDs upon logout/expiring

- [ ] `Websockets`

  - Definition of structure for Web Socket Handler responsible for upgrading an HTTP connection to a WS connection
  - Definition of a Web Socket Client structure, abstracting Client connections, responsible for Read loop.
  - Definition of a Pool Structure, responsible for handling a number of connected clients (Switch case to handle different types of message)
  - Socket Message structure definitions, capable of representing in Go, the JSON payload sent over the WebSocket connection

- [ ] API
  - Handler definitions which take a repo interface as argument `each handler should be a struct which implements http.Handler interface`
