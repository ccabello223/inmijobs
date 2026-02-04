package dto

type Ping struct {
	Message       string `json:"message"`
	Authenticated bool   `json:"authenticated,omitempty"`
	UserID        string `json:"userId,omitempty"`
}
