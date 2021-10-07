package service

import (
	"go-zone/src/dao"
	"go-zone/src/model"
)

func PostItemDataAdd(row *model.PostAttributes) error {
	return dao.PostItemDataAdd(row)
}

func PostListGet() (*[]model.PostAttributes, error) {
	posts := []model.PostAttributes{}
	err := dao.PostDataScan(&posts)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
