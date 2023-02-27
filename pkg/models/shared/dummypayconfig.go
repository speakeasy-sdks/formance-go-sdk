package shared

type DummyPayConfig struct {
	Directory            string  `json:"directory"`
	FileGenerationPeriod *string `json:"fileGenerationPeriod,omitempty"`
	FilePollingPeriod    *string `json:"filePollingPeriod,omitempty"`
}
