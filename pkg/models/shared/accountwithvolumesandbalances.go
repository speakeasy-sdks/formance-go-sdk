package shared

type AccountWithVolumesAndBalances struct {
	Address  string                      `json:"address"`
	Balances map[string]int64            `json:"balances,omitempty"`
	Metadata map[string]interface{}      `json:"metadata,omitempty"`
	Type     *string                     `json:"type,omitempty"`
	Volumes  map[string]map[string]int64 `json:"volumes,omitempty"`
}
