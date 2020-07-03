package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	rest "github.com/IgorAndrade/go-boilerplate/app/api/rest/webserver"
	"github.com/IgorAndrade/go-boilerplate/app/config"
	"github.com/IgorAndrade/go-boilerplate/app/db/mongo"
	"github.com/IgorAndrade/go-boilerplate/internal/service"
	"golang.org/x/sync/errgroup"
)

func main() {
	b := config.NewBuilder()
	config.Define(b)
	mongo.Define(b)
	service.Define(b)
	config.Build(b)
	defer config.Container.Delete()

	ctx, done := context.WithCancel(context.Background())
	defer done()
	g, gctx := errgroup.WithContext(ctx)
	s := rest.NewServer(gctx, done)
	g.Go(s.Start)

	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
		defer s.Stop()

		select {
		case sig := <-signalChannel:
			fmt.Printf("Received signal: %s\n", sig)
		case <-gctx.Done():
			fmt.Printf("closing signal goroutine\n")
			return gctx.Err()
		}

		return nil
	})

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
