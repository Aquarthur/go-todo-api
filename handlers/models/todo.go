package models

// TodoDTO is the struct matching the JSON payload expected from and returned to users for the TODO API.
type TodoDTO struct {
	ID        string `json:"id,omitempty" binding:"-"`
	Task      string `json:"task" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
}
