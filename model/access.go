package model

import (
	"time"
)

type AccessData struct {
	UID             string
	ClientID        string
	AuthorizeDataID string
	AccessDataID    string
	UserID          string
	AccessToken     string
	RefreshToken    string
	ExpiresIn       int32
	Scope           string
	RedirectUri     string
	CreatedAt       time.Time
}
