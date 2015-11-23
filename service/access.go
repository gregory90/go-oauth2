package service

import (
	"database/sql"

	"github.com/RangelReale/osin"

	"github.com/gregory90/go-oauth2/datastore"
	"github.com/gregory90/go-oauth2/model"
	"github.com/gregory90/go-webutils"

	. "github.com/gregory90/go-webutils/logger"
)

func GetAccessByToken(tx *sql.Tx, token string) (*osin.AccessData, error) {
	access, err := datastore.GetAccessByToken(tx, token)
	if err != nil {
		return nil, err
	}

	client, err := GetClientByUID(tx, access.ClientID)
	if err != nil {
		return nil, err
	}

	a := &osin.AccessData{
		Client:        client,
		AuthorizeData: nil,
		AccessData:    nil,
		AccessToken:   access.AccessToken,
		RefreshToken:  access.RefreshToken,
		ExpiresIn:     access.ExpiresIn,
		Scope:         access.Scope,
		RedirectUri:   access.RedirectUri,
		CreatedAt:     access.CreatedAt,
	}

	return a, nil
}

func CreateAccess(tx *sql.Tx, data *osin.AccessData) error {
	client, err := datastore.GetClientByID(tx, data.Client.GetId())
	if err != nil {
		return err
	}

	access := &model.AccessData{
		UID:             utils.NewUUID(),
		ClientID:        client.UID,
		AuthorizeDataID: "",
		AccessDataID:    "",
		AccessToken:     data.AccessToken,
		RefreshToken:    data.RefreshToken,
		ExpiresIn:       data.ExpiresIn,
		Scope:           data.Scope,
		RedirectUri:     data.RedirectUri,
		CreatedAt:       data.CreatedAt,
	}

	Log.Debug("%+v", access)
	err = datastore.CreateAccess(tx, access)
	return err
}

func DeleteAccessByToken(tx *sql.Tx, token string) error {
	err := datastore.DeleteAccessByToken(tx, token)
	return err
}

func DeleteAccessByUserUID(tx *sql.Tx, userUID string) error {
	err := datastore.DeleteAccessByUserUID(tx, userUID)
	return err
}

func GetAccessByRefresh(tx *sql.Tx, token string) (*osin.AccessData, error) {
	access, err := datastore.GetAccessByRefresh(tx, token)
	if err != nil {
		return nil, err
	}

	client, err := GetClientByUID(tx, access.ClientID)
	if err != nil {
		return nil, err
	}

	a := &osin.AccessData{
		Client:        client,
		AuthorizeData: nil,
		AccessData:    nil,
		AccessToken:   access.AccessToken,
		RefreshToken:  access.RefreshToken,
		ExpiresIn:     access.ExpiresIn,
		Scope:         access.Scope,
		RedirectUri:   access.RedirectUri,
		CreatedAt:     access.CreatedAt,
	}

	return a, nil
}

func GetAccessByRefreshModel(tx *sql.Tx, token string) (*model.AccessData, error) {
	access, err := datastore.GetAccessByRefresh(tx, token)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func GetAccessByUser(tx *sql.Tx, userUID string) (*model.AccessData, error) {
	access, err := datastore.GetAccessByUser(tx, userUID)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func UpdateAccessByToken(tx *sql.Tx, token string, userUID string) error {
	err := datastore.UpdateAccessByToken(tx, token, userUID)
	return err
}
