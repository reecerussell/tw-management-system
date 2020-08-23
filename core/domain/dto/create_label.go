package dto

// CreateLabel is used to read request bodies of API request to
// create a new label record.
type CreateLabel struct {
	Name string `json:"name"`
}
