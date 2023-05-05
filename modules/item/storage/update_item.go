package storage

import (
	"TodoApi/common"
	"TodoApi/modules/item/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}

		return common.ErrDB(err)
	}

	return nil
}
