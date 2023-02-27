package shared

type Secret struct {
	Clear      string                 `json:"clear"`
	ID         string                 `json:"id"`
	LastDigits string                 `json:"lastDigits"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Name       string                 `json:"name"`
}
