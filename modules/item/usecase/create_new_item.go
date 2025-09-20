package usecase

import (
	"context"
	"social-todo-list/modules/item/entity"
)

type InsertItemRepo interface {
	InsertItems(ctx context.Context, data *entity.ToDoItemCreatetion) error
}

type CreateItemUsecase struct {
	store InsertItemRepo
}

func NewCreateItemUsecase(store InsertItemRepo) *CreateItemUsecase {
	return &CreateItemUsecase{store: store}
}

func (uc *CreateItemUsecase) CreateNewItem(ctx context.Context, data *entity.ToDoItemCreatetion) error {
	title := data.Title
	if title == "" {
		return entity.ErrTitleBlank
	}
	if err := uc.store.InsertItems(ctx, data); err != nil {
		return err
	}
	return nil
}
