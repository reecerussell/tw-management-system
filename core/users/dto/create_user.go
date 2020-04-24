package dto

// CreateUser is a data-transfer object which holds the required
// data to create a user record.
type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
