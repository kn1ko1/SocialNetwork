package models

// Most fields ommitted for simplicity of example
// eg. CreatedAt, Categories, etc.
type Post struct {
	ID       int
	AuthorID int
	Title    string
	Contents string
}

func (post Post) Validate() error {
	// All validation logic for a new post instance ommitted
	return nil
}
