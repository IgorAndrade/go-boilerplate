package middleware

import "github.com/labstack/echo/v4"

func ApplyMiddleware(e *echo.Echo) {
	e.Use(HTTPErrorHandler)
}
