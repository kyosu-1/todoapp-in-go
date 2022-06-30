package repository

import (
	"todoapp-in-go/domain/model"
)

// TaskRepository task repositoryのinterface
type TaskRepository interface {
	FindAll() ([]*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(task *model.Task) error
}
