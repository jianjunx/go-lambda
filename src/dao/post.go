package dao

import (
	"fmt"
	"go-zone/src/model"

	"github.com/guregu/dynamo"
)

// 添加post
func PostItemDataPut(row *model.PostAttributes) error {
	table := getPostTable()
	return table.Put(row).Run()
}

// 扫描post
func PostDataScan(posts *[]model.PostAttributes, params *model.PostSearchParam) (int64, error) {
	table := getPostTable()
	scan := table.Scan()
	if params.Book != 0 {
		scan = table.Scan().Filter("'book_id' = ?", params.Book)
	}
	total, err := scan.Count()
	if err != nil {
		return 0, err
	}
	if params.Page == 1 {
		scan.Limit(int64(params.Size)).All(posts)
		return total, nil
	}
	var lastKey dynamo.PagingKey
	lm := int64((params.Page - 1) * params.Size)
	lastKey, err = scan.Limit(lm).AllWithLastEvaluatedKey(&[]model.PostAttributes{})
	if err != nil {
		return 0, err
	}
	scan.StartFrom(lastKey).Limit(int64(params.Size)).All(posts)
	return total, nil
}

// 删除post
func PostItemDelete(id int) error {
	table := getPostTable()
	return table.Delete("id", id).Run()
}

// 添加或修改book
func BookItemDataPut(row *model.PostBook) error {
	table := getBookTable()
	return table.Put(row).Run()
}

// 扫描book
func BookDataScan(books *[]model.PostBook) error {
	table := getBookTable()
	ite := table.Scan().SearchLimit(10).Iter()
	fmt.Println("ite", ite)
	return nil
}

// 根据ID删除book
func BookItemDelete(id int) error {
	table := getBookTable()
	return table.Delete("id", id).Run()
}
