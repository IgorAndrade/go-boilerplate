package service

import (
	"context"

	"github.com/IgorAndrade/go-boilerplate/internal/model"
)

type TodoListMock struct {
	CreateFn func(context.Context, *model.TodoList) error
	GetAllFn func(context.Context) ([]model.TodoList, error)
}

func (m *TodoListMock) Create(c context.Context, tl *model.TodoList) error {
	return m.CreateFn(c, tl)
}
func (m *TodoListMock) GetAll(c context.Context) ([]model.TodoList, error) {
	return m.GetAllFn(c)
}
