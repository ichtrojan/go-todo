package main

import (
	"errors"
	"github.com/ichtrojan/go-todo/routes"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
