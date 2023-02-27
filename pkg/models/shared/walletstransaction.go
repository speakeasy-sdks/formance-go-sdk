package shared

import (
	"time"
)

type WalletsTransaction struct {
	Metadata          map[string]interface{}              `json:"metadata,omitempty"`
	PostCommitVolumes map[string]map[string]WalletsVolume `json:"postCommitVolumes,omitempty"`
	Postings          []Posting                           `json:"postings"`
	PreCommitVolumes  map[string]map[string]WalletsVolume `json:"preCommitVolumes,omitempty"`
	Reference         *string                             `json:"reference,omitempty"`
	Timestamp         time.Time                           `json:"timestamp"`
	Txid              int64                               `json:"txid"`
}
