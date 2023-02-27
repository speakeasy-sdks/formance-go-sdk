package shared

import (
	"time"
)

type Workflow struct {
	Config    WorkflowConfig `json:"config"`
	CreatedAt time.Time      `json:"createdAt"`
	ID        string         `json:"id"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
