package shared

type CreateSecretRequest struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name     string                 `json:"name"`
}
