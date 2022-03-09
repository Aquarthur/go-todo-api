package api

import (
	"fmt"
	"net/http"

	"github.com/Aquarthur/go-todo-api/config"
	"github.com/Aquarthur/go-todo-api/util"

	"github.com/Aquarthur/go-todo-api/handlers"
	"github.com/Aquarthur/go-todo-api/repository"
)

type TodoAPI struct {
	healthHandler *handlers.HealthHandler
	todoHandler   *handlers.TodoHandler
	config        *config.ServerConfig
}

func NewTodoAPI(config *config.Config) *TodoAPI {
	healthHandler := handlers.NewHealthHandler()
	todoRepository := repository.NewTodoRepository(config.Postgres)
	todoHandler := handlers.NewTodoHandler(todoRepository)
	return &TodoAPI{
		healthHandler: healthHandler,
		todoHandler:   todoHandler,
		config:        config.Server,
	}
}

func (api *TodoAPI) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = util.SplitPath(req.URL.Path)
	healthPath, _ := util.SplitPath(api.config.Health)
	switch head {
	case healthPath:
		api.healthHandler.ServeHTTP(w, req)
	case "todos":
		api.todoHandler.ServeHTTP(w, req)
	default:
		http.NotFound(w, req)
	}
}

func (api *TodoAPI) Start() {
	http.ListenAndServe(fmt.Sprintf(":%d", api.config.Port), api)
}
