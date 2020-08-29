package errors

type BaseError struct {
	HttpCode        int    `json:"-"`
	Message         string `json:"message"`
	LocalizationKey string `json:"localization_key"`
}

func NewBaseError(httpCode int, localizationKey string, message string) BaseError {
	return BaseError{
		HttpCode:        httpCode,
		Message:         message,
		LocalizationKey: localizationKey,
	}
}

func (e BaseError) Error() string {
	return e.Message
}
