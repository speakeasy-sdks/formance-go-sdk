package shared

type SchemeAuthorization struct {
	Authorization string `security:"name=Authorization"`
}

type Security struct {
	Authorization SchemeAuthorization `security:"scheme,type=oauth2"`
}
