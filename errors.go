package oauth2

type InvalidAccessToken struct {
	ErrorType string `json:"error"`
}

func (e *InvalidAccessToken) Error() string {
	return e.ErrorType
}

func InvalidAccessTokenError() error {
	return &InvalidAccessToken{"access_token_invalid"}
}
