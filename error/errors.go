package error

type RequestError struct {
	StatusCode int
}

type ValidationErrorItem struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Reason string `json:"reason"`
}