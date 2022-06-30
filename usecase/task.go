package usecase

import (
	"todoapp-in-go/domain/model"
	"todoapp-in-go/repository"
)

// TaskUsecase task usecaseのinterface
type TaskUsecase interface {
	FindAll() ([]*model.Task, error)
	Create(title, content string) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(id int, title, content string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	repository repository.TaskRepository
}

// NewTaskUsecase task usecaseのインスタンスを生成する
func NewTaskUsecase(repository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repository}
}

// FindAll taskを全件取得する
func (u *taskUsecase) FindAll() ([]*model.Task, error) {
	tasks, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Create taskをDBに保存する
func (u *taskUsecase) Create(title, content string) (*model.Task, error) {
	task, err := model.NewTask(title, content)
	if err != nil {
		return nil, err
	}

	createdTask, err := u.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

// FindByID taskをIDで検索する
func (u *taskUsecase) FindByID(id int) (*model.Task, error) {
	task, err := u.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Update taskを更新
func (u *taskUsecase) Update(id int, title, content string) (*model.Task, error) {
	task, err := u.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = task.Set(title, content)
	if err != nil {
		return nil, err
	}

	updatedTask, err := u.repository.Update(task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

// Delete taskをDBから削除する
func (u *taskUsecase) Delete(id int) error {
	task, err := u.repository.FindByID(id)
	if err != nil {
		return err
	}

	err = u.repository.Delete(task)
	if err != nil {
		return err
	}

	return nil
}
