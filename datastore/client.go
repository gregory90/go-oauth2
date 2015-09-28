package datastore

import (
	. "bitbucket.org/pqstudio/go-oauth2/db"

	"bitbucket.org/pqstudio/go-oauth2/model"
)

const (
	clientTable string = "client_data"
)

func GetClientByUID(id string) (*model.Client, error) {
	client := &model.Client{}

	// get client
	err := DB.QueryRow("SELECT lower(hex(uid)), clientID, clientSecret, redirectURI FROM "+clientTable+" WHERE uid = unhex(?)", id).Scan(&client.UID, &client.Id, &client.Secret, &client.RedirectUri)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetClientByID(id string) (*model.Client, error) {
	client := &model.Client{}

	// get client
	err := DB.QueryRow("SELECT lower(hex(uid)), clientID, clientSecret, redirectURI FROM "+clientTable+" WHERE clientID = ?", id).Scan(&client.UID, &client.Id, &client.Secret, &client.RedirectUri)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateClient(client *model.Client) error {
	stmt, err := DB.Prepare("INSERT " + clientTable + " SET uid=unhex(?),clientID=?,clientSecret=?,redirectURI=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.UID, client.Id, client.Secret, client.RedirectUri)
	return err
}
