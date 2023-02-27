package shared

type PaymentsCursorCursor struct {
	Data     []Payment `json:"data"`
	HasMore  bool      `json:"hasMore"`
	Next     *string   `json:"next,omitempty"`
	PageSize int64     `json:"pageSize"`
	Previous *string   `json:"previous,omitempty"`
}

type PaymentsCursor struct {
	Cursor PaymentsCursorCursor `json:"cursor"`
}
