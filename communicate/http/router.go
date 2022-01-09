package http

import (
	"net/http"
	"path"
)

type Router struct {
	prefix string
	srv    *server
}

// newRouter inherit by front router
// like this : prefix "/User"
// if add new router "/GetName" , finally will get "/User/GetName"
func newRouter(prefix string, srv *server) *Router {
	return &Router{
		prefix: prefix,
		srv:    srv,
	}
}

// Group create new router group
func (r *Router) Group(prefix string) *Router {
	return newRouter(path.Join(r.prefix, prefix), r.srv)
}

func (r *Router) Get(p string,h http.HandlerFunc) {
	r.Handle(http.MethodGet,p,h)
}

func (r *Router) Post(p string,h http.HandlerFunc) {
	r.Handle(http.MethodPost,p,h)
}

func (r *Router) Put(p string,h http.HandlerFunc) {
	r.Handle(http.MethodPut,p,h)
}

func (r *Router) Delete(p string,h http.HandlerFunc) {
	r.Handle(http.MethodDelete,p,h)
}

func (r *Router) Head(p string,h http.HandlerFunc) {
	r.Handle(http.MethodHead,p,h)
}

func (r *Router) Connect(p string,h http.HandlerFunc) {
	r.Handle(http.MethodConnect,p,h)
}

func (r *Router) Option(p string,h http.HandlerFunc) {
	r.Handle(http.MethodOptions,p,h)
}

func (r *Router) 

func (r *Router) Handle(method string,p string,f http.HandlerFunc) {
	r.srv.router.HandleFunc(path.Join(r.prefix,p),f).Methods(method)
}
