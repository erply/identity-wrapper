package Identity

import (
	"encoding/json"
	"net/url"
)

type GetUserConfirmedCompaniesResponse struct {
	Result struct {
		Companies []Company `json:"companies"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type Company struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Status    bool   `json:"status"`
	IsDefault bool   `json:"isDefault"`
}

// Get list of companies where user has access.
func (a *API) GetUserConfirmedCompanies(jwt string) (GetUserConfirmedCompaniesResponse, error) {

	params := &url.Values{}
	params.Set("api[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Launchpad/getUserConfirmedCompanies")

	if err != nil {
		return GetUserConfirmedCompaniesResponse{}, err
	}

	var apiResponse = GetUserConfirmedCompaniesResponse{}
	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return GetUserConfirmedCompaniesResponse{}, err
	}

	return apiResponse, nil
}
