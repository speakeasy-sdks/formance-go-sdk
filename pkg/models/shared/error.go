// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ErrorErrorCodeEnum string

const (
	ErrorErrorCodeEnumValidation ErrorErrorCodeEnum = "VALIDATION"
)

func (e ErrorErrorCodeEnum) ToPointer() *ErrorErrorCodeEnum {
	return &e
}

func (e *ErrorErrorCodeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VALIDATION":
		*e = ErrorErrorCodeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ErrorErrorCodeEnum: %v", v)
	}
}

// Error - General error
type Error struct {
	ErrorCode    ErrorErrorCodeEnum `json:"errorCode"`
	ErrorMessage string             `json:"errorMessage"`
}
