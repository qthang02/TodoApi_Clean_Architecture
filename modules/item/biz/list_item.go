package biz

import (
	"TodoApi/common"
	"TodoApi/modules/item/model"
	"context"
)

type ListItemStorage interface {
	ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error)
}

type listItemBiz struct {
	storage ListItemStorage
}

func NewListItemBiz(storage ListItemStorage) *listItemBiz {
	return &listItemBiz{storage: storage}
}

func (biz *listItemBiz) GetListItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	data, err := biz.storage.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
