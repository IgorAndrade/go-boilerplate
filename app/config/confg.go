package config

import (
	"fmt"
	"os"

	"github.com/sarulabs/di"
)

type Config struct {
	Rest  Rest
	Mongo Mongo
}

type Rest struct {
	Port string
}

type Mongo struct {
	Address  string
	User     string
	Password string
}

var CONFIG = "config"
var c *Config

func Define(b *di.Builder) {
	c = &Config{
		Rest: Rest{
			Port: fmt.Sprintf(":%s", os.Getenv("API_PORT")),
		},
		Mongo: Mongo{
			Address:  os.Getenv("MONGO_URL"),
			User:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
		},
	}
	b.Add(di.Def{
		Name:  CONFIG,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return c, nil
		},
	})
}
