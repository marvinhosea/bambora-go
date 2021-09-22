package bambora

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	//"github.com/marvinhosea/bambora-go/client"
	"github.com/marvinhosea/bambora-go/config"
	"github.com/marvinhosea/bambora-go/util"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Endpoint interface {
	Call(method, path, passcode string, params map[string]interface{}, v LastResponseSetter) error
}

type Endpoints struct {
	API, Connect Endpoint
	mu sync.RWMutex
}

var endpoints Endpoints

// MerchantId Merchant Id which is used globally
var MerchantId string

// EncodedProfilePasscode AccountPasscode which is used Globally
var EncodedProfilePasscode string

// EncodedPaymentPasscode AccountPasscode which is used Globally
var EncodedPaymentPasscode string

// EncodedPaymentPasscode AccountPasscode which is used Globally
var PaymentMethod string

// ProfilePasscode Passcode is Bambora global request passcode used to query
var ProfilePasscode string

// PaymentPasscode Passcode is Bambora global request passcode used to query
var PaymentPasscode string

var httpClient = &http.Client{
	Timeout: 80*time.Second,
	Transport: &http.Transport{TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper)}}

type ApiResponse struct {
	Headers http.Header
	RawJson []byte
	Status string
	StatusCode int
}

type LastResponseSetter interface {
	SetLastResponse(response *ApiResponse)
}

type APIResource struct {
	LastResponse *ApiResponse `json:"-"`
}

type EndpointImplementation struct {
	HTTPClient *http.Client
	Type string
	Url string
	restClient RestClient
}

func newApiResponse(res *http.Response, body []byte) *ApiResponse {
	return &ApiResponse{
		Headers: res.Header,
		RawJson: body,
		Status: res.Status,
		StatusCode: res.StatusCode,
	}
}

func (a *APIResource) SetLastResponse(response *ApiResponse)  {
	a.LastResponse = response
}

func (i *EndpointImplementation) Call(method, path, passcode string, params map[string]interface{}, v LastResponseSetter) error {
	return i.NewRequest(method, path, passcode, params, v)
}

func (i *EndpointImplementation) NewRequest(method, path, passcode string, params map[string]interface{}, v LastResponseSetter) error {
	url := i.Url + path
	var req *http.Request
	var err error
	if method == http.MethodPost {
		req, err = i.restClient.Post(url, passcode, params, nil)
		if err != nil {
			return err
		}
	}

	if method == http.MethodGet {
		req, err = i.restClient.Post(url, passcode, params, nil)
		if err != nil {
			return err
		}
	}
	var result []byte
	res, err := i.HTTPClient.Do(req)
	if err == nil {
		result, err = ioutil.ReadAll(res.Body)
		log.Println("result", string(result))
		res.Body.Close()
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(result, v)
	if err != nil {
		log.Println("Could not unmarshal response and something happened", err)
		return err
	}
	v.SetLastResponse(newApiResponse(res, result))

	return nil
}

func GetEndpoint(endpointType string) Endpoint {
	var ep Endpoint
	endpoints.mu.RLock()
	switch endpointType {
	case config.APIEndpoint:
		ep = endpoints.API
	case config.Connect:
		ep = endpoints.Connect
	}
	endpoints.mu.RUnlock()
	ep = GetEndpointWithConfig(endpointType, &config.Config{
		HttpClient: httpClient,
		Url: nil,
	})
	log.Println("endpoint", ep)
	if ep != nil {
		return ep
	}
	return ep
}

func GetEndpointWithConfig(ep string, cnf *config.Config) Endpoint {
	if cnf.HttpClient == nil {
		cnf.HttpClient = httpClient
	}
	switch ep {
	case config.APIEndpoint:
		if cnf.Url == nil {
			cnf.Url = util.String(config.APIUrl)
		}
		return newEndpointImplementation(ep, cnf)
	case config.Connect:
		if cnf.Url == nil {
			cnf.Url = util.String(config.OauthUrl)
		}
		return newEndpointImplementation(ep, cnf)
	}

	return nil
}

func GeneratePasscodes(merchantId, profilePasscode string, paymentPasscode string) *config.Config {
	MerchantId = merchantId
	PaymentPasscode = paymentPasscode
	ProfilePasscode = profilePasscode

	paymentP, err := generatePasscode(MerchantId, PaymentPasscode)
	if err != nil {
		return nil
	}
	EncodedPaymentPasscode = paymentP

	profileP, err := generatePasscode(MerchantId, ProfilePasscode)
	if err != nil {
		return nil
	}

	return config.New(MerchantId, profileP, paymentP)
}

func generatePasscode(merchantId, passcode string) (string, error) {
	if len(merchantId) == 0 {
		return "", errors.New("error: merchant id is empty")
	}

	if len(passcode) == 0 {
		return "", errors.New("error: passcode is empty")
	}

	base64pc := base64.StdEncoding.EncodeToString([]byte(merchantId + ":" + passcode))
	if len(base64pc) == 0 {
		return "", errors.New("error: generated passcode is empty")
	}
	return base64pc, nil
}

func newEndpointImplementation(endpointType string, config *config.Config) Endpoint {
	return &EndpointImplementation{
		HTTPClient: config.HttpClient,
		Url: *config.Url,
		Type: endpointType,
		restClient: RestClient{},
	}
}