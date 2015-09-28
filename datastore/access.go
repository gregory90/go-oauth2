package datastore

import (
	. "bitbucket.org/pqstudio/go-oauth2/db"

	"bitbucket.org/pqstudio/go-oauth2/model"
)

const (
	accessTable string = "access_data"
)

func GetAccessByToken(token string) (*model.AccessData, error) {
	access := &model.AccessData{}

	// get access
	err := DB.QueryRow("SELECT lower(hex(uid)), lower(hex(clientID)), lower(hex(userID)),accessToken, refreshToken, expiresIn, scope, redirectURI, createdAt FROM "+accessTable+" WHERE accessToken = ?", token).Scan(&access.UID, &access.ClientID, &access.UserID, &access.AccessToken, &access.RefreshToken, &access.ExpiresIn, &access.Scope, &access.RedirectUri, &access.CreatedAt)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func GetAccessByUser(userUID string) (*model.AccessData, error) {
	access := &model.AccessData{}

	// get access
	err := DB.QueryRow("SELECT lower(hex(uid)), lower(hex(clientID)), lower(hex(userID)),accessToken, refreshToken, expiresIn, scope, redirectURI, createdAt FROM "+accessTable+" WHERE userID = unhex(?) ORDER BY createdAt DESC LIMIT 1", userUID).Scan(&access.UID, &access.ClientID, &access.UserID, &access.AccessToken, &access.RefreshToken, &access.ExpiresIn, &access.Scope, &access.RedirectUri, &access.CreatedAt)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func CreateAccess(data *model.AccessData) error {
	stmt, err := DB.Prepare("INSERT " + accessTable + " SET uid=unhex(?),clientID=unhex(?),authorizeDataID=unhex(?),accessDataID=unhex(?),userID=unhex(?),accessToken=?,refreshToken=?,expiresIn=?,scope=?,redirectURI=?,createdAt=?")
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

func DeleteAccessByToken(token string) error {
	stmt, err := DB.Prepare("DELETE FROM " + accessTable + " WHERE accessToken=?")
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

func GetAccessByRefresh(token string) (*model.AccessData, error) {
	access := &model.AccessData{}

	// get access
	err := DB.QueryRow("SELECT lower(hex(uid)), lower(hex(clientID)), lower(hex(userID)),accessToken, refreshToken, expiresIn, scope, redirectURI, createdAt FROM "+accessTable+" WHERE refreshToken = ?", token).Scan(&access.UID, &access.ClientID, &access.UserID, &access.AccessToken, &access.RefreshToken, &access.ExpiresIn, &access.Scope, &access.RedirectUri, &access.CreatedAt)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func UpdateAccessByToken(token string, userUID string) error {
	stmt, err := DB.Prepare("UPDATE " + accessTable + " SET userID=unhex(?) WHERE accessToken=?")
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
