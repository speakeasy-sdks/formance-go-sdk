package shared

import (
	"time"
)

type WalletWithBalancesBalances struct {
	Main AssetHolder `json:"main"`
}

type WalletWithBalances struct {
	Balances  WalletWithBalancesBalances `json:"balances"`
	CreatedAt time.Time                  `json:"createdAt"`
	ID        string                     `json:"id"`
	Ledger    string                     `json:"ledger"`
	Metadata  map[string]interface{}     `json:"metadata"`
	Name      string                     `json:"name"`
}
