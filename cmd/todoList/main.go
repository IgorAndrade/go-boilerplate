package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IgorAndrade/go-boilerplate/app/api"
	rest "github.com/IgorAndrade/go-boilerplate/app/api/rest/webserver"
	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/IgorAndrade/go-boilerplate/app/db/mongo"
	"github.com/IgorAndrade/go-boilerplate/internal/service"
	"github.com/sarulabs/di"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctn := initContainer()
	defer ctn.Delete()

	ctx, done := context.WithCancel(context.Background())
	defer done()
	g, gctx := errgroup.WithContext(ctx)

	s := rest.NewServer(gctx, done)
	serv := api.List{s}
	serv.StartAll(g)

	g.Go(waitSignalChannel(gctx, serv))

	err := g.Wait()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Print("context was canceled")
		} else {
			fmt.Printf("received error: %v", err)
		}
	} else {
		fmt.Println("finished clean")
	}
}

func waitSignalChannel(gctx context.Context, serv api.List) func() error {
	return func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, os.Kill, syscall.SIGSEGV)
		defer serv.StopAll()

		select {
		case sig := <-signalChannel:
			fmt.Printf("Received signal: %s\n", sig)
		case <-gctx.Done():
			fmt.Printf("closing signal goroutine\n")
			return gctx.Err()
		}

		return nil
	}
}

func initContainer() di.Container {
	b := config.NewBuilder(config.Define, mongo.Define, service.Define)
	return config.Build(b)
}
