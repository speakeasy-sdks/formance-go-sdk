package shared

type BalanceWithAssets struct {
	Assets map[string]float64 `json:"assets"`
	Name   string             `json:"name"`
}
