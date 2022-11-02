package main

import (
	"log"

	todo "github.com/ohpasha/rest-api"
	"github.com/ohpasha/rest-api/pkg/handler"
	"github.com/ohpasha/rest-api/pkg/repository"
	"github.com/ohpasha/rest-api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todo.Server)

	if error := srv.Run("8000", handlers.InitRouters()); error != nil {
		log.Fatalf("error: %s", error.Error())
	}
}
