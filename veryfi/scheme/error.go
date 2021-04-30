package scheme

// Error describes an error response.
type Error struct {
	Status  string `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
