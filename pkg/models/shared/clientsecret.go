package shared

type ClientSecret struct {
	ID         string                 `json:"id"`
	LastDigits string                 `json:"lastDigits"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Name       string                 `json:"name"`
}
