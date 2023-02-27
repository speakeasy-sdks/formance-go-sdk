package shared

type User struct {
	Email   *string `json:"email,omitempty"`
	ID      *string `json:"id,omitempty"`
	Subject *string `json:"subject,omitempty"`
}
