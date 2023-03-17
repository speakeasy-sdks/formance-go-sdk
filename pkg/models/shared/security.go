package shared

type Security struct {
	Authorization string `security:"scheme,type=oauth2,name=Authorization"`
}
