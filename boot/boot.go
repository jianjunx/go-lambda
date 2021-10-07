package boot

import "gin-template/src/dao"

func init() {
	dao.DynamoInit()
}
