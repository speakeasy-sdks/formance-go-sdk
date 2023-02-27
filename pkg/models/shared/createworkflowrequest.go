package shared

type CreateWorkflowRequest struct {
	Stages []map[string]interface{} `json:"stages"`
}
