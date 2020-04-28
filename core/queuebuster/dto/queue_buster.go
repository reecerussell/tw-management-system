package dto

// QueueBuster is a data transfer object used to pass records
// in and out of the application.
type QueueBuster struct {
	Department string `json:"department"`
	Enabled    bool   `json:"enabled"`
}
