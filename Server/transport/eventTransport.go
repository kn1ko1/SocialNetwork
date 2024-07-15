package transport

type EventTransport struct {
	EventId     int    `json:"eventId"`
	Attendance  string `json:"attendance"`
	DateTime    int64  `json:"dateTime"`
	Description string `json:"description"`
	GroupId     int    `json:"groupId"`
	Title       string `json:"title"`
	UserId      int    `json:"userId"`
}
