// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type WebhooksConfig struct {
	Active     *bool      `json:"active,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	Endpoint   *string    `json:"endpoint,omitempty"`
	EventTypes []string   `json:"eventTypes,omitempty"`
	ID         *string    `json:"id,omitempty"`
	Secret     *string    `json:"secret,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}
