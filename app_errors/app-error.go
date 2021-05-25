package app_errors

type AppError struct {
	Key              string      `json:"key"`
	Message          *string     `json:"message,omitempty"`
	InternalMetadata interface{} `json:"metadata,omitempty"`
}

func NewAppError(key string, message *string, metadata map[string]interface{}) error {
	return &AppError{
		Key:              key,
		Message:          message,
		InternalMetadata: metadata,
	}
}

func (e *AppError) Error() string {
	return e.Key
}
