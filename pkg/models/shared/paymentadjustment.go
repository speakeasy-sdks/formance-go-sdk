package shared

import (
	"time"
)

type PaymentAdjustment struct {
	Absolute bool                   `json:"absolute"`
	Amount   int64                  `json:"amount"`
	Date     time.Time              `json:"date"`
	Raw      map[string]interface{} `json:"raw"`
	Status   PaymentStatusEnum      `json:"status"`
}
