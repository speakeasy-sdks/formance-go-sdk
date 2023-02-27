package shared

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
