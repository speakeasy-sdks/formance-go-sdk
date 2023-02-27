package shared

type WalletsErrorResponseErrorCodeEnum string

const (
	WalletsErrorResponseErrorCodeEnumValidation WalletsErrorResponseErrorCodeEnum = "VALIDATION"
)

type WalletsErrorResponse struct {
	ErrorCode    WalletsErrorResponseErrorCodeEnum `json:"errorCode"`
	ErrorMessage string                            `json:"errorMessage"`
}
