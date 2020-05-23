package rest

import (
	"github.com/IgorAndrade/go-boilerplate/app/api"
	route "github.com/IgorAndrade/go-boilerplate/app/api/rest/routes"
	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/labstack/echo/v4"
)

//Server struct
type server struct {
	server *echo.Echo
}

//NewServer Create a new Rest server
func NewServer() api.Server {
	e := echo.New()
	route.ApplyRoutes(e)
	return &server{
		server: e,
	}
}

//Start a rest server
func (s server) Start() error {
	c := config.Container.Get(config.CONFIG).(*config.Config)
	return s.server.Start(c.Rest.Port)
}

//Stop a rest server
func (s server) Stop() error {
	return s.Stop()
}
