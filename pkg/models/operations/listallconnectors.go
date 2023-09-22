// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListAllConnectorsResponse struct {
	// OK
	ConnectorsResponse *shared.ConnectorsResponse
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
}

func (o *ListAllConnectorsResponse) GetConnectorsResponse() *shared.ConnectorsResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorsResponse
}

func (o *ListAllConnectorsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAllConnectorsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAllConnectorsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
