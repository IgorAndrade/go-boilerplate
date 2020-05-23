package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := h(c); err != nil {
			c.Logger().Error(err)
			c.JSON(http.StatusInternalServerError, err)
		}
		return nil
	}
}
