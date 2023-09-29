// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type GetManyConfigsRequest struct {
	// Optional filter by endpoint URL
	Endpoint *string `queryParam:"style=form,explode=true,name=endpoint"`
	// Optional filter by Config ID
	ID *string `queryParam:"style=form,explode=true,name=id"`
}

func (o *GetManyConfigsRequest) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *GetManyConfigsRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type GetManyConfigsResponse struct {
	// OK
	ConfigsResponse *shared.ConfigsResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetManyConfigsResponse) GetConfigsResponse() *shared.ConfigsResponse {
	if o == nil {
		return nil
	}
	return o.ConfigsResponse
}

func (o *GetManyConfigsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetManyConfigsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetManyConfigsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
