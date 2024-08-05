package models

import "time"

// Test email addresses
var (
	sutBio  = []string{"string 1", "string 2", ""}
	sutBody = []string{"Hello, World", "Test", "Example"}
	sutDOB  = []int64{int64((time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC)).Unix()),
		int64((time.Date(1975, time.May, 15, 0, 0, 0, 0, time.UTC)).Unix()),
		int64((time.Date(1990, time.October, 30, 0, 0, 0, 0, time.UTC)).Unix())}
	sutEmail              = []string{"john.doe@example.com", "another@email.co", "user@123.de"}
	sutFirstName          = []string{"Matthew", "Micheal", "Nikoi"}
	sutIsPublic           = []bool{true, false}
	sutIsGoing            = []bool{true, false}
	sutLastName           = []string{"Cheetham", "Cornea", "Fenton"}
	sutPassword           = []string{"Password123", "Password456", "Password789"}
	sutUsername           = []string{"User123", "User456", "User789"}
	sutDescriptions       = []string{"Event 1", "Event 2", "Event 3"}
	sutTitles             = []string{"Title 1", "Title 2", "Title 3"}
	sutTitle              = []string{"Family Time", "Our Group", "Cinema Trip"}
	sutMessageTypes       = []string{"GC", "DM"}
	sutImageURL           = []string{"URL1", "URL2", "URL3", ""}
	sutPrivacy            = []string{"public", "private", "almost private"}
	sutNotificationType   = []string{"Follow Request", "Group Invitation", "Message"}
	sutNotificationStatus = []string{"Pending", "Accepted", "Rejected"}
)
