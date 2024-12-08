package myRedis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"txnbi-backend/conf"
)

var Cli *redis.Client

func init() {
	Cli = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisCfg.Host, conf.RedisCfg.Port),
		Password: conf.RedisCfg.Password, // no password set
		DB:       conf.RedisCfg.DB,       // use default DB
	})
}
