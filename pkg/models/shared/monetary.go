package shared

type Monetary struct {
	Amount int64  `json:"amount"`
	Asset  string `json:"asset"`
}
