package shared

type TasksCursorCursor struct {
	Data     []interface{} `json:"data"`
	HasMore  bool          `json:"hasMore"`
	Next     *string       `json:"next,omitempty"`
	PageSize int64         `json:"pageSize"`
	Previous *string       `json:"previous,omitempty"`
}

type TasksCursor struct {
	Cursor TasksCursorCursor `json:"cursor"`
}
