package service

import (
	"bitbucket.org/pqstudio/go-webutils"

	"bitbucket.org/pqstudio/go-oauth2/datastore"
	"bitbucket.org/pqstudio/go-oauth2/model"

	"github.com/RangelReale/osin"

	//. "bitbucket.org/pqstudio/go-webutils/logger"
)

func GetClientByUID(tx *sql.Tx, uid string) (*osin.DefaultClient, error) {
	client, err := datastore.GetClientByUID(tx, uid)
	if err != nil {
		return nil, err
	}

	c := &osin.DefaultClient{
		Id:          client.Id,
		Secret:      client.Secret,
		RedirectUri: client.RedirectUri,
	}

	return c, nil
}

func GetClientByID(tx *sql.Tx, id string) (*osin.DefaultClient, error) {
	client, err := datastore.GetClientByID(tx, id)
	if err != nil {
		return nil, err
	}

	c := &osin.DefaultClient{
		Id:          client.Id,
		Secret:      client.Secret,
		RedirectUri: client.RedirectUri,
	}

	return c, nil
}

func CreateClient(tx *sql.Tx, client *osin.DefaultClient) error {
	c := &model.Client{
		UID:         utils.NewUUID(),
		Id:          client.Id,
		Secret:      client.Secret,
		RedirectUri: client.RedirectUri,
	}
	err := datastore.CreateClient(tx, c)
	return err
}
