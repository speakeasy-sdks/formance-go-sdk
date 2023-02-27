package shared

type GetTransactionsResponseCursor struct {
	Data     []WalletsTransaction `json:"data"`
	HasMore  *bool                `json:"hasMore,omitempty"`
	Next     *string              `json:"next,omitempty"`
	PageSize int64                `json:"pageSize"`
	Previous *string              `json:"previous,omitempty"`
}

type GetTransactionsResponse struct {
	Cursor GetTransactionsResponseCursor `json:"cursor"`
}
