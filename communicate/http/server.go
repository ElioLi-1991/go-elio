package http

import (
	"context"
	"go-elio/communicate"
	"net"
	"net/http"
)


var (
	_ communicate.Server = (*server)(nil)
)

type ServerOption func(*server)

type server struct {
	*http.Server
	lister net.Listener
	network string
	addr 	string
	err 	error
}

func Address(addr string) ServerOption {
	return func(s *server) {
		s.Addr = addr
	}
}


func (s *server) Start(ctx context.Context) error {
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return nil
}