// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type PaymentsAccountTypeEnum string

const (
	PaymentsAccountTypeEnumTarget PaymentsAccountTypeEnum = "TARGET"
	PaymentsAccountTypeEnumSource PaymentsAccountTypeEnum = "SOURCE"
)

func (e PaymentsAccountTypeEnum) ToPointer() *PaymentsAccountTypeEnum {
	return &e
}

func (e *PaymentsAccountTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TARGET":
		fallthrough
	case "SOURCE":
		*e = PaymentsAccountTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentsAccountTypeEnum: %v", v)
	}
}

type PaymentsAccount struct {
	CreatedAt time.Time               `json:"createdAt"`
	ID        string                  `json:"id"`
	Provider  ConnectorEnum           `json:"provider"`
	Reference string                  `json:"reference"`
	Type      PaymentsAccountTypeEnum `json:"type"`
}
