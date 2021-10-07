package boot

import (
	_ "go-zone/config"
	"go-zone/src/dao"
)

func init() {
	dao.DynamoInit()
}
