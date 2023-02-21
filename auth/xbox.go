package auth

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sandertv/gophertunnel/minecraft/auth"
	"golang.org/x/oauth2"
)

var tokenSource oauth2.TokenSource

func TokenSource() (tokenSource oauth2.TokenSource) {
	if tokenSource != nil {
		return
	}

	token := &oauth2.Token{}

	data, err := ioutil.ReadFile("token")

	if err == nil {
		_ = json.Unmarshal(data, token)
	} else {
		token, err = auth.RequestLiveToken()
		if err != nil {
			panic(err)
		}
	}

	src := auth.RefreshTokenSource(token)
	_, err = src.Token()
	if err != nil {
		// The cached refresh token expired and can no longer be used to obtain a new token. We require the
		// user to log in again and use that token instead.
		token, err = auth.RequestLiveToken()
		if err != nil {
			panic(err)
		}

		src = auth.RefreshTokenSource(token)
	}

	tokenSource = src

	tok, _ := src.Token()
	b, _ := json.Marshal(tok)
	_ = ioutil.WriteFile("token", b, 0644)

	return
}
