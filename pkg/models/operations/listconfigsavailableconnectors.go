// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type ListConfigsAvailableConnectorsResponse struct {
	// OK
	ConnectorsConfigsResponse *shared.ConnectorsConfigsResponse
	ContentType               string
	StatusCode                int
	RawResponse               *http.Response
}

func (o *ListConfigsAvailableConnectorsResponse) GetConnectorsConfigsResponse() *shared.ConnectorsConfigsResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorsConfigsResponse
}

func (o *ListConfigsAvailableConnectorsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConfigsAvailableConnectorsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConfigsAvailableConnectorsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
