package dao

import (
	"go-zone/src/model"
)

func PostItemDataAdd(row *model.PostAttributes) error {
	table := getPostTable()
	return table.Put(row).Run()
}

func PostDataScan(posts *[]model.PostAttributes) error {
	table := getPostTable()
	return table.Scan().All(posts)
}
