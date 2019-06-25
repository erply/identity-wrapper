package identity

import (
	"encoding/json"
	"net/url"
)

//RevokeJWTResponse ...
type RevokeJWTResponse struct {
	Result bool `json:"result"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// RevokeJWT revokes persistence token by sending JWT.
func (a *API) RevokeJWT(jwt string) (*RevokeJWTResponse, error) {

	apiResponse := &RevokeJWTResponse{}

	params := &url.Values{}
	params.Set("api[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Launchpad/revoke")

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(resp), &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}
