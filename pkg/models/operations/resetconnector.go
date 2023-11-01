// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ResetConnectorRequest struct {
	// The name of the connector.
	Connector shared.Connector `pathParam:"style=simple,explode=false,name=connector"`
}

func (o *ResetConnectorRequest) GetConnector() shared.Connector {
	if o == nil {
		return shared.Connector("")
	}
	return o.Connector
}

type ResetConnectorResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ResetConnectorResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ResetConnectorResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ResetConnectorResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
