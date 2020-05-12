package router

import (
	"net/http"

	"github.com/romanitalian/go_server_hello/internal/storages"
)

type rootHandler struct {
	tasksHandler tasksHandler
	taskHandler taskHandler
}

func newRootHandler(st storages.Store) rootHandler {
	return rootHandler {
		tasksHandler: newTasksHandler(st),
		taskHandler: newTaskHandler(st),
	}
}

func (h rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "tasks":
		h.tasksHandler.ServeHTTP(w, r)
	case "task":
		h.taskHandler.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
