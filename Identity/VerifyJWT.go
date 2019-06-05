package Identity

import (
	"encoding/json"
	"net/url"
)

type VerifyJWTResponse struct {
	Result bool `json:"result"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Verify persistence token by sending JWT.
// If it's valid then returns (boolean) TRUE and if FALSE then token is expired or not exist.
// Also, if it return FALSE then in error section is error code 1087 with message "Persistence token is not valid
// or not exist".
func (a *API) VerifyJWT(jwt string) (VerifyJWTResponse, error) {

	params := &url.Values{}
	params.Set("api[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Launchpad/verifyJWT")

	if err != nil {
		return VerifyJWTResponse{}, err
	}

	var apiResponse = VerifyJWTResponse{}
	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		return VerifyJWTResponse{}, err
	}

	return apiResponse, nil
}
