package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"txnbi-backend/internal/handle"
	"txnbi-backend/middleware"
	"txnbi-backend/middleware/myLimiter"
)

func Route() *gin.Engine {
	routes := gin.Default()
	routes.Use(middleware.CORSMiddleware())

	// 用户模块
	userGroutMiddle := []gin.HandlerFunc{myLimiter.New("user", myLimiter.LowLevel)}
	userGroup := routes.Group("/user", userGroutMiddle...)
	{
		userGroup.POST("/login", handle.UserLogin)       // 用户登陆接口
		userGroup.POST("/register", handle.UserRegister) //用户注册接口
		authGroup := userGroup.Group("/auth", middleware.AuthUserToken())
		{
			authGroup.GET("/currentUserDetail", handle.CurrentUserDetail)
			authGroup.POST("/loginOut", handle.UserLoginOut)
		}
	}

	// 图表模块
	// 图表模块中间件
	chartGroupMiddle := []gin.HandlerFunc{myLimiter.New("chart", myLimiter.LowLevel)}
	chartGroup := routes.Group("/chart", chartGroupMiddle...)
	{
		// 示例图表接口
		routes.GET("/chart/exampleChart", handle.ExampleChart)
		authGroup := chartGroup.Group("/auth", middleware.AuthUserToken())
		{
			// 生成图表接口因需调用第三方接口，，所以需要更严格的限流
			authGroup.POST("/gen", myLimiter.New("chart-gen", myLimiter.VeryHighLevel), handle.GenChart)
			authGroup.POST("/myChartDel", handle.DeleteMyChart)
			authGroup.GET("/findMyChart", handle.FindMyChart)
			authGroup.POST("/update", handle.UpdateChart)
			authGroup.POST("/share", handle.ShareChart)
		}

	}

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return routes
}
