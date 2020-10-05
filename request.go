package disgo

import (
	"bytes"
	"net/http"
)

func newRequestWithBody(client *Client, method string, path string, body []byte) (*http.Request, error) {
	s := client.Endpoint + "/" + path

	req, err := http.NewRequest(method, s, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Close = true

	req.Header.Add("Api-Key", client.APIKey)
	req.Header.Add("Api-Username", client.Username)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, err
}

func executeRequest(client *Client, req *http.Request) (resp *http.Response, err error) {
	httpc := &http.Client{
		Timeout: client.timeout,
	}
	if client.transport != nil {
		httpc.Transport = client.transport
	}
	resp, err = httpc.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}