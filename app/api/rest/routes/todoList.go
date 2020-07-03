package route

import (
	"net/http"

	"github.com/IgorAndrade/go-boilerplate/app/apiErrors"
	"github.com/IgorAndrade/go-boilerplate/internal/service"

	"github.com/IgorAndrade/go-boilerplate/internal/model"
	"github.com/labstack/echo/v4"
)

func create(c echo.Context, ctn GetterDI) error {
	todoList := model.TodoList{}
	err := c.Bind(&todoList)
	if err != nil {
		return apiErrors.BadRequest.NewError(err)
	}
	s := ctn.Get(service.TODO_LIST).(service.TodoList)
	if err = s.Create(c.Request().Context(), &todoList); err != nil {
		return err
	}
	c.JSON(http.StatusCreated, todoList)
	return nil
}

func getAll(c echo.Context, ctn GetterDI) error {
	s := ctn.Get(service.TODO_LIST).(service.TodoList)
	list, err := s.GetAll(c.Request().Context())
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, list)
	return nil
}
