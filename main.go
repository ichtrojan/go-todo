package main

import (
	"github.com/ichtrojan/go-todo/routes"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":90", routes.Init())

	if err != nil {
		log.Fatal(err)
	}
}
