// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type StripeConfig struct {
	APIKey string `json:"apiKey"`
	// Number of BalanceTransaction to fetch at each polling interval.
	//
	PageSize *int64 `json:"pageSize,omitempty"`
	// The frequency at which the connector will try to fetch new BalanceTransaction objects from Stripe API.
	//
	PollingPeriod *string `json:"pollingPeriod,omitempty"`
}

func (o *StripeConfig) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *StripeConfig) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *StripeConfig) GetPollingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.PollingPeriod
}
