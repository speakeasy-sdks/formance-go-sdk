package shared

import (
	"time"
)

type Wallet struct {
	CreatedAt time.Time              `json:"createdAt"`
	ID        string                 `json:"id"`
	Ledger    string                 `json:"ledger"`
	Metadata  map[string]interface{} `json:"metadata"`
	Name      string                 `json:"name"`
}
