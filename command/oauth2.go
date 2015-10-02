package command

import (
	"net/http"

	"bitbucket.org/pqstudio/go-oauth2/server"
)

func IsTokenAuthorized(r *http.Request) bool {
	token := server.AccessToken(r)

	_, err := server.Server.Storage.LoadAccess(token)

	if err != nil {
		return false
	} else {
		return true
	}
}
