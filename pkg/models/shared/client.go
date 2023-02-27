package shared

type Client struct {
	Description            *string                `json:"description,omitempty"`
	ID                     string                 `json:"id"`
	Metadata               map[string]interface{} `json:"metadata,omitempty"`
	Name                   string                 `json:"name"`
	PostLogoutRedirectUris []string               `json:"postLogoutRedirectUris,omitempty"`
	Public                 *bool                  `json:"public,omitempty"`
	RedirectUris           []string               `json:"redirectUris,omitempty"`
	Scopes                 []string               `json:"scopes,omitempty"`
	Secrets                []ClientSecret         `json:"secrets,omitempty"`
	Trusted                *bool                  `json:"trusted,omitempty"`
}
