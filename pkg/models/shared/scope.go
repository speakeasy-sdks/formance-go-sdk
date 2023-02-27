package shared

type Scope struct {
	ID        string                 `json:"id"`
	Label     string                 `json:"label"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Transient []string               `json:"transient,omitempty"`
}
