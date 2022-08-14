package router

import (
	"github.com/derekAHua/goLib/middleware"
	"goWeb/controller"

	"github.com/gin-gonic/gin"
)

func Http(engine *gin.Engine) {
	router := engine.Group("/web")

	// 这个会把500的错误转为http code = 200， 特定业务错误码
	router.Use(middleware.Recover)

	testGroup := router.Group("/user")
	{
		testGroup.GET("/getUserInfo", controller.User.GetUserInfo)
	}

}
