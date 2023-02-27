package shared

type ListBalancesResponseCursor struct {
	Data     []Balance `json:"data"`
	HasMore  *bool     `json:"hasMore,omitempty"`
	Next     *string   `json:"next,omitempty"`
	PageSize int64     `json:"pageSize"`
	Previous *string   `json:"previous,omitempty"`
}

type ListBalancesResponse struct {
	Cursor ListBalancesResponseCursor `json:"cursor"`
}
