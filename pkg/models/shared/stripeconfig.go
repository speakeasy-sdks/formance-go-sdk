package shared

type StripeConfig struct {
	APIKey        string  `json:"apiKey"`
	PageSize      *int64  `json:"pageSize,omitempty"`
	PollingPeriod *string `json:"pollingPeriod,omitempty"`
}
