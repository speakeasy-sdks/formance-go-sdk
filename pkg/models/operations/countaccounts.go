// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

// Metadata - Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
type Metadata struct {
}

type CountAccountsRequest struct {
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter accounts by address pattern (regular expression placed between ^ and $).
	Address *string `queryParam:"style=form,explode=true,name=address"`
	// Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata *Metadata `queryParam:"style=deepObject,explode=true,name=metadata"`
}

func (o *CountAccountsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *CountAccountsRequest) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CountAccountsRequest) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type CountAccountsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CountAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CountAccountsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *CountAccountsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CountAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CountAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
