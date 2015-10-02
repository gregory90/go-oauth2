package oauth2

import (
	"database/sql"

	"github.com/RangelReale/osin"

	"bitbucket.org/pqstudio/go-oauth2/db"
	"bitbucket.org/pqstudio/go-oauth2/server"
	"bitbucket.org/pqstudio/go-oauth2/storage"
)

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

func Init(DB *sql.DB) {
	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.TOKEN}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.REFRESH_TOKEN, osin.PASSWORD, osin.ASSERTION}
	sconfig.AllowGetAccessRequest = false
	server.Init(osin.NewServer(sconfig, storage.NewMySQLStorage()))
	db.Init(DB)
}
