package handlers

import "net/http"

// Code 200 Success
type RequestOK struct {
	Code    int    `json:"Code" example:"200"`
	Message string `json:"Message" example:"Request succeeded."`
}

func Success200(w *http.ResponseWriter, msg string) RequestOK {
	(*w).WriteHeader(http.StatusOK)
	return RequestOK{Code: 200, Message: msg}
}

// Code 400 Bad Request
type BadRequest struct {
	Code    int    `json:"Code" example:"400"`
	Message string `json:"Message" example:"Unable to process request."`
}

func Error400(w *http.ResponseWriter, msg string) BadRequest {
	(*w).WriteHeader(http.StatusBadRequest)
	return BadRequest{Code: 400, Message: msg}
}

// Code 403 Forbidden
type Forbidden struct {
	Code    int    `json:"Code" example:"403"`
	Message string `json:"Message" example:"Access to the requested resource is forbidden."`
}

func Error403(w *http.ResponseWriter, msg string) Forbidden {
	(*w).WriteHeader(http.StatusForbidden)
	return Forbidden{Code: 403, Message: msg}
}

// Code 404 Not Found
type NotFound struct {
	Code    int    `json:"Code" example:"404"`
	Message string `json:"Message" example:"The requested resource was not found."`
}

func Error404(w *http.ResponseWriter, msg string) NotFound {
	(*w).WriteHeader(http.StatusNotFound)
	return NotFound{Code: 404, Message: msg}
}

// Code 500 Internal Server Error
type InternalServerError struct {
	Code    int    `json:"Code" example:"500"`
	Message string `json:"Message" example:"A server error occurred. Please try again later."`
}

func Error500(w *http.ResponseWriter, msg string) InternalServerError {
	(*w).WriteHeader(http.StatusInternalServerError)
	return InternalServerError{Code: 500, Message: msg}
}

// Code 501 Not Implemented
type NotImplemented struct {
	Code    int    `json:"Code" example:"501"`
	Message string `json:"Message" example:"This operation has not yet been implemented."`
}

func Error501(w *http.ResponseWriter, msg string) NotImplemented {
	(*w).WriteHeader(http.StatusNotImplemented)
	return NotImplemented{Code: 501, Message: msg}
}
