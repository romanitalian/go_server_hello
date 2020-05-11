package router

import "net/http"

type rootHandler struct {
}


func (rh rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "tasks":
	case "task":
	default:
		http.NotFound(w, r)
	}
}
