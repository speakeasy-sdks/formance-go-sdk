package shared

type CreditWalletRequest struct {
	Amount    Monetary               `json:"amount"`
	Balance   *string                `json:"balance,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Reference *string                `json:"reference,omitempty"`
	Sources   []interface{}          `json:"sources"`
}
