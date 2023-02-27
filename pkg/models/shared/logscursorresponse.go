package shared

type LogsCursorResponseCursor struct {
	Data     []Log   `json:"data"`
	HasMore  bool    `json:"hasMore"`
	Next     *string `json:"next,omitempty"`
	PageSize int64   `json:"pageSize"`
	Previous *string `json:"previous,omitempty"`
}

type LogsCursorResponse struct {
	Cursor LogsCursorResponseCursor `json:"cursor"`
}
