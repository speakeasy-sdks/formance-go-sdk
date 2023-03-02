package shared

type ExpandedDebitHold struct {
	Description    string                 `json:"description"`
	Destination    interface{}            `json:"destination,omitempty"`
	ID             string                 `json:"id"`
	Metadata       map[string]interface{} `json:"metadata"`
	OriginalAmount int64                  `json:"originalAmount"`
	Remaining      int64                  `json:"remaining"`
	WalletID       string                 `json:"walletID"`
}
