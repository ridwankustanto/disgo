package disgo

import (
	"encoding/json"
	"net/http"
)

func parseBodyResponse(resp *http.Response) (interface{}, error) {
	var parsedResp interface{}

	//bodyResp, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}

	if err := json.NewDecoder(resp.Body).Decode(&parsedResp); err != nil {
		return nil, err
	}

	return parsedResp, nil
}
