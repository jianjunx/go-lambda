package dao

import (
	"go-zone/src/model"
)

// 添加或修改book
func BookItemDataPut(row *model.PostBook) error {
	table := getBookTable()
	return table.Put(row).Run()
}

// 扫描book
func BookDataScan(books *[]model.PostBook) error {
	table := getBookTable()
	return table.Scan().All(books)
}

// 根据ID删除book
func BookItemDelete(id int) error {
	table := getBookTable()
	return table.Delete("id", id).Run()
}
