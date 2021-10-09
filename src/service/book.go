package service

import (
	"go-zone/src/dao"
	"go-zone/src/model"
)

func BookListGet() (*[]model.PostBook, error) {
	books := []model.PostBook{}
	return &books, dao.BookDataScan(&books)
}
