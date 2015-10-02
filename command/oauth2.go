package command

import (
	"net/http"
)

func IsTokenAuthorized(r *http.Request) bool {
	token := AccessToken(r)

	_, err := Server.Storage.LoadAccess(token)

	if err != nil {
		return false
	} else {
		return true
	}
}
