package datamodel

// QueueBuster is a data model object used to read and write
// to the queue buster database table.
type QueueBuster struct {
	Department string `json:"department"`
	Status     string `json:"status"`
	Announcements bool `json:"announcements"`
}
