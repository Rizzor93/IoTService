package transport

import (
	"IoT_Service/server/services/restApi/models"
	"encoding/json"
	grpcCode "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"log"
	"net/http"
)

const (
	notFoundErrType         = "resource_not_found_error"
	dataValidationErrType   = "data_validation_error"
	formatValidationErrType = "format_validation_error"
	invalidJSONErrType      = "invalid_json_error"
	serviceErrType          = "service_error"
	serviceErrGrpcType      = "service_error_grpc_server"
)

// JsonEncoder
func JsonEncoder(w http.ResponseWriter, code int) *json.Encoder {
	w.Header().Set("Content-Status", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w)
}

// SendJSON json response to the client
func SendJSON(w http.ResponseWriter, response interface{}, code int) {
	encoder := JsonEncoder(w, code)
	if err := encoder.Encode(response); err != nil {
		log.Printf("could not encode error: %v", err)
	}
}

// SendJSONStream json response to the client
func SendJSONStream(w *json.Encoder, response interface{}) {
	if err := w.Encode(response); err != nil {
		log.Printf("could not encode error: %v", err)
	}
}

// SendError a json error to the client
func SendError(w http.ResponseWriter, err error) {
	e := toHTTPError(err)
	encoder := JsonEncoder(w, e.Code)
	if err := encoder.Encode(e); err != nil {
		log.Printf("could not encode error: %v", err)
	}
}

// toHTTPError convert an error to HTTPError
func toHTTPError(err error) models.HTTPError {
	resErr := models.HTTPError{Message: err.Error()}
	switch e := err.(type) {
	case models.HTTPError:
		return e
	case models.NotFoundError:
		resErr.Code = http.StatusNotFound
		resErr.Type = notFoundErrType
	case models.FormatValidationError:
		resErr.Code = http.StatusBadRequest
		resErr.Type = formatValidationErrType
	case models.DataValidationError:
		resErr.Code = http.StatusBadRequest
		resErr.Type = dataValidationErrType
	case models.InvalidJSONError:
		resErr.Code = http.StatusBadRequest
		resErr.Type = invalidJSONErrType
	default:
		resErr.Code = http.StatusInternalServerError
		resErr.Type = serviceErrType
		resErr.Message = "Internal Server Error"
	}

	log.Printf("Error on 'apiServer': \n\tErrorMessage: '%v'  httpCode: '%v' type:'%v' ", resErr.Message, resErr.Code, resErr.Type)
	return resErr
}

// GrpcStatusToHTTPError convert grpcStatus to httpError
func GrpcStatusToHTTPError(err error) models.HTTPError {
	status, _ := grpcStatus.FromError(err)
	resErr := models.HTTPError{Message: status.Message()}

	switch status.Code() {
	case grpcCode.NotFound:
		resErr.Code = http.StatusNotFound
		resErr.Type = notFoundErrType
	case grpcCode.InvalidArgument:
		resErr.Code = http.StatusBadRequest
		resErr.Type = formatValidationErrType
	case grpcCode.Unavailable:
		resErr.Code = http.StatusInternalServerError
		resErr.Type = serviceErrGrpcType
		resErr.Message = "Internal gRPC-Server Error"
	default:
		resErr.Code = http.StatusInternalServerError
		resErr.Type = serviceErrGrpcType
		resErr.Message = "Internal Server Error"
	}

	return resErr
}
