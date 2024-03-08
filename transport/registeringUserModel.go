package transport
type RegisteringUser struct {
	UserId            int    `json:"userId"`
	Bio               string `json:"bio"`
	CreatedAt         int64
	DOB               string  `json:"dob"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encryptedPassword"`
	FirstName         string `json:"firstName"`
	ImageURL          string `json:"imageURL"`
	IsPublic          bool   `json:"isPublic"`
	LastName          string `json:"lastName"`
	UpdatedAt         int64
	Username          string `json:"username"`
}