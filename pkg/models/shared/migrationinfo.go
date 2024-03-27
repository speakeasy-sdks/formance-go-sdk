// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/utils"
	"time"
)

type State string

const (
	StateToDo State = "to do"
	StateDone State = "done"
)

func (e State) ToPointer() *State {
	return &e
}

func (e *State) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "to do":
		fallthrough
	case "done":
		*e = State(v)
		return nil
	default:
		return fmt.Errorf("invalid value for State: %v", v)
	}
}

type MigrationInfo struct {
	Date    *time.Time `json:"date,omitempty"`
	Name    *string    `json:"name,omitempty"`
	State   *State     `json:"state,omitempty"`
	Version *int64     `json:"version,omitempty"`
}

func (m MigrationInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MigrationInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MigrationInfo) GetDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *MigrationInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MigrationInfo) GetState() *State {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *MigrationInfo) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
