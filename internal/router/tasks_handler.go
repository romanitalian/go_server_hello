package router

import (
	"fmt"
	"net/http"

	"github.com/romanitalian/go_server_hello/internal/storages"
)

type tasksHandler struct {
	store storages.Store
}

func newTasksHandler(st storages.Store) tasksHandler {
	return tasksHandler{
		store: st,
	}
}

func (h tasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	list, err := h.store.GetTaskList(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to get task list from store: %v", err)
		return
	}

	writeJSON(w, &list)
}