package storage

import (
	"context"
	"social-todo-list/modules/item/entity"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	var data entity.TodoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	//return &entity.TodoItem{}
	return &data, nil

}
