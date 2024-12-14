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
	userGroup := routes.Group("/user", myLimiter.New("user", myLimiter.LowLevel))
	{
		userGroup.POST("/login", handle.UserLogin)       // 用户登陆接口
		userGroup.POST("/register", handle.UserRegister) //用户注册接口
		userGroup.GET("/currentUserDetail", middleware.AuthUserToken(), handle.CurrentUserDetail)
		userGroup.POST("/loginOut", middleware.AuthUserToken(), handle.UserLoginOut)
	}

	// 示例图表接口
	routes.GET("/chart/exampleChart", myLimiter.New("exampleChart", myLimiter.MidLevel), handle.ExampleChart)

	// 图表模块
	chartGroupMiddle := []gin.HandlerFunc{myLimiter.New("chart", myLimiter.LowLevel), middleware.AuthUserToken()}
	chartGroup := routes.Group("/chart", chartGroupMiddle...)
	{
		// 生成图表接口因需调用第三方接口，，所以需要更严格的限流
		chartGroup.POST("/gen", myLimiter.New("chart-gen", myLimiter.VeryHighLevel), handle.GenChart)
		chartGroup.POST("/myChartDel", handle.DeleteMyChart)
		chartGroup.GET("/findMyChart", handle.FindMyChart)
	}

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return routes
}
