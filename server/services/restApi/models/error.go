package models

// HTTPError represent an http error to be returned to the client
type HTTPError struct {
	Code    int    `json:"code,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e HTTPError) Error() string {
	return e.Message
}

// FormatValidationError represents the error returned in case the request body was
// a wrong format which the server cannot work with
type FormatValidationError struct {
	Message string
}

func (e FormatValidationError) Error() string {
	return e.Message
}

// DataValidationError represents the error returned when the format of request
// is valid but the data is invalid
type DataValidationError struct {
	Message string
}

func (e DataValidationError) Error() string {
	return e.Message
}

// InvalidJSONError represents the error returned when request body contains invalid JSON
type InvalidJSONError struct {
	Message string
}

func (e InvalidJSONError) Error() string {
	return e.Message
}

// NotFoundError represents the error returned in case a resource or route is not found
type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	if e.Message == "" {
		return "resource not found"
	}
	return e.Message
}
