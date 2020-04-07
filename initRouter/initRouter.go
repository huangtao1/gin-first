package initRouter

import (
	"com.mark/ginFirst/handler"
	"github.com/gin-gonic/gin"
)

//SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	//加载模板
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	//加载资源
	router.Static("/statics", "./statics")
	//根请求
	indexRouter := router.Group("/")
	{
		indexRouter.GET("", handler.IndexFuncHandler)
		indexRouter.POST("", handler.RetHelloGin)
	}
	//关联user下面所有route
	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", handler.UserLogin)
		userRouter.POST("/register", handler.UserRegister)
	}
	return router
}
