package redis

import (
	"context"
	"fmt"
	"gin-template/settings"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init(conf *settings.Redis) (err error) {
	addr := fmt.Sprintf("%s:%v", conf.Host, conf.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Passwd,
		DB:       0,             // use default DB
		PoolSize: conf.Poolsize, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func Close() {
	rdb.Close()
}
