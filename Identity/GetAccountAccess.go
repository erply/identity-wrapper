package identity

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//GetAccountAccessResponse ...
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

// GetAccountAccess gets application IDs to see where user has access.
// Get AccountID to launch app.
func (a *API) GetAccountAccess(jwt string, companyID int) (*GetAccountAccessResponse, error) {

	apiResponse := &GetAccountAccessResponse{}

	params := &url.Values{}
	params.Set("api[jwt]", jwt)
	params.Set("parameters[companyId]", strconv.Itoa(companyID))

	resp, err := a.postRequest(params, "V1/Launchpad/accountAccess")

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(resp), &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}
