package identity

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//LaunchAppResponse ...
type LaunchAppResponse struct {
	Result struct {
		LaunchCode string `json:"launchCode"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

//LaunchApp uses JWT and AccountID to launch app and get launch code.
func (a *API) LaunchApp(jwt string, accountID int) (*LaunchAppResponse, error) {

	var apiResponse = &LaunchAppResponse{}

	params := &url.Values{}
	params.Set("api[jwt]", jwt)
	params.Set("parameters[id]", strconv.Itoa(accountID))

	resp, err := a.postRequest(params, "V1/Launchpad/launch")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return nil, err
	}

	return apiResponse, nil
}
