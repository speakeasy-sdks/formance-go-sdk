// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateScopeResponse struct {
	Data *Scope `json:"data,omitempty"`
}

func (o *CreateScopeResponse) GetData() *Scope {
	if o == nil {
		return nil
	}
	return o.Data
}
