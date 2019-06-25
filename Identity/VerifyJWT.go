package identity

import (
	"encoding/json"
	"net/url"
)

//VerifyJWTResponse ...
type VerifyJWTResponse struct {
	Result bool `json:"result"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// VerifyJWT persistence token by sending JWT.
// If it's valid then returns (boolean) TRUE and if FALSE then token is expired or not exist.
// Also, if it return FALSE then in error section is error code 1087 with message "Persistence token is not valid
// or not exist".
func (a *API) VerifyJWT(jwt string) (*VerifyJWTResponse, error) {

	apiResponse := &VerifyJWTResponse{}

	params := &url.Values{}
	params.Set("api[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Launchpad/verifyJWT")

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(resp), &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}
