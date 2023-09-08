// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type AddMetadataToAccountRequest struct {
	// metadata
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	// Exact address of the account. It must match the following regular expressions pattern:
	// ```
	// ^\w+(:\w+)*$
	// ```
	//
	Address string `pathParam:"style=simple,explode=false,name=address"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *AddMetadataToAccountRequest) GetRequestBody() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AddMetadataToAccountRequest) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *AddMetadataToAccountRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type AddMetadataToAccountResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}

func (o *AddMetadataToAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddMetadataToAccountResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *AddMetadataToAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddMetadataToAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
