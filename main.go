package main

import (
	"log"
	"net/http"

	router "github.com/romanitalian/go_server_hello/internal"
)

func main() {
	rt := router.New()
	if err := http.ListenAndServe(":8080", rt.RootHandler()); err != nil {
		log.Fatal("Error on start server")
	}
}
