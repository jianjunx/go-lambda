package boot

import (
	_ "go-zone/config" // 初始化配置
	_ "go-zone/router" // 初始化路由
	"go-zone/src/dao"
)

func init() {
	// 初始化dynamodb
	dao.DynamoInit()
}
