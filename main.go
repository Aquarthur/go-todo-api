package main

import (
	"log"

	"github.com/Aquarthur/go-todo-api/api"
	"github.com/Aquarthur/go-todo-api/config"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("could not load configuration file")
	}
	api := api.NewTodoAPI(config)
	api.Start()
}
