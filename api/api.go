package api

import (
	"net/http"

	"github.com/Aquarthur/go-todo-api/utils"

	"github.com/Aquarthur/go-todo-api/handlers"
	"github.com/Aquarthur/go-todo-api/repository"
)

type TodoAPI struct {
	todoHandler *handlers.TodoHandler
}

func NewTodoAPI() *TodoAPI {
	todoRepository := repository.NewTodoRepository()
	todoHandler := handlers.NewTodoHandler(todoRepository)
	return &TodoAPI{
		todoHandler: todoHandler,
	}
}

func (api *TodoAPI) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = utils.SplitPath(req.URL.Path)
	switch head {
	case "hello":
		hello := func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("Hello there!"))
		}
		http.HandlerFunc(hello).ServeHTTP(w, req)
	case "todos":
		api.todoHandler.ServeHTTP(w, req)
	default:
		http.NotFound(w, req)
	}
}

func (api *TodoAPI) Start() {
	http.ListenAndServe(":8080", api)
}
