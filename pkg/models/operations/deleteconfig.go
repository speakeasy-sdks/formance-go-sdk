package operations

import (
	"net/http"
)

type DeleteConfigRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteConfigResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
