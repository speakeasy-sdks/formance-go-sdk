package shared

type UpdateScopeRequest struct {
	Label    string                 `json:"label"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}
