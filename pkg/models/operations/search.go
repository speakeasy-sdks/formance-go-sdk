package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchResponse struct {
	ContentType string
	Response    *shared.Response
	StatusCode  int
	RawResponse *http.Response
}
