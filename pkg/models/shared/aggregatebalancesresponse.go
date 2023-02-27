package shared

type AggregateBalancesResponse struct {
	Data map[string]int64 `json:"data"`
}
