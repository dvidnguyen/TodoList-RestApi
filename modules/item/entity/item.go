package entity

import (
	"errors"
	"social-todo-list/common"
)

var (
	ErrTitleBlank  = errors.New("Title cannot be blank")
	ErrItemDeleted = errors.New("Item is  deleted")
)

type TodoItem struct {
	common.SqlModel
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string { return "todo_items" }

type ToDoItemCreatetion struct {
	Id          int        `json:"-" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
}

func (ToDoItemCreatetion) TableName() string { return TodoItem{}.TableName() }

type ToDoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title"`
	Description *string     `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (ToDoItemUpdate) TableName() string { return TodoItem{}.TableName() }
