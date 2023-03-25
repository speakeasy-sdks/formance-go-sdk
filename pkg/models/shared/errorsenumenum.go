// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ErrorsEnumEnum string

const (
	ErrorsEnumEnumInternal          ErrorsEnumEnum = "INTERNAL"
	ErrorsEnumEnumInsufficientFund  ErrorsEnumEnum = "INSUFFICIENT_FUND"
	ErrorsEnumEnumValidation        ErrorsEnumEnum = "VALIDATION"
	ErrorsEnumEnumConflict          ErrorsEnumEnum = "CONFLICT"
	ErrorsEnumEnumNoScript          ErrorsEnumEnum = "NO_SCRIPT"
	ErrorsEnumEnumCompilationFailed ErrorsEnumEnum = "COMPILATION_FAILED"
	ErrorsEnumEnumMetadataOverride  ErrorsEnumEnum = "METADATA_OVERRIDE"
)

func (e *ErrorsEnumEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "INTERNAL":
		fallthrough
	case "INSUFFICIENT_FUND":
		fallthrough
	case "VALIDATION":
		fallthrough
	case "CONFLICT":
		fallthrough
	case "NO_SCRIPT":
		fallthrough
	case "COMPILATION_FAILED":
		fallthrough
	case "METADATA_OVERRIDE":
		*e = ErrorsEnumEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ErrorsEnumEnum: %s", s)
	}
}
