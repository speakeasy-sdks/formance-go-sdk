package shared

type DebitWalletRequest struct {
	Amount      Monetary               `json:"amount"`
	Balances    []string               `json:"balances,omitempty"`
	Description *string                `json:"description,omitempty"`
	Destination *interface{}           `json:"destination,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Pending     *bool                  `json:"pending,omitempty"`
}
