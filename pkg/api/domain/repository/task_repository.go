package repository

import "hackz-api/pkg/api/domain/entity"

type TaskRepository interface {
	GetList() ([]entity.Task, error)
	FindByTaskID(taskID uint64) (entity.Task, error)
	Insert(title string, memo string) (entity.Task, error)
	Update(taskID uint64, title string, memo string) (entity.Task, error)
	Delete(taskID uint64) error
}
