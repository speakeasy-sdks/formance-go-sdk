// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateWalletRequestBody struct {
	// Custom metadata to attach to this wallet.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (o *UpdateWalletRequestBody) GetMetadata() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type UpdateWalletRequest struct {
	RequestBody *UpdateWalletRequestBody `request:"mediaType=application/json"`
	ID          string                   `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateWalletRequest) GetRequestBody() *UpdateWalletRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateWalletRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateWalletResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}

func (o *UpdateWalletResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWalletResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWalletResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWalletResponse) GetWalletsErrorResponse() *shared.WalletsErrorResponse {
	if o == nil {
		return nil
	}
	return o.WalletsErrorResponse
}
