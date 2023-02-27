package shared

type PaymentMetadata struct {
	Changelog *PaymentMetadataChangelog `json:"changelog,omitempty"`
	Key       string                    `json:"key"`
	Value     string                    `json:"value"`
}
