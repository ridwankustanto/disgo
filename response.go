package disgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseBodyResponse(resp *http.Response) (interface{}, error) {
	var parsedResp interface{}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bodyResp, &parsedResp); err != nil {
		return nil, err
	}

	return parsedResp, nil
}