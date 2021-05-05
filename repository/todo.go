package repository

import (
	"fmt"

	"github.com/Aquarthur/go-todo-api/config"
	"github.com/Aquarthur/go-todo-api/domain"
)

type TodoRepository interface {
	FindTodoByID(id string) (*domain.Todo, error)
	FindAllTodos() ([]*domain.Todo, error)
	SaveTodo(todo *domain.Todo) error
	DeleteTodo(id string) error
	UpdateTodo(id string, todo *domain.Todo) error
}

func NewTodoRepository(config *config.PostgresConfig) TodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string]*domain.Todo),
	}
}

type InMemoryTodoRepository struct {
	todos map[string]*domain.Todo
}

func (repo *InMemoryTodoRepository) FindTodoByID(id string) (*domain.Todo, error) {
	if todo, ok := repo.todos[id]; ok {
		return todo, nil
	}
	return nil, fmt.Errorf("could not fetch todo %s: %w", id, domain.ErrNotFound)
}

func (repo *InMemoryTodoRepository) FindAllTodos() ([]*domain.Todo, error) {
	todos := make([]*domain.Todo, 0, len(repo.todos))
	for _, val := range repo.todos {
		todos = append(todos, val)
	}
	return todos, nil
}

func (repo *InMemoryTodoRepository) SaveTodo(todo *domain.Todo) error {
	if _, ok := repo.todos[todo.ID]; ok {
		return fmt.Errorf("could not save todo %s: %w", todo.ID, domain.ErrConflict)
	}
	repo.todos[todo.ID] = todo
	return nil
}

func (repo *InMemoryTodoRepository) DeleteTodo(id string) error {
	if _, ok := repo.todos[id]; ok {
		delete(repo.todos, id)
		return nil
	}
	return fmt.Errorf("could not delete todo %s: %w", id, domain.ErrNotFound)
}

func (repo *InMemoryTodoRepository) UpdateTodo(id string, todo *domain.Todo) error {
	if _, ok := repo.todos[id]; ok {
		repo.todos[id] = &domain.Todo{
			ID:        id,
			Task:      todo.Task,
			Completed: todo.Completed,
		}
		return nil
	}
	return fmt.Errorf("could not update todo %s: %w", id, domain.ErrNotFound)
}
