package middleware

import (
	"io"
	"net/http"

	"github.com/IgorAndrade/go-boilerplate/app/apiErrors"
	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := h(c); err != nil {
			c.Logger().Error(err)
			if e_http, ok := err.(*echo.HTTPError); ok {
				c.Response().WriteHeader(e_http.Code)
				c.Response().Write([]byte(e_http.Internal.Error()))
				return nil
			}
			if wt, ok := err.(io.WriterTo); ok {
				c.Response().WriteHeader(getHttpCode(err))
				wt.WriteTo(c.Response())
				return nil
			}
			c.JSON(http.StatusInternalServerError, err)
		}
		return nil
	}
}

func getHttpCode(err error) int {
	et := apiErrors.GetType(err)
	switch et {
	case apiErrors.BadRequest:
		return http.StatusBadRequest
	case apiErrors.NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
