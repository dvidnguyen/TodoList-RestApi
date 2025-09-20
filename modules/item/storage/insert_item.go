package storage

import (
	"context"
	"social-todo-list/modules/item/entity"
)

func (s *sqlStore) InsertItems(ctx context.Context, data *entity.ToDoItemCreatetion) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
