package operations

import (
	"net/http"
)

type DeleteClientPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type DeleteClientRequest struct {
	PathParams DeleteClientPathParams
}

type DeleteClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
