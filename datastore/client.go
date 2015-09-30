package datastore

import (
	"bitbucket.org/pqstudio/go-oauth2/model"
)

const (
	clientTable string = "client_data"
)

func GetClientByUID(tx *sql.Tx, id string) (*model.Client, error) {
	client := &model.Client{}

	// get client
	err := tx.QueryRow("SELECT lower(hex(uid)), clientID, clientSecret, redirectURI FROM "+clientTable+" WHERE uid = unhex(?)", id).Scan(&client.UID, &client.Id, &client.Secret, &client.RedirectUri)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetClientByID(tx *sql.Tx, id string) (*model.Client, error) {
	client := &model.Client{}

	// get client
	err := tx.QueryRow("SELECT lower(hex(uid)), clientID, clientSecret, redirectURI FROM "+clientTable+" WHERE clientID = ?", id).Scan(&client.UID, &client.Id, &client.Secret, &client.RedirectUri)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateClient(tx *sql.Tx, client *model.Client) error {
	stmt, err := tx.Prepare("INSERT " + clientTable + " SET uid=unhex(?),clientID=?,clientSecret=?,redirectURI=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.UID, client.Id, client.Secret, client.RedirectUri)
	return err
}
