CREATE TABLE POST_USERS (
    PostUserId INTEGER PRIMARY KEY AUTOINCREMENT,
    CreatedAt BIGINT NOT NULL,
    PostId INTEGER NOT NULL,
    UpdatedAt BIGINT NOT NULL,
    UserId INTEGER NOT NULL
);
