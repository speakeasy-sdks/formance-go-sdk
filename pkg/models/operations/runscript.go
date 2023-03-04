package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type RunScriptPathParams struct {
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type RunScriptQueryParams struct {
	Preview *bool `queryParam:"style=form,explode=true,name=preview"`
}

type RunScriptRequest struct {
	PathParams  RunScriptPathParams
	QueryParams RunScriptQueryParams
	Request     shared.Script `request:"mediaType=application/json"`
}

type RunScriptResponse struct {
	ContentType    string
	ScriptResponse *shared.ScriptResponse
	StatusCode     int
	RawResponse    *http.Response
}
