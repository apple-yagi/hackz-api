package model

import (
	"time"
)

type Task struct {
	ID        uint64 `gorm:"primaryKey"`
	Title     string `gorm:"unique"`
	Memo      *string
	Done      bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}
