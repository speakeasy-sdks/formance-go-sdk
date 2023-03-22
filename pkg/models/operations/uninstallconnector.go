// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type UninstallConnectorRequest struct {
	// The name of the connector.
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type UninstallConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
