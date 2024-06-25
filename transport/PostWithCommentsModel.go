package transport

type PostWithComments struct {
	Comments []CommentTransport `json:"comments"`
	Post     PostTransport      `json:"post"`
}
