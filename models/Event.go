package models

type Event struct {
	EventId     int
	CreatedAt   int64
	DateTime    int64
	Description string
	GroupId     int
	Title       string
	UpdatedAt   int64
	UserId      int
}

func (e *Event) Validate() error {
	// Validate logic here
	return nil
}
