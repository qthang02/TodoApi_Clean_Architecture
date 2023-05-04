package storage

import (
	"TodoApi/modules/item/model"
	"context"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.TodoItem{}.TableName()).
		Where(cond).Updates(map[string]interface{}{"status": "Deleted"}).Error; err != nil {
		return err
	}

	return nil
}
