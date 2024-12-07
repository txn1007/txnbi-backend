package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/ulule/limiter/v3"
	"log"
	"time"

	stdlibMiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	redis2 "github.com/ulule/limiter/v3/drivers/store/redis"
)

var limit *limiter.Limiter

func init() {
	rate := limiter.Rate{
		Period: time.Hour,
		Limit:  20,
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})

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
