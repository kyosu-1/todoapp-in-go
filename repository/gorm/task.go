package gorm

import (
	"todoapp-in-go/domain/model"
	"todoapp-in-go/repository"

	"github.com/jinzhu/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository task repositoryのインスタンスを生成する
func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &taskRepository{db}
}

// FindAll taskを全件取得する
func (r *taskRepository) FindAll() ([]*model.Task, error) {
	var tasks []*model.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// Create taskをDBに保存する
func (r *taskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}


// FindByID taskをIDで検索する
func (r *taskRepository) FindByID(id int) (*model.Task, error) {
	var task model.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

// Update taskをDBに更新する
func (r *taskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := r.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Delete taskをDBから削除する
func (r *taskRepository) Delete(task *model.Task) error {
	if err := r.db.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
