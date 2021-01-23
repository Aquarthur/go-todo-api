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

	dtos := make([]*models.TodoDTO, 0, len(todos))
	for _, todo := range todos {
		dtos = append(dtos, todo.ToDTO())
	}

	responseBody, err := json.Marshal(dtos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
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

	responseBody, err := json.Marshal(todo.ToDTO())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}
