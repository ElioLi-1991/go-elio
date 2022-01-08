package http

import "path"

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

func (r *Router) Handle() {

}
