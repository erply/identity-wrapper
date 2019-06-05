package Identity

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type GetAccountAccessResponse struct {
	Result struct {
		AccountID    int   `json:"id"`
		Applications []int `json:"applications"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Get application IDs to see where user has access.
// Get AccountID to launch app.
func (a *API) GetAccountAccess(jwt string, companyID int) (GetAccountAccessResponse, error) {

	params := &url.Values{}
	params.Set("api[jwt]", jwt)
	params.Set("parameters[companyId]", strconv.Itoa(companyID))

	resp, err := a.postRequest(params, "V1/Launchpad/accountAccess")

	if err != nil {
		return GetAccountAccessResponse{}, err
	}

	var apiResponse = GetAccountAccessResponse{}
	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return GetAccountAccessResponse{}, err
	}

	return apiResponse, nil
}
