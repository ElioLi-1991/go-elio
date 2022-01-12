package grpc

import (
	"context"
	"crypto/tls"
	"go-elio/communicate"
	"go-elio/internal/endpoint"
	"go-elio/logger"
	"google.golang.org/grpc"
	"net"
	"net/url"
	"time"
)

var (
	_ communicate.Server = (*server)(nil)
	_ communicate.EndPointer = (*server)(nil)
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
		s.tlsConfig = tls
	}
}

func Logger(h *logger.Helper) ServerOption {
	return func(s *server) {
		s.logger = h
	}
}

type server struct {
	*grpc.Server              // grpc server
	tlsConfig    *tls.Config  // tls config
	lister       net.Listener // net lister
	network      string       // net model like "tcp","udp","http"
	addr         string       // grpc ip:port
	err          error
	timeout      time.Duration
	endpoint     *url.URL
	logger       *logger.Helper

}

func NewServer(options ...ServerOption) *server {
	srv := &server{
		network: "tcp",
		addr: ":0",
		timeout: 1 * time.Second,
		logger: logger.NewHelper(logger.DefaultLogger),
	}
	for _,o := range options {
		o(srv)
	}
	srv.Server = grpc.NewServer()
	srv.err = srv.listenerAndEndpoint()
	return srv
}

func (s *server) listenerAndEndpoint() error {
	if s.lister == nil {
		lis, err := net.Listen(s.network, s.addr)
		if err != nil {
			return err
		}
		s.lister = lis
	}
	s.endpoint = endpoint.NewEndpoint("grpc",s.lister.Addr().String(),s.tlsConfig!=nil)
	return nil
}

func (s *server) EndPoint() (*url.URL,error) {
	if s.err != nil {
		return nil,s.err
	}
	return s.endpoint,nil
}

func (s *server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	s.logger.Infof("[gRPC] server listening on: %v",s.lister.Addr().String())
	return s.Server.Serve(s.lister)
}

func (s *server) Stop(ctx context.Context) error {
	s.GracefulStop()
	s.logger.Info("[gRPC] server is stop")
	return nil
}
