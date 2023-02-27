package shared

type CreateWalletRequest struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name     string                 `json:"name"`
}
