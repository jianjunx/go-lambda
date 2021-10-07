package dao

import (
	"go-zone/config"
	"go-zone/src/model"
)

func PostItemDataAdd(row *model.PostAttributes) error {
	table := getTable(config.GetConfig().PostTableName)
	return table.Put(row).Run()
}
