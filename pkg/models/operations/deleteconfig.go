package operations

import (
	"net/http"
)

type DeleteConfigPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteConfigRequest struct {
	PathParams DeleteConfigPathParams
}

type DeleteConfigResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
