package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, value interface{}) {
	data, err := json.Marshal(&value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed marshal list: %v", err)
		return
	}
	w.Write(data)
}

