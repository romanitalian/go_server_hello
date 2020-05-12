package main

import (
	"log"
	"net/http"

	"github.com/romanitalian/go_server_hello/internal/router"
	"github.com/romanitalian/go_server_hello/internal/storages/memstore"
)

func main() {
	rt := router.New(memstore.New())
	if err := http.ListenAndServe(":8080", rt.RootHandler()); err != nil {
		log.Fatal("Error on start server")
	}
}
