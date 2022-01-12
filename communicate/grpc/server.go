package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"go-elio/communicate"
	"google.golang.org/grpc"
	"net"
	"time"
)

var (
	_ communicate.Server = (*server)(nil)
)

type ServerOption func(*server)

func Network(network string) ServerOption {
	return func(s *server) {
		s.network = network
	}
}

func Address(addr string) ServerOption {
	return func(s *server) {
		s.addr = addr
	}
}

func Timeout(time time.Duration) ServerOption {
	return func(s *server) {
		s.timeout = time
	}
}

func Listener(lister net.Listener) ServerOption {
	return func(s *server) {
		s.lister = lister
	}
}

func TlsConfig(tls *tls.Config) ServerOption {
	return func(s *server) {
		s.tls = tls
	}
}

type server struct {
	*grpc.Server              // grpc server
	tls          *tls.Config  // tls config
	lister       net.Listener // net lister
	network      string       // net model like "tcp","udp","http"
	addr         string       // grpc ip:port
	err          error
	timeout      time.Duration
}

func NewServer(options ...ServerOption) *server {
	srv := &server{
		network: "tcp",
		addr:    ":0",
		timeout: 1 * time.Second,
	}
	for _, o := range options {
		o(srv)
	}
	srv.err = srv.listener()
	return srv
}

func (s *server) listener() error {
	if s.lister == nil {
		lis, err := net.Listen(s.network, s.addr)
		if err != nil {
			return err
		}
		s.lister = lis
	}
	return nil
}

func (s *server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	fmt.Printf("[gRPC] server listening on: %v", s.lister.Addr().String())
	return s.Server.Serve(s.lister)
}

func (s *server) Stop(ctx context.Context) error {
	s.GracefulStop()
	fmt.Println("[gRPC] server stopping")
	return nil
}
