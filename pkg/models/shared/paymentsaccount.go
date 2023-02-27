package shared

import (
	"time"
)

type PaymentsAccountTypeEnum string

const (
	PaymentsAccountTypeEnumTarget PaymentsAccountTypeEnum = "TARGET"
	PaymentsAccountTypeEnumSource PaymentsAccountTypeEnum = "SOURCE"
)

type PaymentsAccount struct {
	CreatedAt time.Time               `json:"createdAt"`
	ID        string                  `json:"id"`
	Provider  ConnectorEnum           `json:"provider"`
	Reference string                  `json:"reference"`
	Type      PaymentsAccountTypeEnum `json:"type"`
}
