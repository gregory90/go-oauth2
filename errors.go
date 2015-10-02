package oauth2

type AuthenticationError struct {
	ErrorType string `json:"error"`
}

func (e *AuthenticationError) Error() string {
	return e.ErrorType
}

func NewAuthenticationError(text string) error {
	return &AuthenticationError{text}
}
