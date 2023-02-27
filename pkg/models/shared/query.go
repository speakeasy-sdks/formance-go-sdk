package shared

type Query struct {
	After    []string `json:"after,omitempty"`
	Cursor   *string  `json:"cursor,omitempty"`
	Ledgers  []string `json:"ledgers,omitempty"`
	PageSize *int64   `json:"pageSize,omitempty"`
	Policy   *string  `json:"policy,omitempty"`
	Sort     *string  `json:"sort,omitempty"`
	Target   *string  `json:"target,omitempty"`
	Terms    []string `json:"terms,omitempty"`
}
