package server

import (
	"net/http"

	"github.com/RangelReale/osin"
)

var Server *osin.Server

func Init(server *osin.Server) {
	Server = server
}

type Data struct {
	GrantType     string `json:"grant_type,omitempty"`
	ClientID      string `json:"client_id,omitempty"`
	ClientSecret  string `json:"client_secret,omitempty"`
	Username      string `json:"username,omitempty"`
	AssertionType string `json:"assertion_type,omitempty"`
	Assertion     string `json:"assertion,omitempty"`
	Token         string `json:"token,omitempty"`
	Password      string `json:"password,omitempty"`
	Scope         string `json:"scope,omitempty"`
	RefreshToken  string `json:"refresh_token,omitempty"`
}

func AccessToken(r *http.Request) string {
	auth := r.Header["Authorization"]

	var token string
	if len(auth) > 0 && len(auth[0]) > 7 {
		token = auth[0][7:]
	}

	return token
}
