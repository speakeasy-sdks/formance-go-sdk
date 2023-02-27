package shared

type Contract struct {
	Account *string                `json:"account,omitempty"`
	Expr    map[string]interface{} `json:"expr"`
}
