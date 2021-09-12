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
	Call(method, path, passcode string, params map[string]string, v LastResponseSetter) error
}

type Endpoints struct {
	API, Connect Endpoint
	mu sync.RWMutex
}

var endpoints Endpoints

// MerchantId Merchant Id which is used globally
var MerchantId string

//AccountPasscode which is used Globally
var AccountPasscode string

// Passcode is Bambora global request passcode used to query
var Passcode string

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

func (i *EndpointImplementation) Call(method, path, passcode string, params map[string]string, v LastResponseSetter) error {
	return i.NewRequest(method, path, passcode, params, v)
}

func (i *EndpointImplementation) NewRequest(method, path, passcode string, params map[string]string, v LastResponseSetter) error {
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

func GeneratePasscode() string {
	config := config.New(MerchantId, AccountPasscode)
	pc, err := generatePasscode(config)
	if err != nil {
		log.Fatalln(err)
	}
	return pc
}

func generatePasscode(config *config.Config) (string, error) {
	if len(config.MerchantId) == 0 {
		return "", errors.New("error: merchant id is empty")
	}

	if len(config.Passcode) == 0 {
		return "", errors.New("error: passcode is empty")
	}

	passcode := base64.StdEncoding.EncodeToString([]byte(config.MerchantId + ":" + config.Passcode))
	if len(passcode) == 0 {
		return "", errors.New("error: generated passcode is empty")
	}
	return passcode, nil
}

func newEndpointImplementation(endpointType string, config *config.Config) Endpoint {
	return &EndpointImplementation{
		HTTPClient: config.HttpClient,
		Url: *config.Url,
		Type: endpointType,
		restClient: RestClient{},
	}
}