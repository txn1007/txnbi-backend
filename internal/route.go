package internal

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"txnbi-backend/internal/handle"
	"txnbi-backend/middleware"
)

func Route() *gin.Engine {
	routes := gin.Default()
	routes.Use(middleware.CORSMiddleware())
	userGroup := routes.Group("/user")
	{
		userGroup.POST("/login", handle.UserLogin)       // 用户登陆接口
		userGroup.POST("/register", handle.UserRegister) //用户注册接口
		userGroup.GET("/CurrentUserDetail", middleware.AuthUserToken(), handle.CurrentUserDetail)
		userGroup.POST("/loginOut", middleware.AuthUserToken(), handle.UserLoginOut)
	}
	routes.GET("/chart/exampleChart", handle.ExampleChart)
	chartGroup := routes.Group("/chart", middleware.AuthUserToken())
	{
		//chartGroup.Use(middleware.Limiter()).POST("/gen", handle.GenChart)
		chartGroup.POST("/gen", handle.GenChart)
		chartGroup.POST("/myChartDel", handle.DeleteMyChart)
		chartGroup.GET("/findMyChart", handle.FindMyChart)
	}

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return routes
}
