package config

import (
	"log"

	"github.com/sarulabs/di"
)

//Container used to DI
var Container di.Container

func NewBuilder() *di.Builder {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
	}

	return builder
}

//Build container
func Build(b *di.Builder) {
	Container = b.Build()
}
