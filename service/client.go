package service

import (
	"bitbucket.org/pqstudio/go-webutils"

	"bitbucket.org/pqstudio/go-oauth2/datastore"
	"bitbucket.org/pqstudio/go-oauth2/model"

	"github.com/RangelReale/osin"

	//. "bitbucket.org/pqstudio/go-webutils/logger"
)

func GetClientByUID(uid string) (*osin.DefaultClient, error) {
	client, err := datastore.GetClientByUID(uid)
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

func GetClientByID(id string) (*osin.DefaultClient, error) {
	client, err := datastore.GetClientByID(id)
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

func CreateClient(client *osin.DefaultClient) error {
	c := &model.Client{
		UID:         utils.NewUUID(),
		Id:          client.Id,
		Secret:      client.Secret,
		RedirectUri: client.RedirectUri,
	}
	err := datastore.CreateClient(c)
	return err
}
