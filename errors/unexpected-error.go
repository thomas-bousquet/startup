package errors

type UnexpectedError struct {
	BaseError
}

func NewUnexpectedError() UnexpectedError {
	return UnexpectedError{BaseError: BaseError{Message: "unexpected-error"}}
}
