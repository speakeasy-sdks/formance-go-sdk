package shared

type Hold struct {
	Description string                 `json:"description"`
	Destination interface{}            `json:"destination,omitempty"`
	ID          string                 `json:"id"`
	Metadata    map[string]interface{} `json:"metadata"`
	WalletID    string                 `json:"walletID"`
}
