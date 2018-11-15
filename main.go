package main

import (
	"log"
	"net/http"
	"todo-go/config"
	"todo-go/dao"
	"todo-go/srvhandler"
	"todo-go/utils"

	"github.com/gorilla/mux"
)

var tododao = dao.TodoDAO{}
var dbconfig = config.DBConfig{}
var srvresponse = utils.ResponseUtil{}

func init() {
	dbconfig.Read()
	tododao.Server = dbconfig.Server
	tododao.Database = dbconfig.Database
	tododao.Connect()

}
func main() {
	r := mux.NewRouter()
	todoHandler := &srvhandler.TodoHandler{}
	r.Handle("/todo", todoHandler)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Not Working")
	}
}
