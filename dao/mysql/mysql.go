package mysql

import (
	"fmt"
	"gin-template/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var DB *sqlx.DB

func Init(conf *settings.Mysql) (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True",
		conf.User, conf.Passwd, conf.Host, conf.Port, conf.Dbname,
	)
	// 也可以使用MustConnect连接不成功就panic
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err", zap.Error(err))
		return
	}
	DB.SetMaxOpenConns(conf.MaxOpenConns)
	DB.SetMaxIdleConns(conf.MaxIdleConns)
	return
}
