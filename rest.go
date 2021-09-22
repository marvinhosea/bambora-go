package bambora

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type RestClient struct {

}

func (r *RestClient) Get(path, passcode string, params map[string]interface{}, headers http.Header) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	req.URL.RawQuery = q.Encode()
	req.Header = headers
	req.Header.Add("Content-Type", `application/json;charset=utf-8`)
	if len(passcode) != 0 {
		req.Header.Add("Authorization", "Passcode " + passcode)
	}

	return req, nil
}

func (r *RestClient) Post(path, passcode string, body interface{}, headers http.Header) (*http.Request, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	log.Println("bddfy", body)

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		req.Header = headers
	}
	req.Header.Add("Content-Type", `application/json;charset=utf-8`)
	req.Header.Add("Authorization", "Passcode " + passcode)
	return req, nil
}