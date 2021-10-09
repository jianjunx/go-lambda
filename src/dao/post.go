package dao

import (
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

func PostItemIndexGet(posts *[]model.PostAttributes, slug string) error {
	table := getPostTable()
	return table.Scan().Filter("'slug' = ?", slug).All(posts)
}

// 删除post
func PostItemDelete(id int) error {
	table := getPostTable()
	return table.Delete("id", id).Run()
}
