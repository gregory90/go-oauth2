package oauth2

type AuthenticationError struct {
	ErrorType string `json:"error"`
}

func (e *InvalidAccessToken) Error() string {
	return e.ErrorType
}

func InvalidAccessTokenError(text string) error {
	return &AuthenticationError{text}
}
