package utils

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// EmptyStruct is to initialize empty struct
type EmptyStruct struct {
}

// ResponseUtil is to represent response
type ResponseUtil struct {
	Code     int
	Message  string
	Response interface{}
	Headers  map[string]string
}

// RenderResponse is to render response
func (s *ResponseUtil) RenderResponse(w http.ResponseWriter) {
	if s.Headers == nil {
		s.Headers = map[string]string{"Content-Type": "application/json",
			"Access-Control-Allow-Headers": "*",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*"}
		for h, val := range s.Headers {
			w.Header().Set(h, val)
		}
	}
	responseFormatted := bson.M{
		"responseData": s.Response,
		"message":      s.Message}
	response, _ := json.MarshalIndent(responseFormatted, "", "")

	w.WriteHeader(s.Code)
	w.Write(response)
}

// OK200 is for 200 ok response
func OK200(message string, response interface{}) ResponseUtil {
	return ResponseUtil{http.StatusOK, message, response, nil}
}

// BadRequest is for 400 bad request
func BadRequest(message string) ResponseUtil {
	var empty EmptyStruct
	return ResponseUtil{http.StatusBadRequest, message, empty, nil}
}

// InternalServerError is for 500 internal server error
func InternalServerError(message string) ResponseUtil {
	var empty EmptyStruct
	return ResponseUtil{http.StatusInternalServerError, message, empty, nil}
}

// MethodNotAllowed ...
func MethodNotAllowed(message string) ResponseUtil {
	var empty EmptyStruct
	return ResponseUtil{http.StatusMethodNotAllowed, message, empty, nil}
}

// MethodNotImplemented ...
func MethodNotImplemented() ResponseUtil {
	var empty EmptyStruct
	return ResponseUtil{http.StatusNotImplemented, "Method not implemented", empty, nil}
}
