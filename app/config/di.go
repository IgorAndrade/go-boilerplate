package config

import (
	"log"

	"github.com/sarulabs/di"
)

var b *di.Builder = newBuilder()

//Container used to DI
var Container di.Container

func newBuilder() *di.Builder {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
	}

	return builder
}

//AddDef add dependency into container
func AddDef(defs ...di.Def) error {
	return b.Add(defs...)
}

//Build container
func Build() {
	Container = b.Build()
}
