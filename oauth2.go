package oauth2

import (
	"database/sql"

	"github.com/RangelReale/osin"

	"github.com/gregory90/go-oauth2/db"
	"github.com/gregory90/go-oauth2/server"
	"github.com/gregory90/go-oauth2/storage"
)

func Init(DB *sql.DB) {
	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.TOKEN}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.REFRESH_TOKEN, osin.PASSWORD, osin.ASSERTION}
	sconfig.AllowGetAccessRequest = false
	server.Init(osin.NewServer(sconfig, storage.NewMySQLStorage()))
	db.Init(DB)
}
