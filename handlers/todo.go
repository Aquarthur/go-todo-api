package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Aquarthur/go-todo-api/domain"

	"github.com/Aquarthur/go-todo-api/handlers/models"

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

func (todoHandler *TodoHandler) FindAll(w http.ResponseWriter, req *http.Request) {
	todos, err := todoHandler.repository.FindAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todosBytes, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(todosBytes)
	w.WriteHeader(http.StatusOK)
}

func (todoHandler *TodoHandler) Create(w http.ResponseWriter, req *http.Request) {
	payload := new(models.TodoDTO)
	err := json.NewDecoder(req.Body).Decode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todo := new(domain.Todo)
	todo.FromDTO(payload)

	err = todoHandler.repository.SaveTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todoBytes, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(todoBytes)
	w.WriteHeader(http.StatusCreated)
}
