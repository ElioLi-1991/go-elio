package http

import (
	"context"
	"crypto/tls"
	"errors"
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

// http server struct interface by communicate.Server
type server struct {
	*http.Server
	tlsConfig *tls.Config
	lister  net.Listener
	timeout time.Duration
	router  *http.ServeMux
	network string
	address string
	err     error
}

// NewServer : register http server
func NewServer(options ...ServerOption) *server {
	srv := &server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
		router: http.NewServeMux(),
	}
	for _, o := range options {
		o(srv)
	}

	srv.Server = &http.Server {
		Handler: srv.router,
		TLSConfig: srv.tlsConfig,
	}

	return srv
}

// HandleFunc register handle
func (s *server) HandleFunc (path string, f func(http.ResponseWriter,*http.Request)) {
	s.router.HandleFunc(path,f)
}

// HandleFuncComplex register handle complex
func (s *server) HandleFuncComplex(hc HandleFuncComplex) {
	for _,h := range hc {
		s.router.HandleFunc(h.Path,h.Func)
	}
}

// Start start http server
func (s *server) Start(ctx context.Context) error {
	if err := s.registerLister();err != nil {
		return err
	}
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	fmt.Printf("[HTTP] server is starting on addr %s", s.lister.Addr().String())
	var err error
	if s.tlsConfig != nil {
		err = s.ServeTLS(s.lister,"","")
	} else {
		err = s.Serve(s.lister)
	}
	// when you Stop this http Server, this function will return server close like panic
	// so use errors.Is to clean up this panic
	if !errors.Is(err,http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *server) registerLister() error {
	if s.lister == nil {
		var err error
		s.lister , err = net.Listen(s.network,s.address)
		if err != nil {
			return err
		}
	}
	return nil
}

// Stop the server
func (s *server) Stop(ctx context.Context) error {
	fmt.Println("[HTTP] server is stop")
	return s.Shutdown(ctx)
}


