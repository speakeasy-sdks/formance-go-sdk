// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Metadata struct {
}

type Account struct {
	Address  string    `json:"address"`
	Metadata *Metadata `json:"metadata,omitempty"`
	Type     *string   `json:"type,omitempty"`
}

func (o *Account) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *Account) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Account) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
