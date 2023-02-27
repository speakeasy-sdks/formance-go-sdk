package shared

type ConfirmHoldRequest struct {
	Amount *int64 `json:"amount,omitempty"`
	Final  *bool  `json:"final,omitempty"`
}
