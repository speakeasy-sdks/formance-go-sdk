// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type DeactivateConfigRequest struct {
	// Config ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeactivateConfigResponse struct {
	// Config successfully deactivated.
	ConfigResponse *shared.ConfigResponse
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}
