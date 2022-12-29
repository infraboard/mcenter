package github

type Config struct {
	AppId        string `json:"app_id" env:"GITHUB_APP_ID"`
	ClientId     string `json:"client_id" env:"GITHUB_CLIENT_ID"`
	ClientSecret string `json:"client_secret" env:"GITHUB_CLIENT_SECRET"`
}
