package internal

import (
	"txnbi-backend/internal/handle"
	"txnbi-backend/middleware"
	"txnbi-backend/middleware/myLimiter"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// 管理员模块
	adminGroupMiddle := []gin.HandlerFunc{myLimiter.New("admin", myLimiter.LowLevel), middleware.AuthUserToken()}
	adminGroup := routes.Group("/admin", adminGroupMiddle...)
	{
		// 管理员用户管理接口
		adminGroup.GET("/user/list", handle.AdminListUsers)
		adminGroup.GET("/user/detail", handle.AdminGetUserDetail)
		adminGroup.POST("/user/create", handle.AdminCreateUser)
		adminGroup.POST("/user/update", handle.AdminUpdateUser)
		adminGroup.POST("/user/delete", handle.AdminDeleteUser)

		// 管理员图表管理接口
		adminGroup.GET("/chart/list", handle.AdminListCharts)
		adminGroup.GET("/chart/detail", handle.AdminGetChartDetail)
		adminGroup.POST("/chart/update", handle.AdminUpdateChart)
		adminGroup.POST("/chart/delete", handle.AdminDeleteChart)

		// 管理员日志相关接口
		adminGroup.GET("/log/list", handle.AdminListLogs)
		adminGroup.GET("/log/detail", handle.AdminGetLogDetail)
		adminGroup.POST("/log/create", handle.AdminCreateLog)
		adminGroup.POST("/log/update", handle.AdminUpdateLog)
		adminGroup.POST("/log/delete", handle.AdminDeleteLog)
		adminGroup.POST("/log/batchDelete", handle.AdminBatchDeleteLogs)
	}

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return routes
}
