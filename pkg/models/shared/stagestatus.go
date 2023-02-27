package shared

import (
	"time"
)

type StageStatus struct {
	Error        *string    `json:"error,omitempty"`
	OccurrenceID string     `json:"occurrenceID"`
	Stage        float64    `json:"stage"`
	StartedAt    time.Time  `json:"startedAt"`
	TerminatedAt *time.Time `json:"terminatedAt,omitempty"`
}
