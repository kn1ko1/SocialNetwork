CREATE TABLE NOTIFICATIONS (
    NotificationId   INTEGER PRIMARY KEY AUTOINCREMENT,
	CreatedAt        BIGINT NOT NULL,
	NotificationType TEXT NOT NULL,
	ObjectId         INTEGER,
	SenderId         INTEGER NOT NULL,  
	Status           TEXT NOT NULL,
	TargetId         INTEGER NOT NULL,
	UpdatedAt        BIGINT NOT NULL
);