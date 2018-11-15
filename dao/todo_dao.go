package dao

import (
	"log"
	"todo-go/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TodoDAO representation
type TodoDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION name
	COLLECTION = "todo"
)

//Connect is establish connection
func (td *TodoDAO) Connect() {
	session, err := mgo.Dial(td.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(td.Database)
}

// FindAll is to get all Todo list
func (td *TodoDAO) FindAll() ([]*models.Todo, error) {
	var todo []*models.Todo
	err := db.C(COLLECTION).Find(nil).All(&todo)
	return todo, err
}

// FindByID is to get one todo list
func (td *TodoDAO) FindByID(id string) (*models.Todo, error) {
	var todo *models.Todo
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&todo)
	return todo, err
}

// Insert is to create new todo list
func (td *TodoDAO) Insert(todo models.Todo) error {
	err := db.C(COLLECTION).Insert(&todo)
	return err
}

// Update is to update the todo records
func (td *TodoDAO) Update(todo models.Todo) error {
	err := db.C(COLLECTION).UpdateId(todo.ID, &todo)
	return err
}

// Delete is to delete the todo list
func (td *TodoDAO) Delete(todo models.Todo) error {
	err := db.C(COLLECTION).RemoveId(todo.ID)
	return err
}
