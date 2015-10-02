package command

import (
	"database/sql"
	"net/http"

	"bitbucket.org/pqstudio/go-oauth2/datastore"
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

func InvalidateToken(tx *sql.Tx, r *http.Request) error {
	token := server.AccessToken(r)
	access, err := datastore.GetAccessByToken(tx, token)
	if err != nil {
		return err
	}
	err = server.Server.Storage.RemoveAccess(token)
	if err != nil {
		return err
	}
	return nil
}
