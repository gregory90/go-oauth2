package datastore

import (
	"database/sql"

	"bitbucket.org/pqstudio/go-oauth2/model"
)

const (
	accessTable string = "access_data"
)

func GetAccessByToken(tx *sql.Tx, token string) (*model.AccessData, error) {
	access := &model.AccessData{}

	// get access
	err := tx.QueryRow("SELECT lower(hex(uid)), lower(hex(clientID)), lower(hex(userID)),accessToken, refreshToken, expiresIn, scope, redirectURI, createdAt FROM "+accessTable+" WHERE accessToken = ?", token).Scan(&access.UID, &access.ClientID, &access.UserID, &access.AccessToken, &access.RefreshToken, &access.ExpiresIn, &access.Scope, &access.RedirectUri, &access.CreatedAt)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func GetAccessByUser(tx *sql.Tx, userUID string) (*model.AccessData, error) {
	access := &model.AccessData{}

	// get access
	err := tx.QueryRow("SELECT lower(hex(uid)), lower(hex(clientID)), lower(hex(userID)),accessToken, refreshToken, expiresIn, scope, redirectURI, createdAt FROM "+accessTable+" WHERE userID = unhex(?) ORDER BY createdAt DESC LIMIT 1", userUID).Scan(&access.UID, &access.ClientID, &access.UserID, &access.AccessToken, &access.RefreshToken, &access.ExpiresIn, &access.Scope, &access.RedirectUri, &access.CreatedAt)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func CreateAccess(tx *sql.Tx, data *model.AccessData) error {
	stmt, err := tx.Prepare("INSERT " + accessTable + " SET uid=unhex(?),clientID=unhex(?),authorizeDataID=unhex(?),accessDataID=unhex(?),userID=unhex(?),accessToken=?,refreshToken=?,expiresIn=?,scope=?,redirectURI=?,createdAt=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.UID, data.ClientID, nil, nil, nil, data.AccessToken, data.RefreshToken, data.ExpiresIn, data.Scope, data.RedirectUri, data.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccessByToken(tx *sql.Tx, token string) error {
	stmt, err := tx.Prepare("DELETE FROM " + accessTable + " WHERE accessToken=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	if err != nil {
		return err
	}

	return nil
}

func GetAccessByRefresh(tx *sql.Tx, token string) (*model.AccessData, error) {
	access := &model.AccessData{}

	// get access
	err := tx.QueryRow("SELECT lower(hex(uid)), lower(hex(clientID)), lower(hex(userID)),accessToken, refreshToken, expiresIn, scope, redirectURI, createdAt FROM "+accessTable+" WHERE refreshToken = ?", token).Scan(&access.UID, &access.ClientID, &access.UserID, &access.AccessToken, &access.RefreshToken, &access.ExpiresIn, &access.Scope, &access.RedirectUri, &access.CreatedAt)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func UpdateAccessByToken(tx *sql.Tx, token string, userUID string) error {
	stmt, err := tx.Prepare("UPDATE " + accessTable + " SET userID=unhex(?) WHERE accessToken=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userUID, token)
	if err != nil {
		return err
	}

	return nil
}
