package grpc

import (
	"context"
	"crypto/tls"
	"go-elio/communicate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"net"
	"time"
)

var (
	_ communicate.Server = (*server)(nil)
)

type ServerOption func(o *server)

type server struct {
	*grpc.Server
	tlsConfig *tls.Config
	lister    net.Listener
	err 	error
	network string
	addr 	string
	timeout time.Duration
	grpcOpts []grpc.ServerOption
	health *health.Server
}

func NewServer(options ...ServerOption) *server {
	srv := &server{
		network: "tcp",
		addr: ":0",
		timeout: 1 * time.Second,
		health: health.NewServer(),
	}
	for _,o := range options {
		o(srv)
	}
	return srv
}

func (s *server) Start(ctx context.Context) error {
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return nil
}
