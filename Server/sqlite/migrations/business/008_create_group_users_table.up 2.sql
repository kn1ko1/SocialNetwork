CREATE TABLE IF NOT EXISTS GROUP_USERS (
GroupUserId INTEGER PRIMARY KEY AUTOINCREMENT,
CreatedAt BIGINT NOT NULL,
GroupId INTEGER NOT NULL, 
UpdatedAt BIGINT NOT NULL,
UserId INTEGER NOT NULL,
FOREIGN KEY (GroupId) REFERENCES GROUPS(GroupId),
FOREIGN KEY (UserId) REFERENCES USERS(UserId)
);