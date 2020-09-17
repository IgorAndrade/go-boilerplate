package api

import "golang.org/x/sync/errgroup"

type Server interface {
	Start() error
	Stop() error
}

type List []Server

func (list List) StartAll(g *errgroup.Group) {
	for _, s := range list {
		g.Go(s.Start)
	}
}

func (list List) StopAll() {
	for _, s := range list {
		s.Stop()
	}
}
