package http

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/gorilla/mux"
	"go-elio/communicate"
	"go-elio/internal/endpoint"
	"go-elio/logger"
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	_ communicate.Server = (*server)(nil)
	_ communicate.EndPointer = (*server)(nil)
)

type ServerOption func(*server)

// Address set Addr
func Address(addr string) ServerOption {
	return func(s *server) {
		s.address = addr
	}
}

// Network set network
func Network(network string) ServerOption {
	return func(s *server) {
		s.network = network
	}
}

// Lister set lister
func Lister(lister net.Listener) ServerOption {
	return func(s *server) {
		s.lister = lister
	}
}

// Timeout set timeout
func Timeout(time time.Duration) ServerOption {
	return func(s *server) {
		s.timeout = time
	}
}

func Logger(h *logger.Helper) ServerOption {
	return func(s *server) {
		s.logger = h
	}
}

// http server struct interface by communicate.Server
type server struct {
	*http.Server
	tlsConfig *tls.Config
	lister    net.Listener
	timeout   time.Duration
	router    *mux.Router
	endpoint  *url.URL
	logger    *logger.Helper
	network   string
	address   string
	err       error
}

// NewServer : register http server
func NewServer(options ...ServerOption) *server {
	srv := &server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
		router:  mux.NewRouter(),
		logger: logger.NewHelper(logger.DefaultLogger),
	}
	for _, o := range options {
		o(srv)
	}
	srv.Server = &http.Server{
		Handler:   srv.router,
		TLSConfig: srv.tlsConfig,
	}
	return srv
}

// create a new route
func (s *server) Route(prefix string) *Router {
	return newRouter(prefix, s)
}

// HandleFunc register handle
func (s *server) HandleFunc(p string, f http.HandlerFunc) {
	s.router.HandleFunc(p, f)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) EndPoint() (*url.URL,error){
	if s.err != nil {
		return nil,s.err
	}
	return s.endpoint,nil
}

// Start start http server
func (s *server) Start(ctx context.Context) error {
	if err := s.listerAndEndpoint(); err != nil {
		return err
	}
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	s.logger.Infof("[HTTP] server is starting on addr %s", s.lister.Addr().String())
	var err error
	if s.tlsConfig != nil {
		err = s.ServeTLS(s.lister, "", "")
	} else {
		err = s.Serve(s.lister)
	}
	// when you Stop this http Server, this function will return server close like panic
	// so use errors.Is to clean up this panic
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *server) listerAndEndpoint() error {
	if s.lister == nil {
		var err error
		s.lister, err = net.Listen(s.network, s.address)
		if err != nil {
			return err
		}
	}
	s.endpoint = endpoint.NewEndpoint("http",s.lister.Addr().String(),s.tlsConfig!=nil)
	return nil
}

// Stop the server
func (s *server) Stop(ctx context.Context) error {
	s.logger.Info("[HTTP] server is stop")
	return s.Shutdown(ctx)
}
