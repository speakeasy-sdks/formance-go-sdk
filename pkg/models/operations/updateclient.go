// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateClientRequest struct {
	UpdateClientRequest *shared.UpdateClientRequest `request:"mediaType=application/json"`
	// Client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

func (o *UpdateClientRequest) GetUpdateClientRequest() *shared.UpdateClientRequest {
	if o == nil {
		return nil
	}
	return o.UpdateClientRequest
}

func (o *UpdateClientRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

type UpdateClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Updated client
	UpdateClientResponse *shared.UpdateClientResponse
}

func (o *UpdateClientResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateClientResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateClientResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateClientResponse) GetUpdateClientResponse() *shared.UpdateClientResponse {
	if o == nil {
		return nil
	}
	return o.UpdateClientResponse
}
