// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetWorkflowOccurrenceResponse - The workflow occurrence
type GetWorkflowOccurrenceResponse struct {
	Data WorkflowOccurrence `json:"data"`
}

func (o *GetWorkflowOccurrenceResponse) GetData() WorkflowOccurrence {
	if o == nil {
		return WorkflowOccurrence{}
	}
	return o.Data
}
