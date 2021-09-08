package config

import "net/http"

const (
	APIUrl string = "https://"
	OauthUrl string = "https://"
	APIEndpoint string = "api"
	// Connect For Oauth backend
	Connect string = "connect"
)

type Config struct {
	HttpClient *http.Client
	Url *string
	MerchantId string
	Passcode string
	ApiVersion string
}