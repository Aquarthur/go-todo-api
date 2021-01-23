package handlers

import (
	"github.com/Aquarthur/go-todo-api/repository"
)

type TodoHandler struct {
	repository repository.TodoRepository
}

func NewTodoHandler(todoRepository repository.TodoRepository) *TodoHandler {
	return &TodoHandler{
		repository: todoRepository,
	}
}
