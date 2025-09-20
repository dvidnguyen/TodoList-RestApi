package usecase

import (
	"context"
	"social-todo-list/modules/item/entity"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error)
}

type getItemUC struct {
	store GetItemStorage
}

func NewGetItemUC(store GetItemStorage) *getItemUC {
	return &getItemUC{store: store}
}

func (s *getItemUC) GetItemByID(ctx context.Context, id int) (*entity.TodoItem, error) {
	data, err := s.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
