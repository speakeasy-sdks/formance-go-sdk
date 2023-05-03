// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type MigrationInfoStateEnum string

const (
	MigrationInfoStateEnumToDo MigrationInfoStateEnum = "to do"
	MigrationInfoStateEnumDone MigrationInfoStateEnum = "done"
)

func (e MigrationInfoStateEnum) ToPointer() *MigrationInfoStateEnum {
	return &e
}

func (e *MigrationInfoStateEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "to do":
		fallthrough
	case "done":
		*e = MigrationInfoStateEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MigrationInfoStateEnum: %v", v)
	}
}

type MigrationInfo struct {
	Date    *time.Time              `json:"date,omitempty"`
	Name    *string                 `json:"name,omitempty"`
	State   *MigrationInfoStateEnum `json:"state,omitempty"`
	Version *int64                  `json:"version,omitempty"`
}
