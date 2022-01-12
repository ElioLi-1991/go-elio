package grpc

import (
	"context"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()
	if err := s.Start(context.Background());err != nil {
		panic(err)
	}
	s.Stop(context.Background())
}
