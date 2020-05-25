package main

import (
	"fmt"

	"github.com/IgorAndrade/go-boilerplate/internal/service"

	"github.com/IgorAndrade/go-boilerplate/internal/repository/mongo"

	rest "github.com/IgorAndrade/go-boilerplate/app/api/rest/webserver"
	"github.com/IgorAndrade/go-boilerplate/app/config"
)

func main() {
	mongo.Init()
	service.Init()
	config.Build()
	defer config.Container.Delete()

	s := rest.NewServer()
	s.Start()
	fmt.Println("fim")
}
