package usecase

import (
	"hackz-api/pkg/api/domain/entity"
	"hackz-api/pkg/api/domain/repository"
)

type TaskInteractor interface {
	GetList() ([]entity.Task, error)
	FindByTaskID(taskID uint64) (entity.Task, error)
	Create(title string, memo string) (entity.Task, error)
	Update(taskID uint64, title string, memo string) (entity.Task, error)
	Delete(taskID uint64) error
}

type TaskInteractorImpl struct {
	TaskRepository repository.TaskRepository
}

func NewTaskUsecase(TaskRepository repository.TaskRepository) TaskInteractor {
	return &TaskInteractorImpl{TaskRepository: TaskRepository}
}

func (ti *TaskInteractorImpl) GetList() ([]entity.Task, error) {
	return ti.TaskRepository.GetList()
}

func (ti *TaskInteractorImpl) FindByTaskID(taskID uint64) (entity.Task, error) {
	return ti.TaskRepository.FindByTaskID(taskID)
}

func (ti *TaskInteractorImpl) Create(title string, memo string) (entity.Task, error) {
	return ti.TaskRepository.Insert(title, memo)
}

func (ti *TaskInteractorImpl) Update(taskID uint64, title string, memo string) (entity.Task, error) {
	return ti.TaskRepository.Update(taskID, title, memo)
}

func (ti *TaskInteractorImpl) Delete(taskID uint64) error {
	return ti.TaskRepository.Delete(taskID)
}
