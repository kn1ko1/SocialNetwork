# Social Network

## API Endpoints

- /api/comments - POST (create a new comment)
- /api/comments/{commentId} - GET (returns a single comment by Comment ID)

- /api/events - POST (create a new event)
- /api/events/{eventId} - GET (returns a single event by Event ID)

- /api/groups - GET (returns a list of all groups); POST (create a new group)
- /api/groups/{groupId} - GET (returns a single group by Group ID)
- /api/groups/{groupId}/events - GET (returns a list of events by Group ID)
- /api/groups/{groupId}/posts - GET (returns a list of posts by Group ID)
- /api/groups/{groupId}/users - GET (returns a list of members by Group ID)

- /api/posts/{postId} - GET (returns a post by Post ID)
- /api/posts/{postId}/comments - GET (returns a list of comments by Post ID)

- /api/users - GET (returns a list of all users)
- /api/users/{userId} - GET (returns a single user by User ID)
- /api/users/{userId}/followers - GET (returns a list of followers by User ID)
- /api/users/{userId}/groups - GET (returns a list of Groups by UserID - utilise GroupUser table)
- /api/users/{userId}/notifications - GET (returns a list of notifications by User ID)

## Locked API Endpoints

- /api/comments - GET (returns a list of all comments)
- /api/events - GET (returns a list of all events)
- /api/messages - GET (returns a list of all messages)
- /api/notifications - GET (returns a list of all notifications)
- /api/posts - GET (returns a list of all posts)
