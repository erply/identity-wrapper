package identity

import (
	"encoding/json"
	"net/url"
)

//GetApplicationsResponse ...
type GetApplicationsResponse struct {
	Result struct {
		Applications []Application `json:"applications"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

//Application ...
type Application struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	URL              string `json:"url"`
	Title            string `json:"title"`
	Icon             string `json:"icon"`
	Active           bool   `json:"active"`
	AppStoreDisabled int    `json:"app_store_disabled"`
}

//GetApplications Gets list of all applications.
func (a *API) GetApplications(jwt string) (*GetApplicationsResponse, error) {

	apiResponse := &GetApplicationsResponse{}

	params := &url.Values{}
	params.Set("api[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Launchpad/applications")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return nil, err
	}

	return apiResponse, nil
}
