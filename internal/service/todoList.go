package service

import (
	"context"

	"github.com/sarulabs/di"

	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/IgorAndrade/go-boilerplate/internal/repository"
)

const TODO_LIST = "todoListService"

func Define(b *di.Builder) {
	b.Add(di.Def{
		Name:  TODO_LIST,
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			r := ctn.Get(repository.TODO_LIST).(repository.TodoList)
			return NewTodoListService(r), nil
		},
	})
}

type TodoList interface {
	Create(context.Context, *model.TodoList) error
	GetAll(context.Context) ([]model.TodoList, error)
}

func NewTodoListService(r repository.TodoList) TodoList {
	return &TodoListImp{
		repository: r,
	}
}

type TodoListImp struct {
	repository repository.TodoList
}

func (t TodoListImp) Create(ctx context.Context, todoList *model.TodoList) error {
	return t.repository.Create(ctx, todoList)
}
func (t TodoListImp) GetAll(ctx context.Context) ([]model.TodoList, error) {
	return t.repository.GetAll(ctx)
}
