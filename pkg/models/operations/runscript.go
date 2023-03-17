package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"net/http"
)

type RunScriptRequest struct {
	Script  shared.Script `request:"mediaType=application/json"`
	Ledger  string        `pathParam:"style=simple,explode=false,name=ledger"`
	Preview *bool         `queryParam:"style=form,explode=true,name=preview"`
}

type RunScriptResponse struct {
	ContentType    string
	ScriptResponse *shared.ScriptResponse
	StatusCode     int
	RawResponse    *http.Response
}
