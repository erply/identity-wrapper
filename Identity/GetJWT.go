package identity

import (
	"encoding/json"
	"net/url"
)

//GetJWTResponse ...
type GetJWTResponse struct {
	Result struct {
		JWT string `json:"JWT"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// GetJWT Gets JWT by launchCode.
// LunchCode is a hash which expires after 30 sec.
// Returns JWT with all permissions you have.
func (a *API) GetJWT(launchCode string) (*GetJWTResponse, error) {

	apiResponse := GetJWTResponse{}

	params := &url.Values{}
	params.Set("parameters[launchCode]", launchCode)

	resp, err := a.postRequest(params, "V1/Launchpad/getJWT")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
