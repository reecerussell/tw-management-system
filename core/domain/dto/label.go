package dto

// Label is a data transfer object used to read data from a database
// or incoming API request bodies.
type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
