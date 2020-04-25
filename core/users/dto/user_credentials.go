package dto

// UserCredentials is a data transfer object used to hold
// a user's credentials.
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
