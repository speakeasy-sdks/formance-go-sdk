// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Mapping struct {
	Contracts []Contract `json:"contracts"`
}

func (o *Mapping) GetContracts() []Contract {
	if o == nil {
		return []Contract{}
	}
	return o.Contracts
}
