package biz

import (
	"TodoApi/modules/item/model"
	"context"
	"strings"
)

type InsertItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type insertItemBiz struct {
	store InsertItemStorage
}

func NewInsertItemBiz(store InsertItemStorage) *insertItemBiz {
	return &insertItemBiz{store: store}
}

func (biz *insertItemBiz) InsertNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
