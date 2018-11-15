package srvhandler

import (
	"net/http"
	"todo-go/utils"
)

// BaseHandler ...
type BaseHandler struct {
}

// Get ...
func (b *BaseHandler) Get(r *http.Request) utils.ResponseUtil {
	return utils.MethodNotImplemented()
}

// GetOne ...
func (b *BaseHandler) GetOne(r *http.Request, id string) utils.ResponseUtil {
	return utils.MethodNotImplemented()
}

// Post ...
func (b *BaseHandler) Post(r *http.Request) utils.ResponseUtil {
	return utils.MethodNotImplemented()
}

// Put ...
func (b *BaseHandler) Put(r *http.Request) utils.ResponseUtil {
	return utils.MethodNotImplemented()
}

// Patch ...
func (b *BaseHandler) Patch(r *http.Request) utils.ResponseUtil {
	return utils.MethodNotImplemented()
}

// Delete ...
func (b *BaseHandler) Delete(r *http.Request) utils.ResponseUtil {
	return utils.MethodNotImplemented()
}

// Options ...
func (b *BaseHandler) Options(r *http.Request) utils.ResponseUtil {
	return utils.MethodNotImplemented()
}
