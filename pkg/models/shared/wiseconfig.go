// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WiseConfig struct {
	APIKey string `json:"apiKey"`
}

func (o *WiseConfig) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
