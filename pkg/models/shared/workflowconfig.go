package shared

type WorkflowConfig struct {
	Stages []map[string]interface{} `json:"stages"`
}
