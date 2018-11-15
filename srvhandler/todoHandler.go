package srvhandler

import (
	"encoding/json"
	"net/http"
	"todo-go/dao"
	"todo-go/models"
	"todo-go/mthdroutr"
	"todo-go/utils"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var tododao dao.TodoDAO

// TodoHandler handles todo request
type TodoHandler struct {
	BaseHandler
}

func (td *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rresponse := mthdroutr.RouteAPICall(td, r)
	rresponse.RenderResponse(w)
}

// Get todo get list
func (td *TodoHandler) Get(r *http.Request) utils.ResponseUtil {
	params := mux.Vars(r)
	id, present := params["id"]
	if present {
		todo, err := tododao.FindByID(id)
		if err != nil {
			return utils.BadRequest("Invalid request Id")
		}
		return utils.OK200("OK response", todo)
	}
	todo, err := tododao.FindAll()
	if err != nil {
		return utils.InternalServerError(err.Error())
	}
	return utils.OK200("OK response", todo)
}

// Post ...
func (td *TodoHandler) Post(r *http.Request) utils.ResponseUtil {
	defer r.Body.Close()
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return utils.BadRequest("Invalid Payload")
	}
	todo.ID = bson.NewObjectId()
	if err := tododao.Insert(todo); err != nil {
		return utils.InternalServerError("Internal Server Error")
	}
	return utils.OK200("OK response", todo)
}

// Put ...
func (td *TodoHandler) Put(r *http.Request) utils.ResponseUtil {
	defer r.Body.Close()
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return utils.BadRequest("Invalid Payload")
	}
	if err := tododao.Update(todo); err != nil {
		return utils.BadRequest("Invalid todo id")
	}
	return utils.OK200("OK response", todo)
}

// Delete ...
func (td *TodoHandler) Delete(r *http.Request) utils.ResponseUtil {
	defer r.Body.Close()
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return utils.BadRequest("Invalid Payload")
	}
	if err := tododao.Delete(todo); err != nil {
		return utils.BadRequest("Invalid todo id")
	}
	return utils.OK200("OK response", todo)
}
