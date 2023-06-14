// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ErrorErrorCode string

const (
	ErrorErrorCodeValidation ErrorErrorCode = "VALIDATION"
)

func (e ErrorErrorCode) ToPointer() *ErrorErrorCode {
	return &e
}

func (e *ErrorErrorCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VALIDATION":
		*e = ErrorErrorCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ErrorErrorCode: %v", v)
	}
}

// Error - General error
type Error struct {
	ErrorCode    ErrorErrorCode `json:"errorCode"`
	ErrorMessage string         `json:"errorMessage"`
}
