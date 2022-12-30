package github

import "fmt"

func NewTokenReponse() *TokenReponse {
	return &TokenReponse{}
}

type TokenReponse struct {
	Token
	Response
}

type Token struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int64  `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	TokenType             string `json:"token_type"`
}

type Response struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
}

func (r *Response) Error() error {
	if r.ErrorCode != "" {
		return fmt.Errorf("code: %s, description: %s, docs: %s",
			r.ErrorCode,
			r.ErrorDescription,
			r.ErrorUri,
		)
	}

	return nil
}
