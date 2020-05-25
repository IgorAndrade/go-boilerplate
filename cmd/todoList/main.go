package main

import (
	"fmt"

	"github.com/IgorAndrade/go-boilerplate/internal/service"

	rest "github.com/IgorAndrade/go-boilerplate/app/api/rest/webserver"
	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/IgorAndrade/go-boilerplate/app/db/mongo"
)

func main() {
	b := config.NewBuilder()
	config.Define(b)
	mongo.Define(b)
	service.Define(b)
	config.Build(b)
	defer config.Container.Delete()

	s := rest.NewServer()
	s.Start()
	fmt.Println("fim")
}
