package shared

type StripeTransferRequest struct {
	Amount      *int64                 `json:"amount,omitempty"`
	Asset       *string                `json:"asset,omitempty"`
	Destination *string                `json:"destination,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}
