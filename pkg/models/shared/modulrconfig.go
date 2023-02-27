package shared

type ModulrConfig struct {
	APIKey    string  `json:"apiKey"`
	APISecret string  `json:"apiSecret"`
	Endpoint  *string `json:"endpoint,omitempty"`
}
