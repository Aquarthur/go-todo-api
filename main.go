package main

import (
	"github.com/Aquarthur/go-todo-api/api"
)

func main() {
	api := api.NewTodoAPI()
	api.Start()
}
