CREATE TABLE USERS (
    UserId INTEGER PRIMARY KEY AUTOINCREMENT,
    Bio TEXT,
    CreatedAt BIGINT NOT NULL,
    DOB BIGINT NOT NULL,
    Email TEXT NOT NULL,
    EncryptedPassword TEXT NOT NULL,
    FirstName TEXT NOT NULL,
    ImageURL TEXT,
    IsPublic BIT NOT NULL, 
    LastName TEXT NOT NULL,
    UpdatedAt BIGINT NOT NULL,
    Username TEXT NOT NULL,
);