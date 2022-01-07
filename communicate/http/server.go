package http

import (
	"context"
	"fmt"
	"go-elio/communicate"
	"net"
	"net/http"
	"time"
)

var (
	_ communicate.Server = (*server)(nil)
)

type ServerOption func(*server)

// Address config Addr
func Address(addr string) ServerOption {
	return func(s *server) {
		s.address = addr
	}
}

// Network config network
func Network(network string) ServerOption {
	return func(s *server) {
		s.network = network
	}
}

// Lister config lister
func Lister(lister net.Listener) ServerOption {
	return func(s *server) {
		s.lister = lister
	}
}

// Timeout config timeout
func Timeout(time time.Duration) ServerOption {
	return func(s *server) {
		s.timeout = time
	}
}

// a http server struct interface by communicate.Server
type server struct {
	*http.Server
	lister  net.Listener
	timeout time.Duration
	network string
	address string
	err     error
}

// NewServer : register http server
func NewServer(options ...ServerOption) communicate.Server {
	s := &server{
		network: "tcp",
		address: ":8000",
		timeout: 1 * time.Second,
	}
	for _, o := range options {
		o(s)
	}
	return s
}

// Start start http server
func (s *server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	fmt.Printf("[HTTP] server is starting on addr : %s", s.lister.Addr().String())
	return s.Serve(s.lister)
}

// Stop the server
func (s *server) Stop(ctx context.Context) error {
	fmt.Println("[HTTP] server is stop")
	return s.Shutdown(ctx)
}
