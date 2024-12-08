package middleware

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

var limit *limiter.Limiter

func init() {
	rate := limiter.Rate{
		Period: time.Hour,
		Limit:  20,
	}

	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       3,  // use default DB
	//})

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.LimiterRedisCfg.Host, conf.LimiterRedisCfg.Port),
		Password: conf.LimiterRedisCfg.Password,
		DB:       conf.LimiterRedisCfg.DB,
	})

	fmt.Println(rdb)

	store, err := redis2.NewStore(rdb)
	if err != nil {
		panic(err)
	}
	limit = limiter.New(store, rate)
	log.Println("Limit middleware initialized")
}

func Limiter() gin.HandlerFunc {
	mid := stdlibMiddleware.NewMiddleware(limit, stdlibMiddleware.WithKeyGetter(func(c *gin.Context) string {
		userID := c.GetInt64("userID")
		if userID == 0 {
			return c.ClientIP()
		}
		return fmt.Sprintf("rate-limit:%d", userID)
	}),
		stdlibMiddleware.WithErrorHandler(func(c *gin.Context, err error) {
			c.JSON(200, gin.H{"statusCode": 1, "messgae": "您请求过快！"})
			c.Abort()
		}))
	return mid
}
