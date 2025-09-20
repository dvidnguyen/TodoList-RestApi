package usecase

import (
	"context"
	"errors"
	"social-todo-list/modules/item/entity"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *entity.ToDoItemUpdate) error
}

type UpdateItemUC struct {
	store UpdateItemStorage
}

func NewUpdateItemUC(store UpdateItemStorage) *UpdateItemUC {
	return &UpdateItemUC{store: store}
}
func (uc *UpdateItemUC) UpdateNewItem(ctx context.Context, id int, data *entity.ToDoItemUpdate) error {
	data, err := uc.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data != nil && *data.Status == entity.ItemStatusDeleted {
		return entity.ErrItemDeleted
	}
	if err := uc.store.UpdateItem(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
