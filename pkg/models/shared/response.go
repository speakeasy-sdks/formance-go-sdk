package shared

type ResponseCursorTotal struct {
	Relation *string `json:"relation,omitempty"`
	Value    *int64  `json:"value,omitempty"`
}

type ResponseCursor struct {
	Data     []map[string]interface{} `json:"data,omitempty"`
	HasMore  *bool                    `json:"hasMore,omitempty"`
	Next     *string                  `json:"next,omitempty"`
	PageSize *int64                   `json:"pageSize,omitempty"`
	Previous *string                  `json:"previous,omitempty"`
	Total    *ResponseCursorTotal     `json:"total,omitempty"`
}

type Response struct {
	Cursor *ResponseCursor        `json:"cursor,omitempty"`
	Data   map[string]interface{} `json:"data,omitempty"`
}
