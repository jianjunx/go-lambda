package service

import (
	"go-zone/src/dao"
	"go-zone/src/model"
)

func PostItemDataAdd(row *model.PostAttributes) error {
	return dao.PostItemDataAdd(row)
}
