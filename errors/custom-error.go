package errors

type CustomError struct {
	HttpCode        int    `json:"-"`
	Message         string `json:"message"`
	LocalizationKey string `json:"localization_key"`
	Metadata        interface{} `json:"metadata,omitempty"`
}

func NewBaseError(httpCode int, localizationKey string, message string, metadata interface{}) CustomError {
	return CustomError{
		HttpCode:        httpCode,
		Message:         message,
		LocalizationKey: localizationKey,
		Metadata: metadata,
	}
}

func (e CustomError) Error() string {
	return e.Message
}