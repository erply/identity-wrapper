package identity

import (
	"encoding/json"
	"log"
	"net/url"
)

//LoginWithCredentialsResponse ...
type LoginWithCredentialsResponse struct {
	Result struct {
		DefaultCompanyID int    `json:"defaultCompanyId"`
		IsNewUser        bool   `json:"isNewUser"`
		JWT              string `json:"jwt"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// LoginWithCredentials Logins to Identity Launchpad with email and password.
// Successfully logged in user will get JWT and default company ID to use it in further requests.
// With this JWT is possible to do requests in limited permissions.
// Use https://jwt.io to see inside your JWT. Use algorithm RS256.
func (a *API) LoginWithCredentials(email string, password string) (*LoginWithCredentialsResponse, error) {

	apiResponse := &LoginWithCredentialsResponse{}

	params := &url.Values{}
	params.Set("parameters[email]", email)
	params.Set("parameters[password]", password)

	resp, err := a.postRequest(params, "V1/Launchpad/login")

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(resp), &apiResponse); err != nil {
		log.Println("Login with credentials unmarshal failed", err)
		return nil, err
	}

	return apiResponse, nil
}
