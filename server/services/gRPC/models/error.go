package models

import (
	grpcCode "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"log"
)

type OK struct {
	Message string
}

func (e OK) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
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

// InternalError represents the error returned when internal errors (grpc/db)
type InternalError struct {
	Message string
}

func (e InternalError) Error() string {
	return e.Message
}

// CreateGRPCStatus convert an error to grpcStatus
func CreateGRPCStatus(err error) error {
	code := grpcCode.Code(0)
	msg := err.Error()
	switch err.(type) {
	case OK:
		code = grpcCode.OK
		return grpcStatus.New(code, msg).Err()
	case NotFoundError:
		code = grpcCode.NotFound
	case FormatValidationError:
		code = grpcCode.InvalidArgument
	case DataValidationError:
		code = grpcCode.InvalidArgument
	default:
		code = grpcCode.Internal
	}

	log.Printf("Error on 'grpcServer': ErrorMessage: '%v'  grpcCode: '%v' ", msg, code)
	return grpcStatus.New(code, msg).Err()

}
