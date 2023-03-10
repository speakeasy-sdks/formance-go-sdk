package shared

import (
	"time"
)

type Transaction struct {
	Metadata          map[string]interface{}       `json:"metadata,omitempty"`
	PostCommitVolumes map[string]map[string]Volume `json:"postCommitVolumes,omitempty"`
	Postings          []Posting                    `json:"postings"`
	PreCommitVolumes  map[string]map[string]Volume `json:"preCommitVolumes,omitempty"`
	Reference         *string                      `json:"reference,omitempty"`
	Timestamp         time.Time                    `json:"timestamp"`
	Txid              int64                        `json:"txid"`
}
