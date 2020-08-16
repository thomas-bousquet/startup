package errors

type BaseError struct {
	Message string           `json:"message"`
}

func (e BaseError) Error() string {
	return e.Message
}