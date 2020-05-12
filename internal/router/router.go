package router

import (
	"net/http"

	"github.com/romanitalian/go_server_hello/internal/storages"
)

type Router struct {
	rootHandler rootHandler
}

func New(st storages.Store) *Router {
	return &Router{
		rootHandler: newRootHandler(st),
	}
}

func (r *Router) RootHandler() http.Handler {
	return r.rootHandler
}
