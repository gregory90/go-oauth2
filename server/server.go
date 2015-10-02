package server

import (
	"net/http"

	"github.com/RangelReale/osin"
)

var Server *osin.Server

func Init(server *osin.Server) {
	Server = server
}

func AccessToken(r *http.Request) string {
	auth := r.Header["Authorization"]

	var token string
	if len(auth) > 0 && len(auth[0]) > 7 {
		token = auth[0][7:]
	}

	return token
}
