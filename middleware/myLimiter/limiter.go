package myLimiter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	stdlibMiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	redis2 "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
	"time"
	"txnbi-backend/conf"
)

type limiterLevel struct{}

// 限流级别
var (
	// LowLevel 每秒3次
	LowLevel = limiterLevel{}
	// MidLevel 每分钟30次
	MidLevel = limiterLevel{}
	// HighLevel 每小时30次
	HighLevel = limiterLevel{}
	// VeryHighLevel 每小时10次
	VeryHighLevel = limiterLevel{}
)

var limit *limiter.Limiter
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.LimiterRedisCfg.Host, conf.LimiterRedisCfg.Port),
		Password: conf.LimiterRedisCfg.Password,
		DB:       conf.LimiterRedisCfg.DB,
	})
	log.Println("Limit middleware initialized")
}

func New(name string, level limiterLevel) gin.HandlerFunc {
	rate := getRate(level)
	store, err := redis2.NewStore(rdb)
	if err != nil {
		panic(err)
	}
	limit = limiter.New(store, rate)

	mid := stdlibMiddleware.NewMiddleware(limit, stdlibMiddleware.WithKeyGetter(func(c *gin.Context) string {
		var userKey string
		userKey = c.GetString("userID")
		// 如果未登陆，就用IP标识
		if userKey == "" {
			userKey = c.ClientIP()
		}
		key := fmt.Sprintf("rate-limit:%s:%s", name, userKey)
		return key
	}),
		stdlibMiddleware.WithErrorHandler(func(c *gin.Context, err error) {
			c.JSON(200, gin.H{"statusCode": 1, "message": "您请求过快！"})
			c.Abort()
		}))

	return mid
}

func getRate(level limiterLevel) limiter.Rate {
	var rate limiter.Rate
	switch level {
	case LowLevel:
		rate = limiter.Rate{
			Period: time.Second,
			Limit:  3,
		}
	case MidLevel:
		rate = limiter.Rate{
			Period: time.Minute,
			Limit:  30,
		}
	case HighLevel:
		rate = limiter.Rate{
			Period: time.Hour,
			Limit:  30,
		}
	case VeryHighLevel:
		rate = limiter.Rate{
			Period: time.Hour,
			Limit:  10,
		}
	}
	return rate
}
