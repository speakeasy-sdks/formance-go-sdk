package shared

type GetHoldsResponseCursor struct {
	Data     []Hold  `json:"data"`
	HasMore  *bool   `json:"hasMore,omitempty"`
	Next     *string `json:"next,omitempty"`
	PageSize int64   `json:"pageSize"`
	Previous *string `json:"previous,omitempty"`
}

type GetHoldsResponse struct {
	Cursor GetHoldsResponseCursor `json:"cursor"`
}
