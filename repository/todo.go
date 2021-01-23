package repository

import (
	"fmt"

	"github.com/Aquarthur/go-todo-api/domain"
)

type TodoRepository interface {
	FindTodoByID(id string) (*domain.Todo, error)
	FindAllTodos() ([]*domain.Todo, error)
	SaveTodo(todo *domain.Todo) error
	DeleteTodo(id string) error
	UpdateTodo(todo *domain.Todo) error
}

func NewTodoRepository() TodoRepository {
	return &InMemoryTodoRepository{}
}

type InMemoryTodoRepository struct {
	todos map[string]*domain.Todo
}

func (repo *InMemoryTodoRepository) FindTodoByID(id string) (*domain.Todo, error) {
	if todo, ok := repo.todos[id]; ok {
		return todo, nil
	}
	return nil, fmt.Errorf("could not find todo with id %s", id)
}

func (repo *InMemoryTodoRepository) FindAllTodos() ([]*domain.Todo, error) {
	todos := make([]*domain.Todo, len(repo.todos))
	for _, val := range repo.todos {
		todos = append(todos, val)
	}
	return todos, nil
}

func (repo *InMemoryTodoRepository) SaveTodo(todo *domain.Todo) error {
	if _, ok := repo.todos[todo.ID]; ok {
		return fmt.Errorf("todo with id %s already exists", todo.ID)
	}
	repo.todos[todo.ID] = todo
	return nil
}

func (repo *InMemoryTodoRepository) DeleteTodo(id string) error {
	if _, ok := repo.todos[id]; ok {
		delete(repo.todos, id)
		return nil
	}
	return fmt.Errorf("could not find todo with id %s", id)
}

func (repo *InMemoryTodoRepository) UpdateTodo(todo *domain.Todo) error {
	repo.todos[todo.ID] = todo
	return nil
}
