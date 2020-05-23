package route

import (
	"net/http"

	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/IgorAndrade/go-boilerplate/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

func create(c echo.Context, ctn di.Container) error {
	todoList := model.TodoList{}
	err := c.Bind(&todoList)
	if err != nil {
		return err
	}
	r := ctn.Get(repository.TODO_LIST).(repository.TodoList)
	if err = r.Create(c.Request().Context(), &todoList); err != nil {
		return err
	}
	c.JSON(http.StatusCreated, todoList)
	return nil
}

func getAll(c echo.Context, ctn di.Container) error {
	r := ctn.Get(repository.TODO_LIST).(repository.TodoList)
	list, err := r.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	c.JSON(http.StatusCreated, list)
	return nil
}
