// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type RunWorkflowRequest struct {
	RequestBody map[string]string `request:"mediaType=application/json"`
	// The flow id
	FlowID string `pathParam:"style=simple,explode=false,name=flowId"`
	// Wait end of the workflow before return
	Wait *bool `queryParam:"style=form,explode=true,name=wait"`
}

func (o *RunWorkflowRequest) GetRequestBody() map[string]string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *RunWorkflowRequest) GetFlowID() string {
	if o == nil {
		return ""
	}
	return o.FlowID
}

func (o *RunWorkflowRequest) GetWait() *bool {
	if o == nil {
		return nil
	}
	return o.Wait
}

type RunWorkflowResponse struct {
	ContentType string
	// General error
	Error *shared.Error
	// The workflow occurrence
	RunWorkflowResponse *shared.RunWorkflowResponse
	StatusCode          int
	RawResponse         *http.Response
}

func (o *RunWorkflowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RunWorkflowResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *RunWorkflowResponse) GetRunWorkflowResponse() *shared.RunWorkflowResponse {
	if o == nil {
		return nil
	}
	return o.RunWorkflowResponse
}

func (o *RunWorkflowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RunWorkflowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
