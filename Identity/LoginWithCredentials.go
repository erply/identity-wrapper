package Identity

import (
	"encoding/json"
	"log"
	"net/url"
)

type LoginWithCredentialsResponse struct {
	Result struct {
		DefaultCompanyId int    `json:"defaultCompanyId"`
		IsNewUser        bool   `json:"isNewUser"`
		JWT              string `json:"jwt"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Login to Identity Launchpad with email and password.
// Successfully logged in user will get JWT and default company ID to use it in further requests.
// With this JWT is possible to do requests in limited permissions.
// Use https://jwt.io to see inside your JWT. Use algorithm RS256.
func (a *API) LoginWithCredentials(email string, password string) (LoginWithCredentialsResponse, error) {

	params := &url.Values{}
	params.Set("parameters[email]", email)
	params.Set("parameters[password]", password)

	resp, err := a.postRequest(params, "V1/Launchpad/login")

	if err != nil {
		return LoginWithCredentialsResponse{}, err
	}

	var apiResponse = LoginWithCredentialsResponse{}
	err = json.Unmarshal([]byte(resp), &apiResponse)

	if err != nil {
		log.Println("Login with credentials unmarshal failed", err)
		return LoginWithCredentialsResponse{}, err
	}

	return apiResponse, nil
}
