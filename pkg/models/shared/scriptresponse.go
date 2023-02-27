package shared

type ScriptResponse struct {
	Details      *string         `json:"details,omitempty"`
	ErrorCode    *ErrorsEnumEnum `json:"errorCode,omitempty"`
	ErrorMessage *string         `json:"errorMessage,omitempty"`
	Transaction  *Transaction    `json:"transaction,omitempty"`
}
