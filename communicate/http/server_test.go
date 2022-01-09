package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"github.com/go-kratos/kratos/v2"
	"time"
)


func a() {
	a := kratos.New()
	fmt.Println(a)
}


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
