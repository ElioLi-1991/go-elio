package http

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestRouter(t *testing.T) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(r.RequestURI)
	}
	srv := NewServer()
	r := srv.Route("/index")
	rGroup := r.Group("/User")
	rGroup.Get("/GetUser", fn)
	rGroup.Get("/GetList", fn)
	ctx := context.Background()
	go func() {
		if err := srv.Start(ctx); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	srv.Stop(ctx)
}
