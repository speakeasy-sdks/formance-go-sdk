// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/utils"
	"time"
)

type TaskWiseDescriptor struct {
	Key       *string `json:"key,omitempty"`
	Name      *string `json:"name,omitempty"`
	ProfileID *int64  `json:"profileID,omitempty"`
}

func (o *TaskWiseDescriptor) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *TaskWiseDescriptor) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaskWiseDescriptor) GetProfileID() *int64 {
	if o == nil {
		return nil
	}
	return o.ProfileID
}

type TaskWiseState struct {
}

type TaskWise struct {
	ConnectorID string             `json:"connectorId"`
	CreatedAt   time.Time          `json:"createdAt"`
	Descriptor  TaskWiseDescriptor `json:"descriptor"`
	Error       *string            `json:"error,omitempty"`
	ID          string             `json:"id"`
	State       TaskWiseState      `json:"state"`
	Status      PaymentStatus      `json:"status"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func (t TaskWise) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskWise) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *TaskWise) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *TaskWise) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TaskWise) GetDescriptor() TaskWiseDescriptor {
	if o == nil {
		return TaskWiseDescriptor{}
	}
	return o.Descriptor
}

func (o *TaskWise) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *TaskWise) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TaskWise) GetState() TaskWiseState {
	if o == nil {
		return TaskWiseState{}
	}
	return o.State
}

func (o *TaskWise) GetStatus() PaymentStatus {
	if o == nil {
		return PaymentStatus("")
	}
	return o.Status
}

func (o *TaskWise) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
