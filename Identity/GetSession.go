package identity

import (
	"encoding/json"
	"net/url"
)

//GetSessionResponse ...
type GetSessionResponse struct {
	Result struct {
		Session string `json:"session"`
	}
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}


// GetSessionID by JWT.
// It returns short session token for Builder applications and Service Engine endpoints.
// Use this token in headers.
// `JSESSIONID` - if using Builder apps.
// `API_KEY` - if using Service Engine endpoints.
func (a *API) GetSessionID(jwt string) (*GetSessionResponse, error) {

	apiResponse := &GetSessionResponse{}

	params := &url.Values{}
	params.Set("parameters[jwt]", jwt)

	resp, err := a.postRequest(params, "V1/Authentication/getSession")

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(resp), &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}
