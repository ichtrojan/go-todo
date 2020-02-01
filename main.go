package main

import (
	"github.com/ichtrojan/go-todo/routes"
	"github.com/ichtrojan/thoth"
	"log"
	"net/http"
)

func main() {
	logger, _ := thoth.Init("log")

	//if err := godotenv.Load(); err != nil {
	//	logger.Log(errors.New("no .env file found"))
	//	log.Fatal("No .env file found")
	//}

	err := http.ListenAndServe(":90", routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
