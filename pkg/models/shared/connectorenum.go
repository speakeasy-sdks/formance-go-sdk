// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ConnectorEnum - The name of the connector.
type ConnectorEnum string

const (
	ConnectorEnumStripe        ConnectorEnum = "STRIPE"
	ConnectorEnumDummyPay      ConnectorEnum = "DUMMY-PAY"
	ConnectorEnumWise          ConnectorEnum = "WISE"
	ConnectorEnumModulr        ConnectorEnum = "MODULR"
	ConnectorEnumCurrencyCloud ConnectorEnum = "CURRENCY-CLOUD"
	ConnectorEnumBankingCircle ConnectorEnum = "BANKING-CIRCLE"
)

func (e ConnectorEnum) ToPointer() *ConnectorEnum {
	return &e
}

func (e *ConnectorEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "STRIPE":
		fallthrough
	case "DUMMY-PAY":
		fallthrough
	case "WISE":
		fallthrough
	case "MODULR":
		fallthrough
	case "CURRENCY-CLOUD":
		fallthrough
	case "BANKING-CIRCLE":
		*e = ConnectorEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectorEnum: %s", s)
	}
}