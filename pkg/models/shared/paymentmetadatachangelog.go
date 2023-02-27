package shared

import (
	"time"
)

type PaymentMetadataChangelog struct {
	Timestamp time.Time `json:"timestamp"`
	Value     string    `json:"value"`
}
