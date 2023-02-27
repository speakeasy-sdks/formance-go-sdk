package shared

import (
	"time"
)

type PostTransactionScript struct {
	Plain string                 `json:"plain"`
	Vars  map[string]interface{} `json:"vars,omitempty"`
}

type PostTransaction struct {
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Postings  []Posting              `json:"postings,omitempty"`
	Reference *string                `json:"reference,omitempty"`
	Script    *PostTransactionScript `json:"script,omitempty"`
	Timestamp *time.Time             `json:"timestamp,omitempty"`
}
