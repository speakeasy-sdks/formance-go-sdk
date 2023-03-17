package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type InsertConfigResponse struct {
	ConfigResponse                 *shared.ConfigResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
	InsertConfig400TextPlainString *string
}
