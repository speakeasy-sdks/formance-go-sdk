package shared

type CurrencyCloudConfig struct {
	APIKey        string  `json:"apiKey"`
	Endpoint      *string `json:"endpoint,omitempty"`
	LoginID       string  `json:"loginID"`
	PollingPeriod *string `json:"pollingPeriod,omitempty"`
}
