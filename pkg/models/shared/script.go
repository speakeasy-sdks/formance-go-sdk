package shared

type Script struct {
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Plain     string                 `json:"plain"`
	Reference *string                `json:"reference,omitempty"`
	Vars      map[string]interface{} `json:"vars,omitempty"`
}
