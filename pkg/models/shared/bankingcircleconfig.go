package shared

type BankingCircleConfig struct {
	AuthorizationEndpoint string `json:"authorizationEndpoint"`
	Endpoint              string `json:"endpoint"`
	Password              string `json:"password"`
	Username              string `json:"username"`
}
