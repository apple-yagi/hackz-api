package entity

import (
	"errors"
	"time"
)

type Task struct {
	ID        uint64
	Title     string
	Memo      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(title string, memo string) (*Task, error) {
	if title == "" {
		return nil, errors.New("タイトルを入力してください")
	}

	task := &Task{
		Title: title,
		Memo:  memo,
	}

	return task, nil
}

func (t *Task) Set(title string, memo string) error {
	if title == "" {
		return errors.New("タイトルを入力してください")
	}

	t.Title = title
	t.Memo = memo

	return nil
}
