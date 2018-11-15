package mthdroutr

import (
	"fmt"
	"net/http"
	"todo-go/utils"
)

// ServiceAPIHandler contains all restfull methods
type ServiceAPIHandler interface {
	GetOne(r *http.Request, id string) utils.ResponseUtil
	Get(r *http.Request) utils.ResponseUtil
	Post(r *http.Request) utils.ResponseUtil
	Put(r *http.Request) utils.ResponseUtil
	Patch(r *http.Request) utils.ResponseUtil
	Delete(r *http.Request) utils.ResponseUtil
	Options(r *http.Request) utils.ResponseUtil
}

// RouteAPICall to route api calls
func RouteAPICall(sah ServiceAPIHandler, r *http.Request) utils.ResponseUtil {
	switch r.Method {
	case "GET":
		return sah.Get(r)
	case "POST":
		return sah.Post(r)
	case "PUT":
		return sah.Put(r)
	case "PATCH":
		return sah.Patch(r)
	case "DELETE":
		return sah.Delete(r)
	case "OPTIONS":
		return sah.Options(r)

	}
	return utils.ResponseUtil{
		Code:     http.StatusMethodNotAllowed,
		Response: fmt.Sprintf("{\"ResponseData\" : \"%s \"}", r.Method),
		Message:  "Method not allowed",
		Headers:  nil}
}
