package gorm

import (
	"hackz-api/pkg/api/domain/entity"
	"hackz-api/pkg/api/domain/repository"
	"hackz-api/pkg/api/infrastructure/persistence/gorm/model"

	"github.com/jinzhu/gorm"
)

type TaskRepositoryImpl struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepositoryImpl{Conn: conn}
}

func (tr *TaskRepositoryImpl) GetList() (e []entity.Task, err error) {
	tasks := []model.Task{}
	if err = tr.Conn.Find(&tasks).Error; err != nil {
		return
	}

	n := len(tasks)
	e = make([]entity.Task, n)
	for i := 0; i < n; i++ {
		e[i].ID = tasks[i].ID
		e[i].Title = tasks[i].Title
		e[i].Memo = *tasks[i].Memo
		e[i].CreatedAt = tasks[i].CreatedAt
		e[i].UpdatedAt = tasks[i].UpdatedAt
	}
	return
}

func (tr *TaskRepositoryImpl) FindByTaskID(taskID uint64) (e entity.Task, err error) {
	task := model.Task{}
	if err = tr.Conn.First(&task, taskID).Error; err != nil {
		return
	}

	e = entity.Task{
		ID:        task.ID,
		Title:     task.Title,
		Memo:      *task.Memo,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return
}

func (tr *TaskRepositoryImpl) Insert(title string, memo string) (e entity.Task, err error) {
	task := &model.Task{
		Title: title,
		Memo:  &memo,
	}

	if err = tr.Conn.Create(task).Error; err != nil {
		return
	}

	e = entity.Task{
		ID:        task.ID,
		Title:     task.Title,
		Memo:      *task.Memo,
		UpdatedAt: task.UpdatedAt,
	}

	return
}

func (tr *TaskRepositoryImpl) Update(taskID uint64, title string, memo string) (e entity.Task, err error) {
	task := &model.Task{
		ID:    taskID,
		Title: title,
		Memo:  &memo,
	}

	if err = tr.Conn.Model(task).Update(task).Error; err != nil {
		return
	}

	e = entity.Task{
		ID:        task.ID,
		Title:     task.Title,
		Memo:      *task.Memo,
		UpdatedAt: task.UpdatedAt,
	}

	return
}

func (tr *TaskRepositoryImpl) Delete(taskID uint64) (err error) {
	task := &model.Task{
		ID: taskID,
	}

	if err = tr.Conn.Delete(task).Error; err != nil {
		return
	}

	return
}
