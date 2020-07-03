package mock

import (
	"context"

	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/IgorAndrade/go-boilerplate/internal/repository"
)

type TodoListMock struct {
	CreateFn func(context.Context, *model.TodoList) error
	GetAllFn func(context.Context) ([]model.TodoList, error)
}

func NewTodoListMock() repository.TodoList {
	return &TodoListMock{}
}

func (m *TodoListMock) Create(c context.Context, t *model.TodoList) error {
	return m.CreateFn(c, t)
}

func (m *TodoListMock) GetAll(c context.Context) ([]model.TodoList, error) {
	return m.GetAllFn(c)
}
