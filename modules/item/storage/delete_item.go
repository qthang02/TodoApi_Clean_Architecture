package storage

import (
	"TodoApi/common"
	"TodoApi/modules/item/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	deleteStatus := model.ItemStatusDeleted

	if err := s.db.
		Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": deleteStatus.String()}).
		Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}

		return common.ErrDB(err)
	}

	return nil
}
