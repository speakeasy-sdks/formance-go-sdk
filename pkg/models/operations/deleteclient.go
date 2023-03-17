package operations

import (
	"net/http"
)

type DeleteClientRequest struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

type DeleteClientResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
