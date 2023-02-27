package shared

type ErrorErrorCodeEnum string

const (
	ErrorErrorCodeEnumValidation ErrorErrorCodeEnum = "VALIDATION"
)

type Error struct {
	ErrorCode    ErrorErrorCodeEnum `json:"errorCode"`
	ErrorMessage string             `json:"errorMessage"`
}
