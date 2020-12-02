package errors

type Error struct {
	HttpCode        int    `json:"-"`
	Message         string `json:"message"`
	LocalizationKey string `json:"localization_key"`
	Metadata        interface{} `json:"metadata,omitempty"`
}

func NewError(httpCode int, localizationKey string, message string, metadata interface{}) *Error {
	return &Error{
		HttpCode:        httpCode,
		Message:         message,
		LocalizationKey: localizationKey,
		Metadata: metadata,
	}
}

func (e Error) Error() string {
	return e.Message
}