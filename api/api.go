package api

import (
	"net/http"

	"github.com/Aquarthur/go-todo-api/handlers"
	"github.com/Aquarthur/go-todo-api/middleware"
	"github.com/Aquarthur/go-todo-api/repository"
)

type TodoAPI struct {
	mux *http.ServeMux
}

func NewTodoAPI() *TodoAPI {
	mux := http.NewServeMux()
	hello := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello there!"))
	}
	mux.Handle("/hello", middleware.Log(http.HandlerFunc(hello)))

	todoRepository := repository.NewTodoRepository()
	todoHandler := handlers.NewTodoHandler(todoRepository)

	mux.Handle("/todo", middleware.Log(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.FindAll(w, r)
		case http.MethodPost:
			todoHandler.Create(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	})))

	return &TodoAPI{
		mux: mux,
	}
}

func (api *TodoAPI) Start() {
	http.ListenAndServe(":8080", api.mux)
}
