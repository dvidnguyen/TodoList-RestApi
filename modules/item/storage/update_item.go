package storage

import (
	"context"
	"social-todo-list/modules/item/entity"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, data *entity.ToDoItemUpdate) error {
	if err := s.db.Where(cond).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
