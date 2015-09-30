package storage

import (
	"database/sql"
	"errors"

	"github.com/RangelReale/osin"

	. "bitbucket.org/pqstudio/go-oauth2/db"
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

func (s *MySQLStorage) GetClient(id string) (osin.Client, error) {
	Log.Notice("OAuth2, get client: %s\n", id)
	var c *osin.DefaultClient
	err := Transact(DB, func(tx *sql.Tx) error {
		var err error
		c, err = service.GetClientByID(tx, id)
		return err
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *MySQLStorage) SetClient(id string, client osin.Client) error {
	Log.Notice("OAuth2, set client: %s\n", id)

	c := &osin.DefaultClient{
		Id:          client.GetId(),
		Secret:      client.GetSecret(),
		RedirectUri: client.GetRedirectUri(),
	}

	err := Transact(DB, func(tx *sql.Tx) error {
		var err error
		err = service.CreateClient(tx, c)
		return err
	})
	return err
}

func (s *MySQLStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	return errors.New("Not implemented")
}

func (s *MySQLStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	return nil, errors.New("Not implemented")
}

func (s *MySQLStorage) RemoveAuthorize(code string) error {
	return errors.New("Not implemented")
}

func (s *MySQLStorage) SaveAccess(data *osin.AccessData) error {
	Log.Notice("OAuth2, save access: %s\n", data.AccessToken)

	err := Transact(DB, func(tx *sql.Tx) error {
		var err error
		err = service.CreateAccess(tx, data)
		return err
	})
	return err
}

func (s *MySQLStorage) LoadAccess(code string) (*osin.AccessData, error) {
	Log.Notice("OAuth2, load access: %s\n", code)

	var a *osin.AccessData
	err := Transact(DB, func(tx *sql.Tx) error {
		var err error
		a, err = service.GetAccessByToken(tx, code)
		return err
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *MySQLStorage) RemoveAccess(code string) error {
	Log.Notice("OAuth2, remove access: %s\n", code)

	err := Transact(DB, func(tx *sql.Tx) error {
		var err error
		err = service.DeleteAccessByToken(tx, code)
		return err
	})
	return err
}

func (s *MySQLStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	Log.Notice("OAuth2, load refresh: %s\n", code)
	var a *osin.AccessData
	err := Transact(DB, func(tx *sql.Tx) error {
		var err error
		a, err = service.GetAccessByRefresh(tx, code)
		return err
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *MySQLStorage) RemoveRefresh(code string) error {
	Log.Notice("OAuth2, remove refresh: %s\n", code)
	return nil
}
