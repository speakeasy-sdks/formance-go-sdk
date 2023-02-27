package shared

import (
	"time"
)

type WorkflowOccurrence struct {
	CreatedAt  time.Time     `json:"createdAt"`
	ID         string        `json:"id"`
	Statuses   []StageStatus `json:"statuses"`
	UpdatedAt  time.Time     `json:"updatedAt"`
	WorkflowID string        `json:"workflowID"`
}
