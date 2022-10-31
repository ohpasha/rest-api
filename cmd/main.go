package main

import (
	"fmt"
	"log"

	todo "github.com/ohpasha/rest-api"
)

func main() {
	fmt.Print("main")

	srv := new(todo.Server)

	if error := srv.Run("8000"); error != nil {
		log.Fatalf("error: %s", error.Error())
	}
}
