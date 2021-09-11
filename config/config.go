package config

import "net/http"

const (
	APIUrl string = "https://"
	OauthUrl string = "https://"
	APIEndpoint string = "api"
	// Connect For Oauth backend
	Connect string = "connect"
	ApiVersion = "v1"
)

type Config struct {
	HttpClient *http.Client
	Url *string
	MerchantId string
	Passcode string
	ApiVersion string
}

var config *Config

func New(merchantId, passcode string) *Config {
	if config == nil {
		config = &Config{
			MerchantId: merchantId,
			Passcode: passcode,
			ApiVersion: ApiVersion,
		}
	}

	return config
}