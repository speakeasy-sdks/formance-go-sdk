package operations

import (
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
)

type InsertConfigRequest struct {
	Request shared.ConfigUser `request:"mediaType=application/json"`
}

type InsertConfigResponse struct {
	ConfigResponse                 *shared.ConfigResponse
	ContentType                    string
	StatusCode                     int
	InsertConfig400TextPlainString *string
}
