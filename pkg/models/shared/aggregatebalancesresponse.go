// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AggregateBalancesResponse - OK
type AggregateBalancesResponse struct {
	Data map[string]int64 `json:"data"`
}

func (o *AggregateBalancesResponse) GetData() map[string]int64 {
	if o == nil {
		return map[string]int64{}
	}
	return o.Data
}
