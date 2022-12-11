package dingding

type GetUserTokenRequest struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Code         string `json:"code"`
	RefreshToken string `json:"refreshToken"`
	GrantType    string `json:"grantType"`
}
