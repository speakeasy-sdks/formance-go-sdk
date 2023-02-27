package shared

type AccountsCursorResponseCursor struct {
	Data     []Account `json:"data"`
	HasMore  bool      `json:"hasMore"`
	Next     *string   `json:"next,omitempty"`
	PageSize int64     `json:"pageSize"`
	Previous *string   `json:"previous,omitempty"`
}

type AccountsCursorResponse struct {
	Cursor AccountsCursorResponseCursor `json:"cursor"`
}
