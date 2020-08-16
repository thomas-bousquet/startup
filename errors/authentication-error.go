package errors

type AuthenticationError struct {
	BaseError
}

func NewAuthenticationError() AuthenticationError {
	return AuthenticationError{BaseError: BaseError{Message: "authentication-error"}}
}
