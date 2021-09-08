package go_bambora

import (
	"crypto/tls"
	"github.com/marvinhosea/bambora-go/config"
	"github.com/marvinhosea/bambora-go/util"
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

var httpClient = &http.Client{
	Timeout: 80*time.Second,
	Transport: &http.Transport{TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper)}}

type ApiResponse struct {
	Headers http.Header
	RawJson []byte
	Status string
	StatusCode string
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
}

func newApiResponse(res *http.Response, body []byte) *ApiResponse {
	return &ApiResponse{
		Headers: res.Header,
		RawJson: body,
	}
}

func (a *APIResource) SetLastResponse(response *ApiResponse)  {

}

func (i *EndpointImplementation) Call(method, path, passcode string, params map[string]string, v LastResponseSetter) error {
	log.Printf("working")
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
	log.Println("end", ep)
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
	//generate passcode from config
	return "djkdajfnkjndsf"
}

func newEndpointImplementation(endpointType string, config *config.Config) Endpoint {
	return &EndpointImplementation{
		HTTPClient: config.HttpClient,
		Url: *config.Url,
	}
}