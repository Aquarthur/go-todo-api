package domain

import (
	"github.com/Aquarthur/go-todo-api/handlers/models"
	uuid "github.com/satori/go.uuid"
)

// Todo is the internal representation of a TODO item
type Todo struct {
	ID        string
	Task      string
	Completed bool
}

// ToDTO converts a Todo struct to a TodoDTO struct
func (t *Todo) ToDTO() *models.TodoDTO {
	return &models.TodoDTO{
		ID:        t.ID,
		Task:      t.Task,
		Completed: t.Completed,
	}
}

// FromDTO consumes a TodoDTO's data and uses it to populate the Todo
func (t *Todo) FromDTO(dto *models.TodoDTO) {
	t.ID = uuid.NewV4().String()
	t.Task = dto.Task
	t.Completed = dto.Completed
}
