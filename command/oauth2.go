package command

import (
	"database/sql"
	"net/http"

	"github.com/RangelReale/osin"

	"github.com/gregory90/go-oauth2/datastore"
	"github.com/gregory90/go-oauth2/server"
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
	_, err := datastore.GetAccessByToken(tx, token)
	if err != nil {
		return err
	}
	err = server.Server.Storage.RemoveAccess(token)
	if err != nil {
		return err
	}
	return nil
}

func HandleAccessRequest(r *http.Request, data server.Data) (*osin.AccessRequest, *osin.Response) {
	r.ParseForm()
	r.Form.Add("grant_type", data.GrantType)
	r.Form.Add("username", data.Username)
	r.Form.Add("password", data.Password)
	r.Form.Add("assertion", data.Assertion)
	r.Form.Add("assertion_type", data.AssertionType)
	r.Form.Add("scope", data.Scope)
	r.Form.Add("token", data.Token)
	r.Form.Add("refresh_token", data.RefreshToken)

	r.SetBasicAuth(data.ClientID, data.ClientSecret)

	resp := server.Server.NewResponse()
	defer resp.Close()
	ar := server.Server.HandleAccessRequest(resp, r)
	return ar, resp
}
