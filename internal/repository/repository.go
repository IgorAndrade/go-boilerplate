package repository

import (
	"context"

	"github.com/IgorAndrade/go-boilerplate/internal/model"
)

const TODO_LIST = "tudoListRepo"

type TodoList interface {
	Create(context.Context, *model.TodoList) error
	GetAll(context.Context) ([]model.TodoList, error)
}
