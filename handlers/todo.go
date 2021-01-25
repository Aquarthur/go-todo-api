package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Aquarthur/go-todo-api/domain"
	"github.com/Aquarthur/go-todo-api/utils"

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

func (handler *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var todoID string
	todoID, r.URL.Path = utils.SplitPath(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if todoID != "" {
		switch r.Method {
		case http.MethodGet:
			handler.Find(todoID, w)
		case http.MethodPut:
			payload := new(models.TodoDTO)
			err := json.NewDecoder(r.Body).Decode(payload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			handler.Update(todoID, w, payload)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			handler.FindAll(w)
		case http.MethodPost:
			payload := new(models.TodoDTO)
			err := json.NewDecoder(r.Body).Decode(payload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			handler.Create(w, payload)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (todoHandler *TodoHandler) Find(id string, w http.ResponseWriter) {
	todo, err := todoHandler.repository.FindTodoByID(id)
	if err != nil {
		handleError(err, w)
		return
	}

	responseBody, err := json.Marshal(todo.ToDTO())
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (todoHandler *TodoHandler) FindAll(w http.ResponseWriter) {
	todos, err := todoHandler.repository.FindAllTodos()
	if err != nil {
		handleError(err, w)
		return
	}

	dtos := make([]*models.TodoDTO, 0, len(todos))
	for _, todo := range todos {
		dtos = append(dtos, todo.ToDTO())
	}

	responseBody, err := json.Marshal(dtos)
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (todoHandler *TodoHandler) Create(w http.ResponseWriter, payload *models.TodoDTO) {
	todo := new(domain.Todo)
	todo.FromDTO(payload)

	err := todoHandler.repository.SaveTodo(todo)
	if err != nil {
		handleError(err, w)
		return
	}

	responseBody, err := json.Marshal(todo.ToDTO())
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

func (todoHandler *TodoHandler) Update(id string, w http.ResponseWriter, payload *models.TodoDTO) {
	todo := new(domain.Todo)
	todo.FromDTO(payload)

	err := todoHandler.repository.UpdateTodo(id, todo)
	if err != nil {
		handleError(err, w)
		return
	}

	responseBody, err := json.Marshal(todo.ToDTO())
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
