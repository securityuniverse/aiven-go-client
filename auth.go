package aiven

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	// Token represents a user token.
	Token struct {
		Token string `json:"token"`
		State string `json:"state"`
	}

	authRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	authResponse struct {
		Errors  []Error `json:"errors"`
		Message string  `json:"message"`
		State   string  `json:"state"`
		Token   string  `json:"token"`
	}
)

// UserToken retrieves a User Auth Token for a given user/password pair.
func UserToken(email, password string, client *http.Client) (*Token, error) {
	if client == nil {
		client = &http.Client{}
	}

	bts, err := json.Marshal(authRequest{email, password})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint("userauth"), bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	bts, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var response *authResponse
	if err := json.Unmarshal(bts, &response); err != nil {
		return nil, err
	}

	return &Token{response.Token, response.State}, nil
}
