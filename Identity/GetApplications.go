package Identity

import (
	"encoding/json"
	"net/url"
)

type GetApplicationsResponse struct {
	Result struct {
		Applications []Application `json:"applications"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Application struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	URL              string `json:"url"`
	Title            string `json:"title"`
	Icon             string `json:"icon"`
	Active           bool   `json:"active"`
	AppStoreDisabled int    `json:"app_store_disabled"`
}

// Get list of all applications.
func (a *API) GetApplications(jwt string) (GetApplicationsResponse, error) {

	params := &url.Values{}
	params.Set("api[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Launchpad/applications")

	if err != nil {
		return GetApplicationsResponse{}, err
	}

	var apiResponse = GetApplicationsResponse{}
	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return GetApplicationsResponse{}, err
	}

	return apiResponse, nil
}
