package shared

type Volume struct {
	Balance *int64 `json:"balance,omitempty"`
	Input   int64  `json:"input"`
	Output  int64  `json:"output"`
}
