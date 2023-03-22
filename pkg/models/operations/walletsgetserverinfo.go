// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type WalletsgetServerInfoResponse struct {
	ContentType string
	// Server information
	ServerInfo  *shared.ServerInfo
	StatusCode  int
	RawResponse *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}
