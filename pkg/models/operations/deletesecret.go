package operations

type DeleteSecretPathParams struct {
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
	SecretID string `pathParam:"style=simple,explode=false,name=secretId"`
}

type DeleteSecretRequest struct {
	PathParams DeleteSecretPathParams
}

type DeleteSecretResponse struct {
	ContentType string
	StatusCode  int
}
