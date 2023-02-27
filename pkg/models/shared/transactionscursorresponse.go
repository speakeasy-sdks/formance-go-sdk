package shared

type TransactionsCursorResponseCursor struct {
	Data     []Transaction `json:"data"`
	HasMore  bool          `json:"hasMore"`
	Next     *string       `json:"next,omitempty"`
	PageSize int64         `json:"pageSize"`
	Previous *string       `json:"previous,omitempty"`
}

type TransactionsCursorResponse struct {
	Cursor TransactionsCursorResponseCursor `json:"cursor"`
}
