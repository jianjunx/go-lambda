package service

import (
	"go-zone/src/dao"
	"go-zone/src/model"
)

func PostItemDataAdd(row *model.PostAttributes) error {
	// Status 0 草稿 1 发布，为草稿状态时不必进行入库和更新操作
	if row.Status == 0 {
		return nil
	}
	// ActionType = delete执行删除操作
	if row.ActionType == "delete" {
		return PostItemDelete(row)
	}
	err := dao.PostItemDataPut(row)
	if err != nil {
		return err
	}
	return nil
}

// 删除文章
func PostItemDelete(row *model.PostAttributes) error {
	err := dao.PostItemDelete(row.Id)
	if err != nil {
		return err
	}
	// 如果当前书中仅剩一个文章则删除该书
	if row.Book.ItemsCount <= 1 {
		return dao.BookItemDelete(row.BookId)
	}
	return nil
}

func PostListGet(params *model.PostSearchParam) (*[]model.PostAttributes, int, error) {
	posts := []model.PostAttributes{}
	total, err := dao.PostDataScan(&posts, params)
	if err != nil {
		return nil, 0, err
	}
	return &posts, int(total), nil
}
