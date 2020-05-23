package route

import (
	"net/http"

	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type handlerDiReq func(c echo.Context, ctn di.Container) error

func ApplyRoutes(e *echo.Echo) {
	e.POST("/todo-list", injectDiReq(create))
	e.GET("/todo-list", injectDiReq(getAll))
}

func injectDiReq(h handlerDiReq) echo.HandlerFunc {
	return func(e echo.Context) error {
		rDI, err := config.Container.SubContainer()
		if err != nil {
			e.JSON(http.StatusInternalServerError, err)
			return err
		}
		return h(e, rDI)
	}
}
