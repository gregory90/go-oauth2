package service

import (
	"bitbucket.org/pqstudio/go-webutils"

	"bitbucket.org/pqstudio/go-oauth2/datastore"
	"bitbucket.org/pqstudio/go-oauth2/model"

	"github.com/RangelReale/osin"

	. "bitbucket.org/pqstudio/go-webutils/logger"
)

func GetAccessByToken(token string) (*osin.AccessData, error) {
	access, err := datastore.GetAccessByToken(token)
	if err != nil {
		return nil, err
	}

	client, err := GetClientByUID(access.ClientID)
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

func CreateAccess(data *osin.AccessData) error {
	client, err := datastore.GetClientByID(data.Client.GetId())
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
	err = datastore.CreateAccess(access)
	return err
}

func DeleteAccessByToken(token string) error {
	err := datastore.DeleteAccessByToken(token)
	return err
}

func GetAccessByRefresh(token string) (*osin.AccessData, error) {
	access, err := datastore.GetAccessByRefresh(token)
	if err != nil {
		return nil, err
	}

	client, err := GetClientByUID(access.ClientID)
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

func GetAccessByRefreshModel(token string) (*model.AccessData, error) {
	access, err := datastore.GetAccessByRefresh(token)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func GetAccessByUser(userUID string) (*model.AccessData, error) {
	access, err := datastore.GetAccessByUser(userUID)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func UpdateAccessByToken(token string, userUID string) error {
	err := datastore.UpdateAccessByToken(token, userUID)
	return err
}
