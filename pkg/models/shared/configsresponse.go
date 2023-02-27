package shared

type ConfigsResponseCursor struct {
	Data     []interface{} `json:"data"`
	HasMore  *bool         `json:"hasMore,omitempty"`
	Next     *string       `json:"next,omitempty"`
	PageSize int64         `json:"pageSize"`
	Previous *string       `json:"previous,omitempty"`
}

type ConfigsResponse struct {
	Cursor ConfigsResponseCursor `json:"cursor"`
}
