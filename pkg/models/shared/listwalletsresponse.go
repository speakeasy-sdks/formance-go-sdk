package shared

type ListWalletsResponseCursor struct {
	Data     []Wallet `json:"data"`
	HasMore  *bool    `json:"hasMore,omitempty"`
	Next     *string  `json:"next,omitempty"`
	PageSize int64    `json:"pageSize"`
	Previous *string  `json:"previous,omitempty"`
}

type ListWalletsResponse struct {
	Cursor ListWalletsResponseCursor `json:"cursor"`
}
