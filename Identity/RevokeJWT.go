package Identity

import (
	"encoding/json"
	"net/url"
)

type RevokeJWTResponse struct {
	Result bool `json:"result"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Revoke persistence token by sending JWT.
func (a *API) RevokeJWT(jwt string) (RevokeJWTResponse, error) {

	params := &url.Values{}
	params.Set("api[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Launchpad/revoke")

	if err != nil {
		return RevokeJWTResponse{}, err
	}

	var apiResponse = RevokeJWTResponse{}
	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return RevokeJWTResponse{}, err
	}

	return apiResponse, nil
}
