package main

import (
	"log"

	todo "github.com/ohpasha/rest-api"
	handler "github.com/ohpasha/rest-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if error := srv.Run("8000", handlers.InitRouters()); error != nil {
		log.Fatalf("error: %s", error.Error())
	}
}
