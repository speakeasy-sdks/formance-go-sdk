package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchRequest struct {
	Request shared.Query `request:"mediaType=application/json"`
}

type SearchResponse struct {
	ContentType string
	Response    *shared.Response
	StatusCode  int
	RawResponse *http.Response
}
