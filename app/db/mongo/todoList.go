package mongo

import (
	"context"

	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/IgorAndrade/go-boilerplate/internal/repository"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type TodoList struct {
	*mongo
}

func newTodoList(s *mgo.Session) repository.TodoList {
	return &TodoList{
		mongo: &mongo{session: s, db: DB_NAME, collectionName: "todoList"},
	}
}

func (t TodoList) Create(c context.Context, tl *model.TodoList) error {
	return t.insert(tl)
}

func (t TodoList) GetAll(context.Context) ([]model.TodoList, error) {
	list := make([]model.TodoList, 0)
	err := t.findAllByFilter(bson.M{}, &list)
	return list, err
}
