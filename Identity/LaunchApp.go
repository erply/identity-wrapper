package Identity

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type LaunchAppResponse struct {
	Result struct {
		LaunchCode string `json:"launchCode"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Use JWT and AccountID to launch app and get launch code.
func (a *API) LaunchApp(jwt string, accountID int) (LaunchAppResponse, error) {

	params := &url.Values{}
	params.Set("api[jwt]", jwt)
	params.Set("parameters[id]", strconv.Itoa(accountID))

	resp, err := a.postRequest(params, "V1/Launchpad/launch")

	if err != nil {
		return LaunchAppResponse{}, err
	}

	var apiResponse = LaunchAppResponse{}
	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return LaunchAppResponse{}, err
	}

	return apiResponse, nil
}
