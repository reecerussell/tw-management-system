package dto

// ChangePassword is a data-transfer objects which holds the
// data needed to change a user's password.
type ChangePassword struct {
	ID              string `json:"id"`
	CurrentPassword string `json:"current"`
	NewPassword     string `json:"new"`
}
