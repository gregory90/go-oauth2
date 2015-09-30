package storage

import (
	"database/sql"
	"errors"

	"github.com/RangelReale/osin"

	"bitbucket.org/pqstudio/go-oauth2/service"

	. "bitbucket.org/pqstudio/go-webutils/logger"
)

type MySQLStorage struct {
}

func NewMySQLStorage() *MySQLStorage {
	r := &MySQLStorage{}

	return r
}

func (s *MySQLStorage) Clone() osin.Storage {
	return s
}

func (s *MySQLStorage) Close() {
}

func (s *MySQLStorage) GetClient(tx *sql.Tx, id string) (osin.Client, error) {
	Log.Notice("OAuth2, get client: %s\n", id)
	c, err := service.GetClientByID(tx, id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *MySQLStorage) SetClient(tx *sql.Tx, id string, client osin.Client) error {
	Log.Notice("OAuth2, set client: %s\n", id)

	c := &osin.DefaultClient{
		Id:          client.GetId(),
		Secret:      client.GetSecret(),
		RedirectUri: client.GetRedirectUri(),
	}

	err := service.CreateClient(tx, c)
	return err
}

func (s *MySQLStorage) SaveAuthorize(tx *sql.Tx, data *osin.AuthorizeData) error {
	return errors.New("Not implemented")
}

func (s *MySQLStorage) LoadAuthorize(tx *sql.Tx, code string) (*osin.AuthorizeData, error) {
	return nil, errors.New("Not implemented")
}

func (s *MySQLStorage) RemoveAuthorize(tx *sql.Tx, code string) error {
	return errors.New("Not implemented")
}

func (s *MySQLStorage) SaveAccess(tx *sql.Tx, data *osin.AccessData) error {
	Log.Notice("OAuth2, save access: %s\n", data.AccessToken)

	err := service.CreateAccess(tx, data)
	return err
}

func (s *MySQLStorage) LoadAccess(tx *sql.Tx, code string) (*osin.AccessData, error) {
	Log.Notice("OAuth2, load access: %s\n", code)

	a, err := service.GetAccessByToken(tx, code)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *MySQLStorage) RemoveAccess(tx *sql.Tx, code string) error {
	Log.Notice("OAuth2, remove access: %s\n", code)

	err := service.DeleteAccessByToken(tx, code)
	return err
}

func (s *MySQLStorage) LoadRefresh(tx *sql.Tx, code string) (*osin.AccessData, error) {
	Log.Notice("OAuth2, load refresh: %s\n", code)
	a, err := service.GetAccessByRefresh(tx, code)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *MySQLStorage) RemoveRefresh(tx *sql.Tx, code string) error {
	Log.Notice("OAuth2, remove refresh: %s\n", code)
	return nil
}
