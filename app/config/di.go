package config

import (
	"log"

	"github.com/sarulabs/di"
)

//Container used to DI
var Container di.Container

func NewBuilder(opts ...func(*di.Builder)) *di.Builder {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range opts {
		f(builder)
	}

	return builder
}

//Build container
func Build(b *di.Builder) di.Container {
	Container = b.Build()
	return Container
}
