package mongo

import (
	"context"

	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/IgorAndrade/go-boilerplate/internal/repository"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sarulabs/di"
)

type TodoList struct {
	*mongo
}

func Init() {
	config.AddDef(di.Def{
		Name:  repository.TODO_LIST,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			s := ctn.Get(MONGO).(*mgo.Session)
			return TodoList{
				mongo: &mongo{
					session:        s,
					db:             DB_NAME,
					collectionName: "todoList",
				},
			}, nil
		},
	})
}

func (t TodoList) Create(c context.Context, tl model.TodoList) error {
	return t.insert(&tl)
}

func (t TodoList) GetAll(context.Context) ([]model.TodoList, error) {
	list := make([]model.TodoList, 0)
	err := t.findAllByFilter(bson.M{}, &list)
	return list, err
}
