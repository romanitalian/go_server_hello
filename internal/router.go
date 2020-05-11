package router

import "net/http"

type Router struct {
	rootHandler rootHandler
}

func New() *Router {
	return &Router{}
}

func (r *Router) RootHandler() http.Handler {
	return r.rootHandler
}
