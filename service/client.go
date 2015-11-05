package service

import (
	"database/sql"

	"github.com/RangelReale/osin"

	"github.com/gregory90/go-webutils"

	"github.com/gregory90/go-oauth2/datastore"
	"github.com/gregory90/go-oauth2/model"
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
