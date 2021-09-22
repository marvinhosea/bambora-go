package config

import (
	"github.com/marvinhosea/bambora-go/util"
	"log"
	"net/http"
)

const (
	APIUrl string = "https://api.na.bambora.com/v1"
	OauthUrl string = "https://api.na.bambora.com/scripts"
	APIEndpoint string = "api"
	// Connect For Oauth backend
	Connect string = "connect"
	ApiVersion = "v1"
	ProfilePaymentMethod = "payment_profile"
)

type Config struct {
	HttpClient *http.Client
	Url *string
	MerchantId       string
	EProfilePasscode string
	EPaymentPasscode string
	ApiVersion       string
}

var config *Config

func New(merchantId, profilePasscode string, paymentPasscode string) *Config {
	if config == nil {
		config = &Config{
			MerchantId:       merchantId,
			EPaymentPasscode: paymentPasscode,
			EProfilePasscode: profilePasscode,
			Url:              util.String(OauthUrl),
			ApiVersion:       ApiVersion,
		}
	}

	return config
}

var Con = func() *Config {
	if config == nil {
		log.Fatalln("set up Bambora config first")
	}

	return config
}