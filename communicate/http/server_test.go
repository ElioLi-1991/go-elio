package http

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)


func TestNewServer(t *testing.T) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(r.RequestURI)
	}
	ctx := context.Background()
	s := NewServer()
	s.HandleFunc("/index", fn)
	s.HandleFunc("/index/{id:[0-9]+}", fn)
	go func() {
		if err := s.Start(ctx); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	s.Stop(ctx)
}
