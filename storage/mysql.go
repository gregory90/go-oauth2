package storage

import (
	. "bitbucket.org/pqstudio/go-webutils/logger"
	"errors"
	"github.com/RangelReale/osin"

	"bitbucket.org/pqstudio/go-oauth2/service"
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
	c, err := service.GetClientByID(id)
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

	err := service.CreateClient(c)
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

	err := service.CreateAccess(data)
	return err
}

func (s *MySQLStorage) LoadAccess(code string) (*osin.AccessData, error) {
	Log.Notice("OAuth2, load access: %s\n", code)

	a, err := service.GetAccessByToken(code)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *MySQLStorage) RemoveAccess(code string) error {
	Log.Notice("OAuth2, remove access: %s\n", code)

	err := service.DeleteAccessByToken(code)
	return err
}

func (s *MySQLStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	Log.Notice("OAuth2, load refresh: %s\n", code)
	a, err := service.GetAccessByRefresh(code)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *MySQLStorage) RemoveRefresh(code string) error {
	Log.Notice("OAuth2, remove refresh: %s\n", code)
	return nil
}
