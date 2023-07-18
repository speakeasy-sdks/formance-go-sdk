// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Wallet struct {
	CreatedAt time.Time `json:"createdAt"`
	// The unique ID of the wallet.
	ID     string `json:"id"`
	Ledger string `json:"ledger"`
	// Metadata associated with the wallet.
	Metadata map[string]interface{} `json:"metadata"`
	Name     string                 `json:"name"`
}

func (o *Wallet) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Wallet) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Wallet) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *Wallet) GetMetadata() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Metadata
}

func (o *Wallet) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
