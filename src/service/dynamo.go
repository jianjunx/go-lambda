package service

import "gin-template/src/dao"

func PrintDynamoName() ([]string, error) {
	return dao.GetTableNames()
}
