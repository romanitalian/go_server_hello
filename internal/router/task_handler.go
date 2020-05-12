package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/romanitalian/go_server_hello/internal/models"
	"github.com/romanitalian/go_server_hello/internal/storages"
)

type taskHandler struct {
	store storages.Store
}


func newTaskHandler(st storages.Store) taskHandler {
	return taskHandler{
		store: st,
	}
}

func (h taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createTask(w, r)
	case http.MethodDelete:
		h.deleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h taskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	jsDecoder := json.NewDecoder(r.Body)
	err := jsDecoder.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to decode body: %v", err)
		return
	}

	respTask, err := h.store.CreateTask(r.Context(), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to create Task in Store: %v", err)
		return
	}

	writeJSON(w, &respTask)
}

func (h taskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := shiftPath(r.URL.Path)
	log.Println(id)
	if id == "" {
		http.NotFound(w, r)
		return
	}

	err := h.store.DeleteTask(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to Delet task from Store: %v", err)
		return
	}
}