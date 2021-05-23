package controller

import (
	"hackz-api/pkg/api/interface/request"
	"hackz-api/pkg/api/interface/response"
	"hackz-api/pkg/api/usecase"
	"strconv"
)

type TaskController interface {
	Index(c request.Context)
	Show(c request.Context)
	Create(c request.Context)
	Put(c request.Context)
	Delete(c request.Context)
}

type TaskControllerImpl struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(TaskInteractor usecase.TaskInteractor) TaskController {
	return &TaskControllerImpl{
		Interactor: TaskInteractor,
	}
}

func (tc *TaskControllerImpl) Index(c request.Context) {
	tasks, err := tc.Interactor.GetList()
	if err != nil {
		c.JSON(500, response.NewError(500, err.Error()))
		return
	}

	c.JSON(200, response.NewSuccess(200, response.DataOpts(tasks)))
}

func (tc *TaskControllerImpl) Show(c request.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, response.NewError(400, err.Error()))
		return
	}

	task, err := tc.Interactor.FindByTaskID(id)
	if err != nil {
		c.JSON(500, response.NewError(500, err.Error()))
		return
	}
	c.JSON(200, response.NewSuccess(200, response.DataOpts(task)))
}

func (tc *TaskControllerImpl) Create(c request.Context) {
	type (
		Request struct {
			Title string `json:"title"`
			Memo  string `json:"memo"`
		}
	)
	req := Request{}
	if err := c.Bind(&req); err != nil {
		c.JSON(400, response.NewError(400, err.Error()))
		return
	}

	task, err := tc.Interactor.Create(req.Title, req.Memo)
	if err != nil {
		c.JSON(500, response.NewError(500, err.Error()))
		return
	}
	c.JSON(201, response.NewSuccess(201, response.DataOpts(task)))
}

func (tc *TaskControllerImpl) Put(c request.Context) {
	type (
		Request struct {
			Title string `json:"title"`
			Memo  string `json:"memo"`
		}
	)
	req := Request{}
	if err := c.Bind(&req); err != nil {
		c.JSON(400, response.NewError(400, err.Error()))
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, response.NewError(400, err.Error()))
		return
	}

	task, err := tc.Interactor.Update(id, req.Title, req.Memo)
	if err != nil {
		c.JSON(500, response.NewError(500, err.Error()))
		return
	}
	c.JSON(200, response.NewSuccess(200, response.DataOpts(task)))
}

func (tc *TaskControllerImpl) Delete(c request.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, response.NewError(400, err.Error()))
		return
	}

	err = tc.Interactor.Delete(id)
	if err != nil {
		c.JSON(500, response.NewError(500, err.Error()))
		return
	}
	c.JSON(204, response.NewSuccess(204))
}
