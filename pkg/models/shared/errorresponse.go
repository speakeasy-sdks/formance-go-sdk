package shared

type ErrorResponse struct {
	Details      *string         `json:"details,omitempty"`
	ErrorCode    *ErrorsEnumEnum `json:"errorCode,omitempty"`
	ErrorMessage *string         `json:"errorMessage,omitempty"`
}
